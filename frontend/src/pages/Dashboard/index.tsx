import { useEffect, useState } from "react";
import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import { Link } from "react-router-dom";



interface Assignment {
    id: number; 
    rubric_id?: number | null; 
    active: boolean;
    assignment_classroom_id: number;
    semester_id: number;
    name: string;
    local_id: number;
    main_due_date: string;
  }

const Dashboard: React.FC = () => {
    const [response, setResponse] = useState<string | null>(null);
    const [assignments, setAssignments] = useState<Assignment[]>([]);
    const [loading, setLoading] = useState<boolean>(false);
    const [activeAssgns, setActiveAssgns] = useState<Assignment[]>([]);
    const [inactiveAssgns, setInactiveAssgns] = useState<Assignment[]>([]);


    // API call when the component is rendered
    useEffect(() => {
        const fetchAssignments = async () => {
            try {
                setLoading(true);
                const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                const result = await fetch(`${base_url}/assignments`, {
                    method: 'GET',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                });

                if (!result.ok) {
                    throw new Error('Network response was not ok');
                }

                const data = await result.json();
                console.log(data)
                setAssignments(data); 
            } catch (error) {
                console.error('Error fetching assignments:', error);
                setResponse('Error fetching assignments');
            } finally {
                setLoading(false);
            }
        };

        const handleButtonClick = async () => {
            try {
                const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                const result = await fetch(`${base_url}/github/sync`, {
                    method: 'POST',
                    credentials: 'include',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ classroom_id: 237209 }), //237210
                })
    
                if (!result.ok) {
                    throw new Error('Network response was not ok');
                }
    
                const data = await result.json();
                setResponse(JSON.stringify(data));
            } catch (error) {
                console.error('Error making API call:', error);
                setResponse('Error fetching data');
            } finally {
                console.log("Done")
            }
        };

        handleButtonClick();

        fetchAssignments();

        

    }, []);

    useEffect(() => {
        assignments.forEach((assignment: Assignment) => {
            if(assignment.active) {
                setActiveAssgns(activeAssgns.concat(assignment))
            } else {
                setInactiveAssgns(inactiveAssgns.concat(assignment))
            }
        }); 
    }, [assignments])

    



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
                <h2 style={{ marginBottom: 0 }}>Active Assignments</h2>
                <Table cols={3}>
                    <TableRow style={{ borderTop: "none" }}>
                        <TableCell>Assignment Name</TableCell>
                        <TableCell>Released</TableCell>
                        <TableCell>Due Date</TableCell>
                    </TableRow>
                    {activeAssgns.map((assignment, i: number) => (
                        <TableRow key={i} className="Assignment__submission">
                            <TableCell> <Link to="/app/assignments/assignmentdetails" className="Dashboard__assignmentLink">{assignment.name}</Link></TableCell>
                            <TableCell>5 Sep, 9:00 AM</TableCell>
                            <TableCell>{assignment.main_due_date}</TableCell>
                        </TableRow>
                    ))}
                </Table>

                <h2 style={{ marginBottom: 0 }}>Inactive Assignments</h2>
                <Table cols={3}>
                    <TableRow style={{ borderTop: "none" }}>
                        <TableCell>Assignment Name</TableCell>
                        <TableCell>Released</TableCell>
                        <TableCell>Due Date</TableCell>
                    </TableRow>
                    {inactiveAssgns.map((assignment, i: number) => (
                        <TableRow key={i} className="Assignment__submission">
                            <TableCell> <Link to="/app/assignments/assignmentdetails" className="Dashboard__assignmentLink">{assignment.name}</Link></TableCell>
                            <TableCell>5 Sep, 9:00 AM</TableCell>
                            <TableCell>{assignment.main_due_date}</TableCell>
                        </TableRow>
                    ))}
                </Table>
            </div>
        </div>
    )
}

export default Dashboard;