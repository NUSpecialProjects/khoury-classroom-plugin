import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { useContext, useEffect, useState } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";

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
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getPaginatedStudentWork } from "@/api/student_works";
import { getFileTree, getFileBlob, createPRComment } from "@/api/file_tree";

import "./styles.css";

const Grader: React.FC = () => {
  const navigate = useNavigate();

  // params
  const { assignmentID, studentWorkID } = useParams();
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  // states
  const [studentWork, setstudentWork] = useState<IPaginatedStudentWork | null>(
    null
  );
  const [gitTree, setGitTree] = useState<IGitTreeNode[]>([]);
  const [cachedFiles, setCachedFiles] = useState<Record<string, IGraderFile>>(
    {}
  );
  const [currentFilePath, setCurrentFilePath] = useState<string>("");
  const [currentFile, setCurrentFile] = useState<IGraderFile | null>(null);

  // fetch requested student assignment
  useEffect(() => {
    if (!selectedClassroom || !assignmentID || !studentWorkID) return;

    getPaginatedStudentWork(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID)
    )
      .then((resp) => {
        console.log(resp);
        setstudentWork(resp);
      })
      .catch((err: unknown) => {
        console.log(err);
        navigate("/404", { replace: true });
      });
  }, [selectedClassroom, assignmentID, studentWorkID]);

  // fetch git tree from student assignment repo
  useEffect(() => {
    if (!selectedClassroom || !assignmentID || !studentWorkID) return;

    getFileTree(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID)
    )
      .then((resp) => {
        setGitTree(resp);
      })
      .catch((err: unknown) => {
        // todo: reroute 404
        console.log(err);
      });
  }, [studentWork]);

  // when a new file is selected, import any necessary
  // prismjs language syntax files and trigger a rehighlight
  useEffect(() => {
    if (currentFile) {
      const lang = ext2lang[extractExtension(currentFile.name)];
      const loadLanguages = async () => {
        try {
          const deps: string | string[] = dependencies[lang];
          if (deps) {
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
    setCurrentFilePath(node.path);

    // Check if the content is already cached
    if (node.sha in cachedFiles) {
      setCurrentFile(cachedFiles[node.sha]);
      return;
    }

    if (!selectedClassroom || !assignmentID || !studentWorkID) return;
    getFileBlob(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID),
      node
    )
      .then((resp) => {
        setCurrentFile(resp);
        setCachedFiles((prev) => ({
          ...prev,
          [node.sha]: resp,
        }));
      })
      .catch((err: unknown) => {
        // todo: reroute 404
        console.log(err);
      });
  };

  const submitComment = (e: React.FormEvent) => {
    e.preventDefault();

    if (
      !selectedClassroom ||
      !studentWork ||
      !gitTree ||
      !studentWork.repo_name
    )
      return;
    const data = new FormData(e.target as HTMLFormElement);
    createPRComment(
      selectedClassroom.org_name,
      studentWork.repo_name,
      currentFilePath,
      Number(data.get("line")),
      String(data.get("comment"))
    );
  };

  return (
    studentWork && (
      <div className="Grader">
        <div className="Grader__head">
          <div className="Grader__title">
            <Link to="/app/grading">
              <FaChevronLeft />
            </Link>
            <div>
              <h2>{studentWork.assignment_name}</h2>
              <span>{studentWork.contributors}</span>
            </div>
          </div>
          <div className="Grader__nav">
            <span>
              Student Work {studentWork.row_num}/
              {studentWork.total_student_works}
            </span>
            <div>
              {studentWork.previous_student_work_id && (
                <Link
                  to={`/app/grading/assignment/${assignmentID}/student/${studentWork.previous_student_work_id}`}
                >
                  <Button variant="secondary">
                    <FaChevronLeft />
                    Previous
                  </Button>
                </Link>
              )}
              {studentWork.next_student_work_id && (
                <Link
                  to={`/app/grading/assignment/${assignmentID}/student/${studentWork.next_student_work_id}`}
                >
                  <Button variant="secondary">
                    Next
                    <FaChevronRight />
                  </Button>
                </Link>
              )}
            </div>
          </div>
        </div>
        {gitTree && (
          <div className="Grader__body">
            <FileTree
              className="Grader__files"
              gitTree={gitTree}
              selectFileCallback={openFile}
            />
            <div className="Grader__browser">
              <pre
                className={currentFile ? "line-numbers" : "language-undefined"}
              >
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
        )}
        <div>
          <form onSubmit={submitComment}>
            <input type="number" name="line" placeholder="line number" />
            <input type="text" name="comment" placeholder="comment" />
            <button type="submit">submit comment</button>
          </form>
        </div>
      </div>
    )
  );
};

export default Grader;
