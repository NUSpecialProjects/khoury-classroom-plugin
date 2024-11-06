import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import React, { useContext } from "react";

interface GenericRolePageProps {
  role_type: string;
}

const GenericRolePage: React.FC<GenericRolePageProps> = ({
  role_type,
}: GenericRolePageProps) => {
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );

  return (
    <div>
      <h1>Create Role Token</h1>
      <div>
        <p>
          Users with role [{role_type}] in org [{selectedClassroom?.org_id}]
        </p>
        <p>(actually put the list here)</p>
      </div>
    </div>
  );
};

export default GenericRolePage;
