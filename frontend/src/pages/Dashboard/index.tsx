import { useEffect, useState } from "react";
import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import { Link } from "react-router-dom";
import useSelectedSemester from "@/contexts/useClassroom";


interface IAssignment {
    id: number; 
    rubric_id: number | null; 
    active: boolean;
    assignment_classroom_id: number;
    semester_id: number;
    name: string;
    local_id: number;
    main_due_date: Date | null;
  }

const Dashboard: React.FC = () => {
    const [assignments, setAssignments] = useState<IAssignment[]>([]);
    const {selectedSemester} = useSelectedSemester();


    const options: Intl.DateTimeFormatOptions = {
        weekday: 'short', year: 'numeric', month: 'short', day: 'numeric',
        hour: '2-digit', minute: '2-digit', timeZoneName: 'short'
      };

    // API call when the component is rendered
    useEffect(() => {
        const fetchAssignments = async (semester: Semester) => {
            try {
                if (semester.id) {
                    const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                    const result = await fetch(`${base_url}/assignments/${semester.id}`, {
                        method: 'GET',
                        credentials: 'include',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                    });
    
                    if (!result.ok) {
                        throw new Error('Network response was not ok');
                    }
    
                    const data: IAssignment[] = (await result.json() as IAssignment[])
                    const assignmentGoodDate = data.map((assignment: IAssignment) => ({
                        ...assignment,
                        main_due_date: assignment.main_due_date ? new Date(assignment.main_due_date) : null,
                    }))
                    console.log("Setting Assignment data: ", assignmentGoodDate)
                    setAssignments(assignmentGoodDate); 
                }
                
            } catch (error: unknown) {
                console.error('Error fetching assignments:', error);
            }
        };

        const SyncWithClassroom = async (semester: Semester) => {
            try {
                const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                const result = await fetch(`${base_url}/github/sync`, {
                    method: 'POST',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ classroom_id: semester.classroom_id}),
                })
    
                if (!result.ok) {
                    throw new Error('Network response was not ok');
                }
    
            } catch (error: unknown) {
                console.error('Error making API call:', error);
            }
        };

        console.log("We in dashboard: ", selectedSemester)
        if (selectedSemester !== null && selectedSemester !== undefined) {
            SyncWithClassroom(selectedSemester).then(() => {
                fetchAssignments(selectedSemester).catch((error: unknown) => {
                    console.log("Error fetching:", error)
                })
            }).catch((error: unknown) => {
                console.error('Error syncing:', error);
            });
        }

    }, [selectedSemester]);

    return (
        <div className="Dashboard">
            {/* Header group cards */}
            <div className="Dashboard__classroomDetailsWrapper">
                <UserGroupCard label="Professors" number={1} />
                <UserGroupCard label="TAs" number={12} />
                <UserGroupCard label="Students" number={38} />
            </div>

            {/* Assignments */}
            <div className="Dashboard__assignmentsWrapper">
                <h2 style={{ marginBottom: 0 }}>Assignments</h2>
                <Table cols={2}>
                    <TableRow style={{ borderTop: "none" }}>
                        <TableCell>Assignment Name</TableCell>
                        <TableCell>Due Date</TableCell>
                    </TableRow>
                    {assignments.map((assignment, i: number) => (
                        <TableRow key={i} className="Assignment__submission">
                            <TableCell> <Link to={`/app/assignments/?id=${assignment.id}`} className="Dashboard__assignmentLink">{assignment.name}</Link></TableCell>
                            <TableCell> {assignment.main_due_date ? assignment.main_due_date.toLocaleDateString("en-US", options) : "N/A"}</TableCell>
                        </TableRow>
                    ))}
                </Table>

            </div>
        </div>
    )
}

export default Dashboard;