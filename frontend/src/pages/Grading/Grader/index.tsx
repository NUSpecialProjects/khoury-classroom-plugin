import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { useContext, useEffect, useState } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";

import FileTree from "@/components/FileTree";
import Button from "@/components/Button";
import CodeBrowser from "@/components/CodeBrowser";
import { GraderContext, GraderProvider } from "@/contexts/grader";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getPaginatedStudentWork } from "@/api/student_works";
import { createPRComment } from "@/api/grading";

import "./styles.css";

const GraderWrapper: React.FC = () => {
  const { assignmentID, studentWorkID } = useParams();
  return (
    <GraderProvider assignmentID={assignmentID} studentWorkID={studentWorkID}>
      <Grader />
    </GraderProvider>
  );
};

const Grader: React.FC = () => {
  const navigate = useNavigate();

  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { assignmentID, studentWorkID, comments } = useContext(GraderContext);

  const [studentWork, setStudentWork] = useState<IPaginatedStudentWork | null>(
    null
  );
  const [currentFile, setCurrentFile] = useState<IFileTreeNode | null>(null);

  // fetch requested student assignment
  useEffect(() => {
    // reset states
    setCurrentFile(null);
    setStudentWork(null);

    if (!selectedClassroom || !assignmentID || !studentWorkID) return;

    getPaginatedStudentWork(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID)
    )
      .then((resp) => {
        setStudentWork(resp);
      })
      .catch((err: unknown) => {
        console.log(err);
        navigate("/404", { replace: true });
      });
  }, [studentWorkID]);

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
              <h2>{studentWork.contributors.join(", ")}</h2>
              <span>{studentWork.assignment_name}</span>
            </div>
          </div>
          <div className="Grader__nav">
            <span>
              Student Work {studentWork.row_num}/
              {studentWork.total_student_works}
            </span>
            <div>
              <div className="Grader__navButtons">
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
        </div>
        <div className="Grader__body">
          <FileTree
            className="Grader__files"
            selectFileCallback={setCurrentFile}
          />
          <CodeBrowser
            assignmentID={assignmentID}
            studentWorkID={studentWorkID}
            file={currentFile}
          />
          <button onClick={submitComments}>POST COMMENTS</button>
        </div>
      </div>
    )
  );
};

export default GraderWrapper;
