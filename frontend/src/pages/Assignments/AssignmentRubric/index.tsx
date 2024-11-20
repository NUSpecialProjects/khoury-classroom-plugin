import { useEffect, useState } from "react";
import { FaChevronLeft } from "react-icons/fa6";
import { Link, useLocation } from "react-router-dom";

import "./styles.css";
import { getRubric } from "@/api/rubrics";
import Button from "@/components/Button";
import NewRubric from "@/pages/Rubric/NewRubric";

const AssignmentRubric: React.FC = () => {
  const location = useLocation();
  const assignment = location.state.assignment || {};

  const [rubricData, setRubricData] = useState<IFullRubric>()

  useEffect(() => {

    if (assignment && assignment.rubric_id) {
      (async () => {
        try {
          const rubric = await getRubric(assignment.rubric_id)
          if (rubric !== null) {
            setRubricData(rubric)
            console.log(rubricData)
          }
        } catch (error) {
          console.error("Could not get rubric: ", error)
        }

      })();
    }

    
  }, []);


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
              <NewRubric rubricData={rubricData}/>
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