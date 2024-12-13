import { useContext } from "react";

import { SelectedClassroomContext } from "@/contexts/selectedClassroom";

import BreadcrumbPageHeader from "@/components/PageHeader/BreadcrumbPageHeader";
import { GradeDistBoxPlot } from "@/components/Viz/BoxPlot/GradeDistBoxPlot";

const Assignments: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

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
