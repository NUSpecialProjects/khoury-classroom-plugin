import { useEffect, useState } from "react";
import { FaChevronDown, FaChevronLeft, FaChevronRight } from "react-icons/fa6";
import { Link, Navigate, useLocation } from "react-router-dom";

import "./styles.css";
import { getRubric } from "@/api/rubrics";
import Button from "@/components/Button";
import { Table, TableCell, TableDiv, TableRow } from "@/components/Table";
import RubricItem from "@/components/RubricItem";


const AssignmentRubric: React.FC = () => {
  const location = useLocation();

  const [assignment, setAssignmentData] = useState<IAssignmentOutline>()
  const [rubricData, setRubricData] = useState<IFullRubric>()

  useEffect(() => {
    if (location.state) {
      setAssignmentData(location.state.assignment)
      const aData = location.state.assignment
      if (aData && aData.rubric_id) {
        (async () => {
          try {
            const rubric = await getRubric(aData.rubric_id)
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
    } else {
      console.log("no assignment state")
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
              <Table cols={2}>
                <TableRow>
                  <TableCell>Explanation </TableCell>
                  <TableCell>Point Value </TableCell>
                </TableRow>
                 
              {rubricData &&
                rubricData.rubric_items.map((item, i: number) => (
                    <TableRow key={i}>
                        <TableCell> {item.explanation} </TableCell>
                        <TableCell> {item.point_value > 0 ? "+" + item.point_value : item.point_value} </TableCell>
                    </TableRow>
              ))}
              </Table>
             
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