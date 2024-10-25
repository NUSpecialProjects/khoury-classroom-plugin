import { FaChevronLeft } from "react-icons/fa";

import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import Button from "@/components/Button";

import "./styles.css";
import { useLocation, useParams } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { SelectedSemesterContext } from "@/contexts/selectedSemester";




const Assignment: React.FC = () => {
  const location = useLocation();
  const [assignment, setAssignment] = useState<IAssignment>()
  const [studentAssignments, setStudentAssignment] = useState<IStudentAssignment[]>([]);
  const { selectedSemester } = useContext(SelectedSemesterContext);
  const { assignmentId } = useParams();


  useEffect(() => {
    const SyncStudentAssignments = async (sem: ISemester, assignment: IAssignment) => {
      try {
        if (sem.classroom_id) {
          const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
          const result = await fetch(`${base_url}/github/sync/assignment/student`, {
            method: 'POST',
            credentials: 'include',
            headers: {
              'Content-Type': 'application/json',
            },  
            body: JSON.stringify({ classroom_id: sem.classroom_id, assignment_id: assignment.assignment_classroom_id })
          });

          if (!result.ok) {
            throw new Error('Network response was not ok');
          }

        }

      } catch (error: unknown) {
        console.error('Error fetching assignments:', error);
      }
    };

    const fetchAssignmentIndirectNav = async (assignmentID: string, sem: ISemester) => {
      try {
        const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
        const result = await fetch(`${base_url}/semester/${sem.classroom_id}/assignments/${assignmentID}`, {
          method: 'GET',
          credentials: 'include',
          headers: {
            'Content-Type': 'application/json',
          },
        });

        if (!result.ok) {
          throw new Error('Network response was not ok');
        }

        const data: IAssignment = (await result.json() as IAssignment)
        setAssignment(data)

      } catch (error: unknown) {
        console.error(`Error fetching assignment ${assignmentID}`, error)
      }
    };

    const fetchStudentAssignments = async (semesterID: number, assignmentID: number) => {
      try {
        const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
        const result = await fetch(`${base_url}/semesters/${semesterID}/assignments/${assignmentID}/student-assignments`,
          {
            method: "GET",
            credentials: "include",
            headers: {
              "Content-Type": "application/json",
            },
          });
        if (!result.ok) {
          throw new Error("Network response was not ok");
        }

        const data: IStudentAssignment[] = (await result.json())
        setStudentAssignment(data) 

      } catch (error: unknown) {
        console.log("Bad fetch, ", error)
      }
    };


    // check if assignment has been passed through 
    if (location.state) {
      setAssignment(location.state.assignment)
      const a: IAssignment = location.state.assignment

      // sync student assignments
      if (selectedSemester !== null && selectedSemester !== undefined) {
        SyncStudentAssignments(selectedSemester, a).then(() => {
          console.log("Sync didn't error, fetching")
          fetchStudentAssignments(selectedSemester.classroom_id, a.assignment_classroom_id)
            .catch((error: unknown) => { console.log("Error fetching: ", error) })

        }).catch((error: unknown) => {
          console.error("Error syncing: ", error)
        })
      }


    } else {
      console.log("Fetching assignment from backend")
      // fetch the assignment from backend
      if (assignmentId && selectedSemester !== null && selectedSemester !== undefined) {
        fetchAssignmentIndirectNav(assignmentId, selectedSemester).catch((error: unknown) => {
          console.error("Could not get assignment: ", error)
        })
      }

    }

  }, [selectedSemester]);



  return (
    <div className="Assignment">
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

          <div className="Assignment__externalButtons">
            <Button href="" variant="secondary">View in Github Classroom</Button>
            <Button href="" variant="secondary">View Starter Code</Button>
            <Button href="" variant="secondary">View Rubric</Button>
          </div>

          <h2>Metrics</h2>
          <div>Metrics go here</div>

          <h2 style={{ marginBottom: 0 }}>Student Assignments</h2>
          <Table cols={3}>
            <TableRow style={{ borderTop: "none" }}>
              <TableCell>Student Name</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Last Commit</TableCell>
            </TableRow>
            {studentAssignments && studentAssignments.length > 0 &&
              studentAssignments.map((sa, i) => (
                <TableRow key={i} className="Assignment__submission">
                  <TableCell>{sa.student_gh_username.join(", ")}</TableCell>
                  <TableCell>Passing</TableCell>
                  <TableCell>12 Sep, 11:34pm</TableCell>
                </TableRow>
              ))
            }
            
          </Table>
        </>
      )}
    </div>

  );
};

export default Assignment;