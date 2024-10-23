import { FaChevronLeft } from "react-icons/fa";

import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import Button from "@/components/Button";

import "./styles.css";

const Assignment: React.FC = () => {
  return (
    <div className="Assignment">
      <div className="Assignment__head">
        <div className="Assignment__title">
          <FaChevronLeft />
          <h2>Assignment 3</h2>
        </div>
        <div className="Assignment__dates">
          <span>Released: 5 Sep, 10:00am</span>
          <span>Due: 15 Sep, 11:59pm</span>
        </div>
      </div>

      <div className="Assignment__externalButtons">
        <Button href="" variant="secondary">View in Github Classroom</Button>
        <Button href="" variant="secondary">View Starter Code</Button>
        <Button href="" variant="secondary">View Rubric</Button>
      </div>

      <div className="Assignment__contentWrapper">
        <div className="Assignment__metricsWrapper">
          <h1>Metrics</h1>
          <div>Metrics go here</div>
        </div>

        <div className="Assignment__assignmentsWrapper">
          <h1>Student Assignments</h1>
          <Table cols={3}>
            <TableRow style={{ borderTop: "none" }}>
              <TableCell>Student Name</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Last Commit</TableCell>
            </TableRow>
            {Array.from({ length: 10 }).map((_, i: number) => (
              <TableRow key={i} className="Assignment__submission">
                <TableCell>Jane Doe</TableCell>
                <TableCell>Passing</TableCell>
                <TableCell>12 Sep, 11:34pm</TableCell>
              </TableRow>
            ))}
          </Table>
        </div>
      </div>
    </div>
  );
};

export default Assignment;
