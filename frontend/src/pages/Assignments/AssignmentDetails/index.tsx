import { FaChevronLeft } from "react-icons/fa";
import {
  Table,
  TableRow,
  TableCell,
  TableDiv,
} from "@/components/Table/index.tsx";
import Button from "@/components/Button";

import "./styles.css";

const AssignmentDetails: React.FC = () => {
  return (
    <div className="AssignmentDetails">
      <div className="AssignmentDetails__head">
        <div className="AssignmentDetails__title">
          <FaChevronLeft />
          <h2>Assignment 3</h2>
        </div>
        <div className="AssignmentDetails__dates">
          <span>Released: 5 Sep, 10:00am</span>
          <span>Due: 15 Sep, 11:59pm</span>
        </div>
      </div>

      <div className="AssignmentDetails__externalButtons">
        <Button href="">View in Github Classroom</Button>
        <Button href="">View Starter Code</Button>
        <Button href="">View Rubric</Button>
      </div>

      <h2>Metrics</h2>
      <div>Metrics go here</div>

      <h2 style={{ marginBottom: 0 }}>Submissions</h2>
      <Table cols={3}>
        <TableRow style={{ borderTop: "none" }}>
          <TableCell>Student Name</TableCell>
          <TableCell>Status</TableCell>
          <TableCell>Last Commit</TableCell>
        </TableRow>
        {Array.from({ length: 10 }).map((_, i: number) => (
          <TableRow key={i} className="AssignmentDetails__submission">
            <TableCell>Jane Doe</TableCell>
            <TableCell>Passing</TableCell>
            <TableCell>12 Sep, 11:34pm</TableCell>
          </TableRow>
        ))}
      </Table>
    </div>
  );
};

export default AssignmentDetails;
