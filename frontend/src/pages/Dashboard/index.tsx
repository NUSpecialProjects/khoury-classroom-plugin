/* eslint-disable */

import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import { Link } from "react-router-dom";
import useSelectedSemester from "@/contexts/useClassroom";
import AlertBanner from "@/components/Banner/AlertBanner";
import { Semester } from "@/types/semester";
import { activateSemester, deactivateSemester } from "@/api/semesters";

const Dashboard: React.FC = () => {
    const [selectedSemester, setSelectedSemester] = useSelectedSemester();

    console.log("VIEWING DASHBOARD FOR SEMESTER: ", selectedSemester);

    const handleActivate = async (newSemester: Semester) => {
        console.log("Activated semester:", newSemester);
        setSelectedSemester(newSemester);
    }

    const handleActivateClick = async () => {
        if (selectedSemester) {
            const newSemester = await activateSemester(selectedSemester.id);
            handleActivate(newSemester);
        }
    };

    const handleDeactivateClick = async () => {
        if (selectedSemester) {
            const newSemester = await deactivateSemester(selectedSemester.id);
            handleActivate(newSemester);
        }
    };

//TODO: think about where the deactivate button should be
    return (
        <div className="Dashboard">
            {selectedSemester && (
                console.log("Showing alertbanner: ", selectedSemester),
                <>
                <AlertBanner semester={selectedSemester} onActivate={handleActivate} />
                <div className="Dashboard__semesterActions">
                {selectedSemester && !selectedSemester.active && (
                    <button onClick={handleActivateClick}>Activate Class</button>
                )}
                {selectedSemester && selectedSemester.active && (
                    <button onClick={handleDeactivateClick}>Deactivate Class</button>
                )}
            </div>
                </>
            )}
            <div className="Dashboard__classroomDetailsWrapper">
                <UserGroupCard label="Professors" number={1} />
                <UserGroupCard label="TAs" number={12} />
                <UserGroupCard label="Students" number={38} />
            </div>
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