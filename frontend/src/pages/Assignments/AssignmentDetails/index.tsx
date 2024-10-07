import { FaChevronLeft } from "react-icons/fa";

import "./styles.css";
import Button from "@/components/Button";

const AssignmentDetails: React.FC = () => {
  return (
    <div className="AssignmentDetails">
      <div className="AssignmentDetails__head">
        <div className="AssignmentDetails__title">
          <FaChevronLeft />
          <span>Assignment 3</span>
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
    </div>
  );
};

export default AssignmentDetails;
