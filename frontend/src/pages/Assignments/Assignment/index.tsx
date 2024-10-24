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
  const [studentAssignment, setStudentAssignment] = useState<IStudentAssignment[]>([]);
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
            body: JSON.stringify( { classroom_id: sem.classroom_id, assignment_id: assignment.assignment_classroom_id} ) 
          });

          if (!result.ok) {
            throw new Error('Network response was not ok');
          }

          const data: IStudentAssignment[] = (await result.json() as IStudentAssignment[])
          // Need to actually store info
          console.log(data)
          setStudentAssignment(data)
        }

      } catch (error: unknown) {
        console.error('Error fetching assignments:', error);
      }
    };

    const fetchAssignment = async (assignmentID: string, sem: ISemester) => {
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



    // check if assignment has been passed through 
    if (location.state) {
      setAssignment(location.state.assignment)
      const a = location.state.assignment
      // sync student assignments
      if (selectedSemester !== null && selectedSemester !== undefined) {
        SyncStudentAssignments(selectedSemester, a).then(() => {
          console.log(studentAssignment)
        }).catch((error: unknown) => {
          console.error("Error syncing: ", error)
        })
      }


    } else {
      // fetch the assignment from backend
      if (assignmentId && selectedSemester !== null && selectedSemester !== undefined) {
        fetchAssignment(assignmentId, selectedSemester).catch((error: unknown) => {
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
            {Array.from({ length: 10 }).map((_, i: number) => (
              <TableRow key={i} className="Assignment__submission">
                <TableCell>Jane Doe</TableCell>
                <TableCell>Passing</TableCell>
                <TableCell>12 Sep, 11:34pm</TableCell>
              </TableRow>
            ))}
          </Table>
        </>
      )}
    </div>

  );
};

export default Assignment;