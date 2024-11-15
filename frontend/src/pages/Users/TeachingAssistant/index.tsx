import React from "react";
import { useLocation } from "react-router-dom";
import GenericRolePage from "..";

const TAListPage: React.FC = () => {
  const location = useLocation();
  const state = location.state as { users: IClassroomUser[] };
  const role_type = "TA";
  const role_label = "Teaching Assistant";
  return (
    <>
      <GenericRolePage role_label={role_label} role_type={role_type} userList={state.users} />
    </>
  );
};

export default TAListPage;
