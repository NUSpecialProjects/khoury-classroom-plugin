/* eslint-disable */

import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./styles.css";
import { Semester, UserSemestersResponse } from "@/types/semester";
import { getUserSemesters } from "@/api/semesters";
import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import {
    Table,
    TableRow,
    TableCell,
    TableDiv,
} from "@/components/Table/index.tsx";
import useSelectedSemester from "@/contexts/useSelectedSemester";

const SemesterSelection: React.FC = () => {
    const [semestersByOrg, setSemestersByOrg] = useState<{ [key: number]: Semester[] }>({});
    const [collapsed, setCollapsed] = useState<{ [key: number]: boolean }>({});
    const [loading, setLoading] = useState(true);
    const {setSelectedSemester} = useSelectedSemester();

    const navigate = useNavigate();

    useEffect(() => {
        const fetchSemesters = async () => {
            try {
                setSelectedSemester(null);
                const data: UserSemestersResponse = await getUserSemesters();
                const groupedSemesters: { [key: number]: Semester[] } = {};
                const initialCollapsedState: { [key: number]: boolean } = {};

                data.active_semesters.concat(data.inactive_semesters).forEach((semester) => {
                    if (!groupedSemesters[semester.org_id]) {
                        groupedSemesters[semester.org_id] = [];
                        initialCollapsedState[semester.org_id] = false; // Set default to uncollapsed
                    }
                    groupedSemesters[semester.org_id].push(semester);
                });

                setSemestersByOrg(groupedSemesters);
                setCollapsed(initialCollapsedState);
            } catch (error) {
                console.error("Error fetching semesters:", error);
            } finally {
                setLoading(false);
            }
        };

        void fetchSemesters();
    }, []);

    const handleSemesterSelect = (semester: Semester) => {
        console.log("Selected semester:", semester);
        setSelectedSemester(semester);
        navigate(`/app/dashboard`);
    };

    const toggleCollapse = (orgId: number) => {
        setCollapsed((prev) => ({ ...prev, [orgId]: !prev[orgId] }));
    };

    const hasSemesters = Object.keys(semestersByOrg).length > 0;

    return (
        <div className="SemesterSelection">
            <h1>Select a Semester</h1>
            {loading ? (
                <p>Loading...</p>
            ) : hasSemesters ? (
            <Table cols={5} primaryCol={1} className="SemestersTable">
                {Object.keys(semestersByOrg).map((orgId) => (
                    <React.Fragment key={orgId}>
                        <TableRow className="HeaderRow"
                            style={{ borderTop: "none" }}>
                            <TableCell></TableCell>
                            <TableCell>Course Name</TableCell>
                            <TableCell>Status</TableCell>
                            <TableCell>Organization ID</TableCell>
                            <TableCell>Action</TableCell>
                        </TableRow>
                        <TableRow
                            className={`ChildRow ${!collapsed[Number(orgId)] ? "TableRow--expanded" : ""}`}
                        >
                            <TableCell className="fixed-width-button"
                                onClick={() => toggleCollapse(Number(orgId))}>
                                {collapsed[Number(orgId)] ? <FaChevronRight /> : <FaChevronDown />}
                            </TableCell>
                            <TableCell>{`${semestersByOrg[Number(orgId)][0].name.split(':')[0]}`}</TableCell>
                            <TableCell>{semestersByOrg[Number(orgId)].some((semester) => (
                                semester.active)) ? "Active" : "Inactive"}
                            </TableCell>
                            <TableCell
                                className="fixed-width-id-column">
                                {orgId}
                            </TableCell>
                            <TableCell>
                                <button
                                    className="fixed-width-button"
                                    onClick={() => toggleCollapse(Number(orgId))}>
                                    {collapsed[Number(orgId)] ? "Expand" : "Hide"}
                                </button>
                            </TableCell>
                        </TableRow>
                        {!collapsed[Number(orgId)] && (
                            <TableDiv>
                                <Table cols={4} primaryCol={0} className="DetailsTable SubTable">
                                    <TableRow className="HeaderRow" style={{ borderTop: "none" }}>
                                        {/* <TableCell></TableCell> */}
                                        <TableCell>Semester Name</TableCell>
                                        <TableCell>Status</TableCell>
                                        <TableCell>Classroom ID</TableCell>
                                        <TableCell>Action</TableCell>
                                    </TableRow>
                                    {semestersByOrg[Number(orgId)].map((semester) => (
                                        <TableRow key={semester.id} className="SemesterRow">
                                            {/* <TableCell className="fixed-width-button"></TableCell> */}
                                            <TableCell>{semester.name.split(":")[1]}</TableCell>
                                            <TableCell>{semester.active ? "Active" : "Inactive"}</TableCell>
                                            <TableCell
                                                className="fixed-width-id-column">
                                                {semester.classroom_id}
                                            </TableCell>
                                            <TableCell>
                                                <button
                                                    className="fixed-width-button"
                                                    onClick={() => handleSemesterSelect(semester)}>
                                                    View
                                                </button>
                                            </TableCell>
                                        </TableRow>
                                    ))}
                                </Table>
                            </TableDiv>
                        )}
                    </React.Fragment>
                ))}
            </Table>
            ) : (
                <div>
                    <p>You have no semesters.</p>
                </div>
            )}
            <button onClick={() => navigate("/semester-creation")}>Create a Semester</button>
        </div>
    );
};

export default SemesterSelection;