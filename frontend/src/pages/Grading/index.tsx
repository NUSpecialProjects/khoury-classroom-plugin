import { FaChevronRight, FaChevronDown } from "react-icons/fa";
import React, { useState } from "react";

import { Table, TableRow, TableCell } from "@/components/Table";
import { IGradingAssignmentRow } from "./types.ts";

import "./styles.css";

const GradingAssignmentRow: React.FC<IGradingAssignmentRow> = ({
  children,
  submissions,
}) => {
  const [collapsed, setCollapsed] = useState(true);

  return (
    <>
      <TableRow onClick={() => setCollapsed(!collapsed)}>
        <TableCell>
          {collapsed ? <FaChevronRight /> : <FaChevronDown />}
        </TableCell>
        {children}
      </TableRow>
      {!collapsed && (
        <Table className="SubmissionTable" style={{ width: "100%" }}>
          <TableRow labelRow>
            <TableCell primary>Student</TableCell>
            <TableCell>Score</TableCell>
          </TableRow>
          {Array.from({ length: 2 }).map((_, i: number) => (
            <TableRow key={i}>
              <TableCell primary>Jane Doe</TableCell>
              <TableCell>-/100</TableCell>
            </TableRow>
          ))}
        </Table>
      )}
    </>
  );
};

const Grading: React.FC = () => {
  return (
    <div className="Grading">
      <h2 style={{ marginBottom: 0 }}>Assignments</h2>
      <Table className="AssignmentsTable">
        <TableRow labelRow>
          <TableCell>
            <FaChevronRight style={{ opacity: 0 }} />
          </TableCell>
          <TableCell primary>Assignment Name</TableCell>
          <TableCell>Status</TableCell>
          <TableCell>Due Date</TableCell>
        </TableRow>
        {Array.from({ length: 2 }).map((_, i: number) => (
          <GradingAssignmentRow key={i} submissions={[]}>
            <TableCell primary>Assignment Name</TableCell>
            <TableCell>Status</TableCell>
            <TableCell>Due Date</TableCell>
          </GradingAssignmentRow>
        ))}
      </Table>
    </div>
  );
};

export default Grading;
