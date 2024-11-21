import { useEffect, useState } from "react";
import { FaChevronLeft } from "react-icons/fa6";
import { Link, useLocation } from "react-router-dom";

import "./styles.css";
import { getRubric } from "@/api/rubrics";
import Button from "@/components/Button";
import NewRubric from "@/pages/Rubrics/NewRubric";

const AssignmentRubric: React.FC = () => {
  const location = useLocation();
  const assignment = location.state.assignment || {};

  const [rubricData, setRubricData] = useState<IFullRubric>()

  useEffect(() => {
    console.log("use effecting")
    if (assignment && assignment.rubric_id) {
      (async () => {
        try {
          const rubric = await getRubric(assignment.rubric_id)
          if (rubric !== null) {
            console.log("Assignment rubric retrieved rubric data, ", rubric)
            setRubricData(rubric)
          } else {
            console.log("no rubric data found from this assignment")
          }

        } catch (error) {
          //do nothing
        }

      })();
    }

    
  }, [assignment]);


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

          {rubricData ? (
            <div>
              <NewRubric assignment={assignment} givenRubricData={rubricData}/>
            </div>

          ) : (

            <div className="AssignmentRubric__noRubric">
              <div className="AssignmentRubric__noRubricTitle">This Assignment does not have a Rubric yet.</div>
              <Button href="" variant="secondary"> Import existing rubric</Button>
              <Link to={`/app/rubrics/new`} state={{ assignment }}>
                <Button href="" > Add new rubric</Button>
              </Link>
            </div>
          )}



        </>
      )}
    </div>
  );
};

export default AssignmentRubric;