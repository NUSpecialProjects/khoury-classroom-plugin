import { useEffect, useState } from "react";
import { FaChevronLeft } from "react-icons/fa6";
import { useLocation } from "react-router-dom";

import "./styles.css";
import { getAssignmentRubric } from "@/api/assignments";
import Button from "@/components/Button";

const AssignmentRubric: React.FC = () => {
  const location = useLocation();
  const assignment = location.state.assignment || {};

  const [rubricData, setRubricData] = useState<IFullRubric>()

  useEffect(() => {

    if (assignment && assignment.rubric_id) {
      (async () => {
        try {
          const rubric = await getAssignmentRubric(assignment.rubric_id)
          if (rubric !== null && rubricData !== undefined) {
            setRubricData(rubric)
          }
        } catch (error) {
          console.error("Could not get rubric: ", error)
        }
      })();
    }


  });



  return (
    <div className="AssignmentRubric">
      {assignment && (
        <>
          <div className="Assignment__head">
            <div className="Assignment__title">
              <FaChevronLeft />
              <h2>{assignment.name}</h2>
            </div>

            <div className="Assignment__dates">
              <span>
                Due Date:{" "}
                {assignment.main_due_date
                  ? assignment.main_due_date.toString()
                  : "N/A"}
              </span>
            </div>
          </div>


          <div className="AssignmentRubric__title"> RUBRIC </div>

          
          {rubricData ? (
            <>
              <div>Your rubric content goes here</div>
            </>
          ) : (
            <div className="AssignmentRubric__noRubric">
              <div>This Assignment does not have a Rubric yet.</div>
              <Button href="" variant="secondary"> Import existing rubric</Button>
              <Button href=""> Add new rubric</Button>
            </div>
          )}



        </>
      )}
    </div>
  );
};

export default AssignmentRubric;
