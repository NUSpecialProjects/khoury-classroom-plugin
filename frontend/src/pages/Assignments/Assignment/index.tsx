import Button from "@/components/Button";

import "./styles.css";
import { useLocation, useParams } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Table, TableCell, TableRow } from "@/components/Table";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { getAssignmentIndirectNav } from "@/api/assignments";
import { getStudentWorks } from "@/api/student_works";

const Assignment: React.FC = () => {
  const location = useLocation();
  const [assignment, setAssignment] = useState<IAssignmentOutline>();
  const [studentWorks, setStudentAssignment] = useState<IStudentWork[]>([]);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { id } = useParams();

  useEffect(() => {
    // check if assignment has been passed through
    if (location.state) {
      setAssignment(location.state.assignment);
      const a: IAssignmentOutline = location.state.assignment;

      // sync student assignments
      if (selectedClassroom !== null && selectedClassroom !== undefined) {
        (async () => {
          try {
            const studentWorks = await getStudentWorks(
              selectedClassroom.id,
              a.id
            );
            if (studentWorks !== null && studentWorks !== undefined) {
              setStudentAssignment(studentWorks);
            }
          } catch (_) {
            // do nothing
          }
        })();
      }
    } else {
      // fetch the assignment from backend
      if (id && selectedClassroom !== null && selectedClassroom !== undefined) {
        (async () => {
          try {
            const fetchedAssignment = await getAssignmentIndirectNav(
              selectedClassroom.id,
              +id
            );
            if (fetchedAssignment !== null && fetchedAssignment !== undefined) {
              setAssignment(fetchedAssignment);
              const studentWorks = await getStudentWorks(
                selectedClassroom.id,
                fetchedAssignment.id
              );
              if (studentWorks !== null && studentWorks !== undefined) {
                setStudentAssignment(studentWorks);
              }
            }
          } catch (_) {
            // do nothing
          }
        })();
      }
    }
  }, [selectedClassroom]);

  return (
    <div className="Assignment">
      {assignment && (
        <>
          <SubPageHeader pageTitle={assignment.name}>
            <div className="Assignment__dates">
              <div className="Assignment__date">
                <div className="Assignment__date--title"> {"Released on:"}</div>
                {assignment.created_at
                  ? new Date(assignment.created_at).toLocaleString("en-US", {
                    month: "long",
                    day: "numeric",
                    year: "numeric",
                    hour: "2-digit",
                    minute: "2-digit",
                    hour12: true,
                  })
                  : "N/A"}
              </div>
              <div className="Assignment__date">
                <div className="Assignment__date--title"> {"Due Date:"}</div>
                {assignment.main_due_data
                  ? new Date(assignment.main_due_data).toLocaleString("en-US", {
                    month: "long",
                    day: "numeric",
                    year: "numeric",
                    hour: "2-digit",
                    minute: "2-digit",
                    hour12: true,
                  })
                  : "N/A"}
              </div>
            </div>
          </SubPageHeader>

          <div className="Assignment__externalButtons">
            <Button href="" variant="secondary">
              View in Github Classroom
            </Button>
            <Button href="" variant="secondary">
              View Starter Code
            </Button>
            <Button href="" variant="secondary">
              View Rubric
            </Button>
          </div>

          <div className="Assignment__subSectionWrapper">
            <h2>Metrics</h2>
            <p>Metrics go here</p>
          </div>

          <div className="Assignment__subSectionWrapper">
            <h2 style={{ marginBottom: 0 }}>Student Assignments</h2>
            <Table cols={3}>
              <TableRow style={{ borderTop: "none" }}>
                <TableCell>Student Name</TableCell>
                <TableCell>Status</TableCell>
                <TableCell>Last Commit</TableCell>
              </TableRow>
              {studentWorks &&
                studentWorks.length > 0 &&
                studentWorks.map((sa, i) => (
                  <TableRow key={i} className="Assignment__submission">
                    <TableCell>{sa.contributors.join(", ")}</TableCell>
                    <TableCell>Passing</TableCell>
                    <TableCell>12 Sep, 11:34pm</TableCell>
                  </TableRow>
                ))}
            </Table>
          </div>
        </>
      )}
    </div>
  );
};

export default Assignment;
