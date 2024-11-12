import React, { useContext } from "react";
import { useLocation } from "react-router-dom";
import GenericRolePage from "..";
import LinkGenerator from "@/components/LinkGenerator";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";

const TAListPage: React.FC = () => {
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );
<<<<<<< HEAD
=======
  const location = useLocation();
  const state = location.state as { users: IClassroomUser[] };
>>>>>>> main
  const role_type = "TA";
  const role_label = "Teaching Assistant";
  return (
    <>
<<<<<<< HEAD
      <GenericRolePage role_type={role_type} />
      <div>
        <p>Add {role_type}</p>
        <LinkGenerator role_type={role_type} classroom={selectedClassroom} />
      </div>
=======
      <GenericRolePage role_label={role_label} userList={state.users} />
      <LinkGenerator
        role_type={role_type}
        role_label={role_label}
        classroom={selectedClassroom}
      />
>>>>>>> main
    </>
  );
};

export default TAListPage;
