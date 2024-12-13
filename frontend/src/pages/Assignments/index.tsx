import { useContext, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { MdAdd } from "react-icons/md";

import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { getAssignments } from "@/api/assignments";
import { formatDateTime } from "@/utils/date";

import BreadcrumbPageHeader from "@/components/PageHeader/BreadcrumbPageHeader";
import { GradeDistBoxPlot } from "@/components/Viz/BoxPlot/GradeDistBoxPlot";
import { Table, TableRow, TableCell } from "@/components/Table";
import Button from "@/components/Button";

const Assignments: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const [assignments, setAssignments] = useState<IAssignmentOutline[]>([]);

  useEffect(() => {
    const fetchAssignments = async (classroom: IClassroom) => {
      if (classroom) {
        getAssignments(classroom.id)
          .then((assignments) => {
            setAssignments(assignments);
          })
          .catch((_: unknown) => {
            // do nothing
          });
      }
    };

    if (selectedClassroom !== null && selectedClassroom !== undefined) {
      fetchAssignments(selectedClassroom).catch((_: unknown) => {
        // do nothing
      });
    }
  }, [selectedClassroom]);

  return (
    selectedClassroom && (
      <>
        <BreadcrumbPageHeader
          pageTitle={selectedClassroom?.org_name}
          breadcrumbItems={[selectedClassroom?.name, "Assignments"]}
        />

        <div className="Dashboard__sectionWrapper">
          <div className="Dashboard__assignmentsHeader">
            <h2 style={{ marginBottom: 0 }}>Assignments</h2>
            <div className="Dashboard__createAssignmentButton">
              <Button
                variant="secondary"
                size="small"
                href={`/app/assignments/create?org_name=${selectedClassroom?.org_name}`}
              >
                <MdAdd className="icon" /> Create Assignment
              </Button>
            </div>
          </div>
          <Table cols={2}>
            <TableRow style={{ borderTop: "none" }}>
              <TableCell>Assignment Name</TableCell>
              <TableCell>Created Date</TableCell>
            </TableRow>
            {assignments.map((assignment, i: number) => (
              <TableRow key={i} className="Assignment__submission">
                <TableCell>
                  <Link
                    to={`/app/assignments/${assignment.id}`}
                    state={{ assignment }}
                    className="Dashboard__assignmentLink"
                  >
                    {assignment.name}
                  </Link>
                </TableCell>
                <TableCell>{formatDateTime(assignment.created_at)}</TableCell>
              </TableRow>
            ))}
          </Table>
        </div>

        <GradeDistBoxPlot />
      </>
    )
  );
};

export default Assignments;
