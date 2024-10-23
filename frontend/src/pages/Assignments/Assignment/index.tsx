import { FaChevronLeft } from "react-icons/fa";

import { Table, TableRow, TableCell } from "@/components/Table/index.tsx";
import Button from "@/components/Button";

import "./styles.css";
import { useLocation } from "react-router-dom";
import { useEffect, useState } from "react";


const Assignment: React.FC = () => {
    const location = useLocation();
    const assignment: IAssignment = location.state.assignment || {}
    console.log("In Assignment Screen", assignment)
    const [studentAssignment, setstudentAssignment] = useState<IAssignment[]>([]);
    

    useEffect(() => {
        console.log("In Assignment Screen", assignment)
        const SyncStudentAssignments = async (sem: ISemester, assignment: IAssignment) => {
            try {
                if (sem.classroom_id) {
                    const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                    const result = await fetch(`${base_url}/${sem.classroom_id}/${assignment.assignment_classroom_id}`, {
                        method: 'GET',
                        credentials: 'include',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                    });

                    if (!result.ok) {
                        throw new Error('Network response was not ok');
                    }

                    const data: IStudentAssignment[] = (await result.json() as IStudentAssignment[])

                    // Need to actually store info
                }

            } catch (error: unknown) {
                console.error('Error fetching assignments:', error);
            }
        };

    }, []);



    return (
        <div className="Assignment">
            <div className="Assignment__head">
                <div className="Assignment__title">
                    <FaChevronLeft />
                    <h2>{assignment.name}</h2>
                </div>
                <div className="Assignment__dates">
                    <span>Due Date: {assignment.main_due_date ?
                        assignment.main_due_date.toString()
                        : "N/A"}</span>
                </div>
            </div>

            <div className="Assignment__externalButtons">
                <Button href="">View in Github Classroom</Button>
                <Button href="">View Starter Code</Button>
                <Button href="">View Rubric</Button>
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
        </div>
    );
};

export default Assignment;