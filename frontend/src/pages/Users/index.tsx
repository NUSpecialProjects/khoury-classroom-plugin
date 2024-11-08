import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import React, { useContext } from "react";

interface GenericRolePageProps {
  role_label: string;
  userList: IClassroomUser[];
}

const GenericRolePage: React.FC<GenericRolePageProps> = ({
  role_label,
  userList,
}: GenericRolePageProps) => {
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );

  return (
    <div>
      <div>
        <h1>
          {role_label}s in {selectedClassroom?.org_name}
        </h1>
        <div>
          <ul>
            {userList.map((user) => (
              <li key={user.id}>
                {user.first_name} {user.last_name}
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default GenericRolePage;
