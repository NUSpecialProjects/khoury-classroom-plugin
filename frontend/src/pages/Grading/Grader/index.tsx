import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { useContext, useEffect, useState } from "react";
import { useParams } from "react-router-dom";

import Prism from "prismjs";
import "prismjs/plugins/line-numbers/prism-line-numbers";
import "prismjs/plugins/line-numbers/prism-line-numbers.css";
import "@/assets/prism-vs-dark.css";

import {
  dependencies,
  ext2lang,
  ext2langLoader,
  extractExtension,
} from "./funcs";
import FileTree from "@/components/FileTree";
import Button from "@/components/Button";

import { SelectedSemesterContext } from "@/contexts/selectedSemester";
import { getStudentAssignment } from "@/api/student_assignments";

import "./styles.css";

const Grader: React.FC = () => {
  // params
  const { assignmentId, studentAssignmentId } = useParams();
  const { selectedSemester } = useContext(SelectedSemesterContext);

  // states
  const [studentAssignment, setStudentAssignment] =
    useState<IStudentAssignment | null>(null);
  const [gitTree, setGitTree] = useState<IGitTreeNode[]>([]);
  const [cachedFiles, setCachedFiles] = useState<Record<string, IGraderFile>>(
    {}
  );
  const [currentFile, setCurrentFile] = useState<IGraderFile | null>(null);

  // fetch requested student assignment
  useEffect(() => {
    console.log(selectedSemester);
    // TODO: reroute 404
    if (!selectedSemester || !assignmentId || !studentAssignmentId) return;

    getStudentAssignment(
      selectedSemester?.classroom_id,
      Number(assignmentId),
      Number(studentAssignmentId)
    ).then((resp) => {
      setStudentAssignment(resp);
      console.log(resp);
    });
  }, [studentAssignmentId]);

  // fetch git tree from student assignment repo
  useEffect(() => {
    fetch(
      `http://localhost:8080/file-tree/org/${selectedSemester?.org_name}/assignment/${assignmentId}/student/${studentAssignment?.repo_name}`
    )
      .then((response) => response.json())
      .then((data: IGitTreeNode[]) => {
        setGitTree(data);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  }, [studentAssignment]);

  // when a new file is selected, import any necessary
  // prismjs language syntax files and trigger a rehighlight
  useEffect(() => {
    if (currentFile) {
      const lang = ext2lang[extractExtension(currentFile.name)];
      const loadLanguages = async () => {
        try {
          const deps: string | string[] = dependencies[lang];
          if (deps) {
            console.log(deps);
            if (typeof deps === "string") {
              await ext2langLoader[deps]();
            }
            if (Array.isArray(deps)) {
              for (const dep of deps) {
                await ext2langLoader[dep]();
              }
            }
          }
          await ext2langLoader[lang]();
        } catch (err: unknown) {
          // Prism does not support language or mapping does not exist
          console.log(err);
        }
      };
      loadLanguages()
        .then(() => {
          Prism.highlightAll();
        })
        .catch((err: unknown) => {
          console.log(err);
        });
    }
  }, [currentFile]);

  const openFile = (node: IFileTreeNode) => {
    // Check if the content is already cached
    if (node.sha in cachedFiles) {
      setCurrentFile(cachedFiles[node.sha]);
      return;
    }

    fetch(
      "http://localhost:8080/file-tree/org/NUSpecialProjects/assignment/1/student/92pLytz-SgW~mKeuxDyuJg/blob/" +
        node.sha
    )
      .then((response) => response.text())
      .then((content) => {
        const file: IGraderFile = { content, name: node.name };
        setCurrentFile(file);
        // Cache the content
        setCachedFiles((prev) => ({
          ...prev,
          [node.sha]: file,
        }));
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  };

  return (
    <div className="Grader">
      <div className="Grader__head">
        <div className="Grader__title">
          <FaChevronLeft />
          <div>
            <h2>Assignment 3</h2>
            <span>Jane Doe</span>
          </div>
        </div>
        <div className="Grader__nav">
          <span>Submission 2/74</span>
          <div>
            <Button>
              <FaChevronLeft />
              Previous
            </Button>
            <Button>
              Next
              <FaChevronRight />
            </Button>
          </div>
        </div>
      </div>
      <div className="Grader__body">
        <FileTree
          className="Grader__files"
          gitTree={gitTree}
          selectFileCallback={openFile}
        />
        <div className="Grader__browser">
          <pre className={currentFile ? "line-numbers" : "language-undefined"}>
            <code
              className={
                currentFile
                  ? "line-numbers language-" +
                    ext2lang[extractExtension(currentFile.name)]
                  : "language-undefined"
              }
            >
              {currentFile
                ? currentFile.content
                : "Select a file to view its contents."}
            </code>
          </pre>
        </div>
      </div>
    </div>
  );
};

export default Grader;
