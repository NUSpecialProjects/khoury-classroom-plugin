import React from "react";
import GenericRolePage from "..";

const ProfessorListPage: React.FC = () => {
  const role_type = "Professor";
  return (
    <>
      <GenericRolePage role_type={role_type} />
    </>
  );
};

export default ProfessorListPage;
