import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import React, { useContext, useState } from "react";
import { useQuery } from "@tanstack/react-query";

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
import LoadingSpinner from "@/components/LoadingSpinner";
import EmptyDataBanner from "@/components/EmptyDataBanner";

import "./styles.css";
import Button from "@/components/Button";
import { MdAdd } from "react-icons/md";

interface IGradingAssignmentRow extends React.HTMLProps<HTMLDivElement> {
  assignmentId: number;
}

const GradingAssignmentRow: React.FC<IGradingAssignmentRow> = ({
  assignmentId,
  children,
}) => {
  const [collapsed, setCollapsed] = useState(true);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const navigate = useNavigate();

  const { data: studentAssignments } = useQuery({
    queryKey: ['studentWorks', selectedClassroom?.id, assignmentId],
    queryFn: () => getStudentWorks(selectedClassroom!.id, assignmentId),
    enabled: !!selectedClassroom && !collapsed,
  });

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
                    navigate(
                      `assignment/${assignmentId}/student/${studentAssignment.student_work_id}`
                    );
                  }}
                >
                  <TableCell>
                    {studentAssignment.contributors.join(", ")}
                  </TableCell>
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
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  const { data: assignments, isLoading, error } = useQuery({
    queryKey: ['assignments', selectedClassroom?.id],
    queryFn: () => getAssignments(selectedClassroom!.id),
    enabled: !!selectedClassroom,
  });

  return (
    <div className="Grading">
      <PageHeader pageTitle="Assignments"></PageHeader>
      {isLoading ? (
        <EmptyDataBanner>
          <LoadingSpinner />
        </EmptyDataBanner>
      ) : error ? (
        <EmptyDataBanner>
          Error loading assignments: {error instanceof Error ? error.message : "Unknown error"}
        </EmptyDataBanner>
      ) : assignments && assignments.length > 0 ? (
        <Table cols={4} primaryCol={1} className="AssignmentsTable">
          <TableRow style={{ borderTop: "none" }}>
            <TableCell></TableCell>
            <TableCell>Assignment Name</TableCell>
            <TableCell>Assigned Date</TableCell>
            <TableCell>Due Date</TableCell>
          </TableRow>
          {assignments && assignments.map((assignment, i: number) => (
            <GradingAssignmentRow key={i} assignmentId={assignment.id}>
              <TableCell>{assignment.name}</TableCell>
              <TableCell>{formatDateTime(assignment.created_at)}</TableCell>
              <TableCell>{formatDateTime(assignment.main_due_date)}</TableCell>
            </GradingAssignmentRow>
          ))}
        </Table>
      ) : (
        <EmptyDataBanner>
          <div className="emptyDataBannerMessage">
            No assignments found.
          </div>
          <Button
              variant="secondary"
              size="small"
              href={`/app/assignments/create?org_name=${selectedClassroom?.org_name}`}
              >
              <MdAdd className="icon" /> Create Assignment
            </Button>
        </EmptyDataBanner>
      )}
    </div>
  );
};

export default Grading;
