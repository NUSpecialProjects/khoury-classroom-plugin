import Button from "@/components/Button";

import "./styles.css";
import { useLocation, useParams } from "react-router-dom";
import { useContext, useEffect, useState } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { Table, TableCell, TableRow } from "@/components/Table";
import { FaChevronLeft } from "react-icons/fa6";
import StudentListPage from "@/pages/Users/Student";

const Assignment: React.FC = () => {
  const location = useLocation();
  const [assignment, setAssignment] = useState<IAssignment>()
  const [studentAssignments, setStudentAssignment] = useState<IStudentAssignment[]>([]);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const { assignmentId } = useParams();


  useEffect(() => {
  

    const fetchAssignmentIndirectNav = async (assignmentID: string, classroom: IClassroom) => {
      try {
        const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
        const result = await fetch(`${base_url}/classrooms/classroom/${classroom}/assignments/assignment/${assignmentID}`, {
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

    const fetchStudentAssignments = async (classroomID: number, assignmentID: number) => {
      try {
        const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
        const result = await fetch(`${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works`,
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
      if (selectedClassroom !== null && selectedClassroom !== undefined) {
        fetchStudentAssignments(selectedClassroom.id, a.assignment_classroom_id)
        .catch((error: unknown) => { console.log("Error fetching: ", error) })
      }


    } else {
      console.log("Fetching assignment from backend")
      // fetch the assignment from backend
      if (assignmentId && selectedClassroom !== null && selectedClassroom !== undefined) {
        fetchAssignmentIndirectNav(assignmentId, selectedClassroom).catch((error: unknown) => {
          console.error("Could not get assignment: ", error)
        })
      }

    }

  }, [selectedClassroom]);
  
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
            {/* {studentAssignments && studentAssignments.length > 0 &&
              studentAssignments.map((sa, i) => (
                <TableRow key={i} className="Assignment__submission">
                  <TableCell>{StudentListPage}</TableCell>
                  <TableCell>Passing</TableCell>
                  <TableCell>12 Sep, 11:34pm</TableCell>
                </TableRow>
              ))
            } */}

          </Table>
          </>
      )}
    </div>

  );
};

export default Assignment;