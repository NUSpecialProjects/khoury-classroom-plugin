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
<<<<<<< HEAD
import {
  getStudentWork,
  getGitTree,
  getGitBlob,
  getTotalStudentAssignments,
} from "@/api/student_assignments";
=======
import { getPaginatedStudentWork } from "@/api/student_works";
import { getFileTree, getFileBlob, createPRComment } from "@/api/grading";
>>>>>>> main

import "./styles.css";

const Grader: React.FC = () => {
  const navigate = useNavigate();

  // params
<<<<<<< HEAD
  const { assignmentId, studentAssignmentId } = useParams();
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  // states
  const [totalStudentAssignments, setTotalStudentAssignments] = useState(0);
  const [studentAssignment, setStudentAssignment] =
    useState<IStudentWork | null>(null);
=======
  const { assignmentID, studentWorkID } = useParams();
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  // states
  const [studentWork, setStudentWork] = useState<IPaginatedStudentWork | null>(
    null
  );
>>>>>>> main
  const [gitTree, setGitTree] = useState<IGitTreeNode[]>([]);
  const [cachedFiles, setCachedFiles] = useState<Record<string, IGraderFile>>(
    {}
  );
  const [currentFilePath, setCurrentFilePath] = useState<string>("");
  const [currentFile, setCurrentFile] = useState<IGraderFile | null>(null);
<<<<<<< HEAD

  // fetch totals for indexing purposes
  useEffect(() => {
    if (!selectedClassroom || !assignmentId || !studentAssignmentId) return;

    getTotalStudentAssignments(selectedClassroom.id, Number(assignmentId))
      .then((resp) => {
        setTotalStudentAssignments(resp);
      })
      .catch((err: unknown) => {
        console.log(err);
      });
  }, [selectedClassroom, assignmentId]);

  // fetch requested student assignment
  useEffect(() => {
    if (!selectedClassroom || !assignmentId || !studentAssignmentId) return;

    getStudentWork(
      selectedClassroom.id,
      Number(assignmentId),
      Number(studentAssignmentId)
=======
  const [comments, setComments] = useState<IGradingComment[]>([]);

  // fetch requested student assignment
  useEffect(() => {
    // reset states
    setCurrentFilePath("");
    setCurrentFile(null);
    setComments([]);
    setStudentWork(null);
    setGitTree([]);

    if (!selectedClassroom || !assignmentID || !studentWorkID) return;

    getPaginatedStudentWork(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID)
>>>>>>> main
    )
      .then((resp) => {
        setStudentWork(resp);
      })
      .catch((_: unknown) => {
        navigate("/404", { replace: true });
      });
<<<<<<< HEAD
  }, [selectedClassroom, assignmentId, studentAssignmentId]);

  // fetch git tree from student assignment repo
  useEffect(() => {
    if (!selectedClassroom || !studentAssignment) return;

    getGitTree(selectedClassroom.org_name, (studentAssignment.repo_name ? studentAssignment.repo_name : ""))
=======
  }, [studentWorkID]);

  // fetch git tree from student assignment repo
  useEffect(() => {
    if (!selectedClassroom || !assignmentID || !studentWorkID || !studentWork)
      return;

    getFileTree(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID)
    )
>>>>>>> main
      .then((resp) => {
        setGitTree(resp);
      })
      .catch((_: unknown) => {
        setGitTree([]);
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
        } catch (_: unknown) {
          // Prism does not support language or mapping does not exist
          // do nothing
        }
      };
      loadLanguages()
        .then(() => {
          Prism.highlightAll();
        })
        .catch((_: unknown) => {
          // do nothing
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

<<<<<<< HEAD
    if (!selectedClassroom || !studentAssignment) return;
    getGitBlob(selectedClassroom.org_name, (studentAssignment.repo_name ? studentAssignment.repo_name : ""), node)
=======
    if (!selectedClassroom || !assignmentID || !studentWorkID) return;
    getFileBlob(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID),
      node
    )
>>>>>>> main
      .then((resp) => {
        setCurrentFile(resp);
        setCachedFiles((prev) => ({
          ...prev,
          [node.sha]: resp,
        }));
      })
      .catch((_: unknown) => {
        // todo: reroute 404
      });
  };

  const saveComment = (e: React.FormEvent) => {
    e.preventDefault();

    if (!selectedClassroom || !assignmentID || !studentWorkID) return;
    const form = e.target as HTMLFormElement;
    const data = new FormData(form);
    const comment: IGradingComment = {
      path: currentFilePath,
      line: Number(data.get("line")),
      body: String(data.get("comment")),
    };
    setComments([...comments, comment]);
    form.reset();
  };

  const submitComments = () => {
    if (!selectedClassroom || !assignmentID || !studentWorkID) return;
    createPRComment(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID),
      comments
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
<<<<<<< HEAD
              <h2>{studentAssignment.assignment_name}</h2>
              <span>{studentAssignment.contributors}</span>
=======
              <h2>{studentWork.assignment_name}</h2>
              <span>{studentWork.contributors}</span>
>>>>>>> main
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
                  : gitTree.length == 0
                    ? "Student has not submitted work for grading yet or repository is empty."
                    : "Select a file to view its contents."}
              </code>
            </pre>
          </div>
        </div>
        <div style={{ display: "flex", justifyContent: "space-between" }}>
          <form onSubmit={saveComment}>
            <input type="number" name="line" placeholder="line number" />
            <input type="text" name="comment" placeholder="comment" />
            <button type="submit">SAVE COMMENT</button>
          </form>
          <button onClick={submitComments}>POST COMMENTS</button>
        </div>
      </div>
    )
  );
};

export default Grader;
