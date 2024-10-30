import { SelectedSemesterContext } from "@/contexts/selectedSemester";
import React, { useContext } from "react";

interface GenericRolePageProps {
  role_type: string;
}

const GenericRolePage: React.FC<GenericRolePageProps> = ({
  role_type,
}: GenericRolePageProps) => {
  const { selectedSemester } = useContext(SelectedSemesterContext);

  return (
    <div>
      <h1>Create Role Token</h1>
      <div>
        <p>
          Users with role [{role_type}] in org [{selectedSemester?.org_id}]
        </p>
        <p>(actually put the list here)</p>
      </div>
    </div>
  );
};

export default GenericRolePage;
