import { FaChevronUp, FaChevronDown } from "react-icons/fa";

import "./styles.css";

interface IAssignmentRow {
  name: string;
  status: string;
  date: string;
}

interface IGradingAssignmentRow {
  assignment: IAssignmentRow;
  submissions: IGradingAssignmentSubmission[];
}

interface IGradingAssignmentSubmission {
  studentName: string;
  submissionId: number;
}

const GradingAssignmentRow: React.FC<IGradingAssignmentRow> = ({
  assignment,
  submissions,
}) => {
  return (
    <>
      <div className="GradingAssignmentRow">
        <div className="GradingAssignmentRow__inner">
          <div className="GradingAssignmentRow__chevron">
            <FaChevronDown />
          </div>
          <div className="GradingAssignmentRow__cell">{assignment.name}</div>
          <div className="GradingAssignmentRow__cell">{assignment.status}</div>
          <div className="GradingAssignmentRow__cell">{assignment.date}</div>
        </div>
      </div>
      <div className="GradingAssignmentRow__submissions">
        <div className="GradingAssignmentRow__submissionsList">
          sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss
          <div className="GradingAssignmentRow__submission">
            <div>Student Name</div>
            <div>Grade</div>
          </div>
          <div className="GradingAssignmentRow__submission">
            <div>Student Name</div>
            <div>Grade</div>
          </div>
          <div className="GradingAssignmentRow__submission">
            <div>Student Name</div>
            <div>Grade</div>
          </div>
        </div>
      </div>
    </>
  );
};

const Grading: React.FC = () => {
  return (
    <div className="Grading">
      <h2 style={{ marginBottom: 0 }}>Assignments</h2>
      <div className="Grading__assignments">
        <div className="GradingAssignmentRow">
          <div className="GradingAssignmentRow__inner">
            <div className="GradingAssignmentRow__cell"></div>
            <div className="GradingAssignmentRow__cell">Assignment Name</div>
            <div className="GradingAssignmentRow__cell">Status</div>
            <div className="GradingAssignmentRow__cell">Due Date</div>
          </div>
        </div>
        {Array.from({ length: 20 }).map((_, i: number) => (
          <GradingAssignmentRow
            assignment={{
              name: "Assignment " + (i + 1),
              status: "Active",
              date: "12 Sep, 11:34pm",
            }}
            submissions={[]}
            key={i}
          />
        ))}
      </div>
    </div>
  );
};

export default Grading;
