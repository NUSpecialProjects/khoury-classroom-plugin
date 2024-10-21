import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import { Link } from "react-router-dom";

const Dashboard: React.FC = () => {
    return (
        <div className="Dashboard">
            <div className="Dashboard__classroomDetailsWrapper">
                <UserGroupCard label="Professors" number={1} />
                <UserGroupCard label="TAs" number={12} />
                <UserGroupCard label="Students" number={38} />
            </div>
            <div className="Dashboard__assignmentsWrapper">
                <h1>Active Assignments</h1>
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
            </div>

            <div className="Dashboard__assignmentsWrapper">
                <h1>Inactive Assignments</h1>
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