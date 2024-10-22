/* eslint-disable */

import "./styles.css";
import UserGroupCard from "@/components/UserGroupCard";
import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import { Link } from "react-router-dom";
import useSelectedSemester from "@/contexts/useSelectedSemester";
import AlertBanner from "@/components/Banner/AlertBanner";
import { Semester } from "@/types/semester";
import { activateSemester, deactivateSemester } from "@/api/semesters";
import { useState } from "react";
import ErrorMessage from "@/components/Error";

const Dashboard: React.FC = () => {
    const {selectedSemester, setSelectedSemester} = useSelectedSemester();
    const [error, setError] = useState<string | null>(null);

    const handleActivate = async (newSemester: Semester) => {
        setSelectedSemester(newSemester);
    }

    const handleActivateClick = async () => {
        if (selectedSemester) {
            try {
                const newSemester = await activateSemester(selectedSemester.org_id, selectedSemester.classroom_id);
                handleActivate(newSemester);
                setError(null);
            } catch (err) {
                setError("Failed to activate the semester. Please try again.");
            }
        }
    };

    const handleDeactivateClick = async () => {
        if (selectedSemester) {
            try {
                const newSemester = await deactivateSemester(selectedSemester.org_id, selectedSemester.classroom_id);
                handleActivate(newSemester);
                setError(null);
            } catch (err) {
                setError("Failed to deactivate the semester. Please try again.");
            }
        }
    };

    return (
        <div className="Dashboard">
            {selectedSemester && (
                <>
            <h1>{selectedSemester.org_name + " - " + selectedSemester.classroom_name}</h1>
                    <AlertBanner semester={selectedSemester} onActivate={handleActivate} />
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
            <div>
                <p>Temporary Classroom Settings</p>
                        {selectedSemester && !selectedSemester.active && (
                            <button onClick={handleActivateClick}>Activate Class</button>
                        )}
                        {selectedSemester && selectedSemester.active && (
                            <button onClick={handleDeactivateClick}>Deactivate Class</button>
                        )}
                    </div>
                    {error && (
                        <ErrorMessage message={error} />
                    )}
            </>
            )}
        </div>
    )
}

export default Dashboard;