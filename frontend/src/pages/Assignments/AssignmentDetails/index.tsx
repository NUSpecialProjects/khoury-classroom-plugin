import { FaChevronLeft } from "react-icons/fa";

import "./styles.css";
import Button from "@/components/Button";

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
      <div className="AssignmentDetails__submissions">
        <div className="AssignmentDetails__submission">
          <div>Student Name</div>
          <div>Status</div>
          <div>Last Commit</div>
        </div>
        {Array.from({ length: 20 }).map((_, i: number) => (
          <div key={i} className="AssignmentDetails__submission">
            <div>Jane Doe</div>
            <div>Passing</div>
            <div>12 Sep, 11:34pm</div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default AssignmentDetails;
