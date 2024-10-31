import React, { useContext } from "react";
import GenericRolePage from "..";
import LinkGenerator from "@/components/LinkGenerator";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";

const TAListPage: React.FC = () => {
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );
  const role_type = "TA";
  return (
    <>
      <GenericRolePage role_type={role_type} />
      <div>
        <p>Add {role_type}</p>
        <LinkGenerator role_type={role_type} classroom={selectedClassroom} />
      </div>
    </>
  );
};

export default TAListPage;
