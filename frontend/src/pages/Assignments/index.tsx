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

        <GradeDistBoxPlot />
      </>
    )
  );
};

export default Assignments;
