import { useState } from "react";
import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import { Link } from "react-router-dom";

const Dashboard: React.FC = () => {
    const [response, setResponse] = useState<string | null>(null);


    const handleButtonClick = async () => {
    
        try {
            const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
            const result = await fetch(`${base_url}/github/sync`, {
                method: 'POST', 
                credentials: 'include',
                headers: {
                  'Content-Type': 'application/json',
                },
                body: JSON.stringify({classroom_id : 237209}), //237210
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
                    {Array.from({ length: 1 }).map((_, i: number) => (
                        <TableRow key={i} className="Assignment__submission">
                            <TableCell> <Link to="/app/assignments/assignmentdetails" className="Dashboard__assignmentLink">Assignment 1</Link></TableCell>
                            <TableCell>5 Sep, 9:00 AM</TableCell>
                            <TableCell>15 Sep, 11:59 PM</TableCell>
                        </TableRow>
                    ))}
                </Table>
                <button onClick={handleButtonClick}>
                    ASSINGMENT DATA SYNC
                </button>
                <h2 style={{ marginBottom: 0 }}>Inactive Assignments</h2>
                <Table cols={3}>
                    <TableRow style={{ borderTop: "none" }}>
                        <TableCell>Assignment Name</TableCell>
                        <TableCell>Released</TableCell>
                        <TableCell>Due Date</TableCell>
                    </TableRow>
                    {Array.from({ length: 2 }).map((_, i: number) => (
                        <TableRow key={i} className="Assignment__submission">
                            <TableCell> <Link to="/app/assignments/assignmentdetails" className="Dashboard__assignmentLink">Assignment 1</Link></TableCell>
                            <TableCell>5 Sep, 9:00 AM</TableCell>
                            <TableCell>15 Sep, 11:59 PM</TableCell>
                        </TableRow>
                    ))}
                </Table>
            </div>
        </div>
    )
}

export default Dashboard;