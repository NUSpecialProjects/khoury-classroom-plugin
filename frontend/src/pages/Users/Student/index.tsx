import React, { useContext } from "react";
import { useLocation } from "react-router-dom";
import GenericRolePage from "..";
import LinkGenerator from "@/components/LinkGenerator";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { ClassroomRole } from "@/types/users";

const StudentListPage: React.FC = () => {
  const location = useLocation();
  const state = location.state as { users: IClassroomUser[] };
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const role_type: ClassroomRole = ClassroomRole.STUDENT;
  const role_label: string = "Student";
  return (
    <>
      <GenericRolePage role_label={role_label} role_type={role_type} userList={state.users} />
      <LinkGenerator
        role_type={role_type}
        role_label={role_label}
        classroom={selectedClassroom}
      />
    </>
  );
};

export default StudentListPage;
