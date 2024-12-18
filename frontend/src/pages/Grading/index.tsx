import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import React, { useContext, useEffect, useState } from "react";

import {
  Table,
  TableRow,
  TableCell,
  TableDiv,
} from "@/components/Table/index.tsx";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getAssignments } from "@/api/assignments";
import { getStudentWorks } from "@/api/student_works";
import { formatDateTime } from "@/utils/date";
import PageHeader from "@/components/PageHeader";

import "./styles.css";
import { StudentWorkState } from "@/types/enums";

interface IGradingAssignmentRow extends React.HTMLProps<HTMLDivElement> {
  assignmentId: number;
}

const GradingAssignmentRow: React.FC<IGradingAssignmentRow> = ({
  assignmentId,
  children,
}) => {
  const [collapsed, setCollapsed] = useState(true);
  const [studentAssignments, setStudentAssignments] = useState<IStudentWork[]>(
    []
  );
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );
  const navigate = useNavigate();

  useEffect(() => {
    if (!selectedClassroom) return;
    getStudentWorks(selectedClassroom.id, assignmentId)
      .then((studentAssignments) => {
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
          <Table cols={3} className="SubmissionTable">
            <TableRow style={{ borderTop: "none" }}>
              <TableCell>Student</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Score</TableCell>
            </TableRow>
            {studentAssignments &&
              studentAssignments
                .filter(
                  (studentAssignment) =>
                    studentAssignment.work_state !==
                    StudentWorkState.NOT_ACCEPTED
                )
                .map((studentAssignment, i: number) => (
                  <TableRow
                    key={i}
                    onClick={() => {
                    navigate(
                      `assignment/${assignmentId}/student/${studentAssignment.student_work_id}`
                    );
                  }}
                >
                  <TableCell>
                    {studentAssignment.contributors.map(c => `${c.full_name}`).join(", ")}
                  </TableCell>
                  <TableCell>{studentAssignment.work_state}</TableCell>
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
  const [assignments, setAssignments] = useState<IAssignmentOutline[]>([]);
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );
  useEffect(() => {
    if (!selectedClassroom) return;
    getAssignments(selectedClassroom.id)
      .then((assignments) => {
        setAssignments(assignments);
      })
      .catch((err: unknown) => {
        console.error("Error fetching assignments:", err);
      });
  }, []);

  return (
    <div className="Grading">
      <PageHeader pageTitle="Assignments"></PageHeader>
      <Table cols={4} primaryCol={1} className="AssignmentsTable">
        <TableRow style={{ borderTop: "none" }}>
          <TableCell></TableCell>
          <TableCell>Assignment Name</TableCell>
          <TableCell>Assigned Date</TableCell>
          <TableCell>Due Date</TableCell>
        </TableRow>
        {assignments.map((assignment, i: number) => (
          <GradingAssignmentRow key={i} assignmentId={assignment.id}>
            <TableCell>{assignment.name}</TableCell>
            <TableCell>{formatDateTime(assignment.created_at)}</TableCell>
            <TableCell>{formatDateTime(assignment.main_due_date)}</TableCell>
          </GradingAssignmentRow>
        ))}
      </Table>
    </div>
  );
};

export default Grading;
