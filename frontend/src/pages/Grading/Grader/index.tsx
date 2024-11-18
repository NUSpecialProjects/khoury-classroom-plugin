import { FaChevronLeft, FaChevronRight } from "react-icons/fa";
import { useContext, useEffect, useState } from "react";
import { useParams, useNavigate, Link } from "react-router-dom";

import FileTree from "@/components/FileTree";
import Button from "@/components/Button";
import CodeBrowser from "@/components/CodeBrowser";
import { GraderContext, GraderProvider } from "@/contexts/grader";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getPaginatedStudentWork } from "@/api/student_works";

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
  const {
    assignmentID,
    studentWorkID,
    studentWork,
    selectedFile,
    setSelectedFile,
    postFeedback,
  } = useContext(GraderContext);

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
            selectFileCallback={setSelectedFile}
          />
          <CodeBrowser
            assignmentID={assignmentID}
            studentWorkID={studentWorkID}
            file={selectedFile}
          />
          <button onClick={postFeedback}>POST COMMENTS</button>
        </div>
      </div>
    )
  );
};

export default GraderWrapper;
