import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import React, { useContext, useEffect, useState } from "react";

import {
  Table,
  TableRow,
  TableCell,
  TableDiv,
} from "@/components/Table/index.tsx";
import { SelectedSemesterContext } from "@/contexts/selectedClassroom";
import { getAssignments } from "@/api/assignments";
import { getStudentAssignments } from "@/api/student_assignments";
import { formatDate } from "@/utils/date";

import "./styles.css";

const GradingAssignmentRow: React.FC<IGradingAssignmentRow> = ({
  assignmentId,
  children,
}) => {
  const [collapsed, setCollapsed] = useState(true);
  const [studentAssignments, setStudentAssignments] = useState<
    IStudentAssignment[]
  >([]);
  const { selectedClassroom: selectedSemester } = useContext(SelectedSemesterContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (!selectedSemester) return;
    getStudentAssignments(selectedSemester.classroom_id, assignmentId)
      .then((studentAssignments) => {
        console.log(studentAssignments);
        setStudentAssignments(studentAssignments);
      })
      .catch((err: unknown) => {
        console.error("Error fetching student assignments:", err);
      });
  }, []);

  return (
    <>
      <TableRow
        className={!collapsed ? "TableRow--expanded" : undefined}
        onClick={() => {
          setCollapsed(!collapsed);
        }}
      >
        <TableCell>
          {collapsed ? <FaChevronRight /> : <FaChevronDown />}
        </TableCell>
        {children}
      </TableRow>
      {!collapsed && (
        <TableDiv>
          <Table cols={2} className="SubmissionTable">
            <TableRow style={{ borderTop: "none" }}>
              <TableCell>Student</TableCell>
              <TableCell>Score</TableCell>
            </TableRow>
            {studentAssignments &&
              studentAssignments.map((studentAssignment, i: number) => (
                <TableRow
                  key={i}
                  onClick={() => {
                    navigate(`assignment/${assignmentId}/student/${i + 1}`);
                  }}
                >
                  <TableCell>{studentAssignment.student_gh_username}</TableCell>
                  <TableCell>-/100</TableCell>
                </TableRow>
              ))}
          </Table>
        </TableDiv>
      )}
    </>
  );
};

const Grading: React.FC = () => {
  const [assignments, setAssignments] = useState<IAssignment[]>([]);
  const { selectedClassroom: selectedSemester } = useContext(SelectedSemesterContext);
  useEffect(() => {
    if (!selectedSemester) return;
    getAssignments(selectedSemester.classroom_id)
      .then((assignments) => {
        setAssignments(assignments);
      })
      .catch((err: unknown) => {
        console.error("Error fetching assignments:", err);
      });
  }, []);

  return (
    <div className="Grading">
      <h2 style={{ marginBottom: 0 }}>Assignments</h2>
      <Table cols={4} primaryCol={1} className="AssignmentsTable">
        <TableRow style={{ borderTop: "none" }}>
          <TableCell></TableCell>
          <TableCell>Assignment Name</TableCell>
          <TableCell>Assigned Date</TableCell>
          <TableCell>Due Date</TableCell>
        </TableRow>
        {assignments.map((assignment, i: number) => (
          <GradingAssignmentRow key={i} assignmentId={i + 1}>
            <TableCell>{assignment.name}</TableCell>
            <TableCell>{formatDate(assignment.inserted_date)}</TableCell>
            <TableCell>{formatDate(assignment.main_due_date)}</TableCell>
          </GradingAssignmentRow>
        ))}
      </Table>
    </div>
  );
};

export default Grading;
