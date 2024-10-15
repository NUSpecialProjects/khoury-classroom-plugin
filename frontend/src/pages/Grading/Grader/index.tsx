import { FaChevronLeft } from "react-icons/fa";

import "./styles.css";

const Grader: React.FC = () => {
  return (
    <div className="Grader">
      <div className="Grader__head">
        <div className="Grader__title">
          <FaChevronLeft />
          <div>
            <h2>Assignment 2</h2> <span>Jane Doe</span>
          </div>
        </div>
        <div className="Grader__dates">
          <span>Released: 5 Sep, 10:00am</span>
          <span>Due: 15 Sep, 11:59pm</span>
        </div>
      </div>
    </div>
  );
};

export default Grader;
