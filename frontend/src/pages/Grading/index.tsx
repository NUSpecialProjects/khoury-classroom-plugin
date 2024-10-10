import { FaChevronUp, FaChevronDown } from "react-icons/fa";
import { Table, TableRow, TableCell } from "@/components/Table";

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
      <Table className="Grading__assignments">
        <TableRow>
          <TableCell></TableCell>
          <TableCell>Assignment Name</TableCell>
          <TableCell>Status</TableCell>
          <TableCell>Due Date</TableCell>
        </TableRow>
        {Array.from({ length: 20 }).map((_, i: number) => (
          <TableRow>
            <TableCell>
              <FaChevronDown />
            </TableCell>
            <TableCell>Assignment Name</TableCell>
            <TableCell>Status</TableCell>
            <TableCell>Due Date</TableCell>
          </TableRow>
        ))}
      </Table>
    </div>
  );
};

export default Grading;
