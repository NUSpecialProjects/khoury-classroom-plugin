import { useContext, useEffect, useState } from "react";
import { FaChevronLeft } from "react-icons/fa6";
import { Link, useLocation } from "react-router-dom";

import "./styles.css";
import { getRubric, getRubricsInClassroom } from "@/api/rubrics";
import Button from "@/components/Button";
import { Table, TableCell, TableRow } from "@/components/Table";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { setAssignmentRubric } from "@/api/assignments";


const AssignmentRubric: React.FC = () => {
  const location = useLocation();
  const { selectedClassroom } = useContext(SelectedClassroomContext)


  const [assignment, setAssignmentData] = useState<IAssignmentOutline>()
  const [rubricData, setRubricData] = useState<IFullRubric>()
  const [rubrics, setRubrics] = useState<IFullRubric[]>([])

  const [importing, setImporting] = useState(false)


  const choseExisting = async (event: React.ChangeEvent<HTMLSelectElement>) => {
    const selectedId = parseInt(event.target.value, 10);
    if (selectedId) {
      const selectedRubric = rubrics.find((rubric) => rubric.rubric.id === selectedId);
      if (selectedRubric) {
        setRubricData(selectedRubric)
        if (selectedClassroom && assignment) {
          setAssignmentRubric(
            selectedRubric.rubric.id!,
            selectedClassroom.id,
            assignment.id)
    
        }
      }
    }
  };


  const notImporting = () => {
    setImporting(false)
  }

  const allowImporting = () => {
    setImporting(true)
  }


  useEffect(() => {
    if (location.state) {
      setAssignmentData(location.state.assignment)
      const aData = location.state.assignment
      if (aData && aData.rubric_id) {
        (async () => {
          try {
            const rubric = await getRubric(aData.rubric_id)
            if (rubric !== null) {
              setRubricData(rubric)
            }
          } catch (_) {
            //do nothing
          }

        })();
      }
    }

    if (selectedClassroom) {
      (async () => {
        try {
          const retrievedRubrics = await getRubricsInClassroom(selectedClassroom.id)
          if (retrievedRubrics !== null) {
            setRubrics(retrievedRubrics)
          }
        } catch (_) {
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
              <div className="AssignmentRubric__title">
                {rubricData.rubric.name}

                <Link to={`/app/rubrics/new`} state={{ assignment, rubricData }}>
                  <Button href=""> Edit Rubric </Button>
                </Link>


              </div>


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


              {importing ?
                <div>
                  <select id="dropdown" onChange={choseExisting}>
                    <option value="">-- Select a rubric --</option>
                    {rubrics.map((rubric) => (
                      <option key={rubric.rubric.id} value={rubric.rubric.id!}>
                        {rubric.rubric.name}
                      </option>
                    ))}
                  </select>

                  <Button href="" variant="secondary" onClick={notImporting}>
                    Cancel
                  </Button>
                </div>

                :
                <div>
                  <Button href="" variant="secondary" onClick={allowImporting}> Import existing rubric</Button>


                  <Link to={`/app/rubrics/new`} state={{ assignment }}>
                    <Button href="" > Add new rubric</Button>
                  </Link>
                </div>

              }
            </div>
          )}
        </>
      )}
    </div>
  );
};

export default AssignmentRubric;