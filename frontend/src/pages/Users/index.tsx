import { sendOrganizationInvitesToRequestedUsers, sendOrganizationInviteToUser, revokeOrganizationInvite, removeUserFromClassroom } from "@/api/classrooms";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { ClassroomRole, ClassroomUserStatus } from "@/types/users";
import React, { useContext, useState } from "react";

interface GenericRolePageProps {
  role_label: string;
  role_type: ClassroomRole;
  userList: IClassroomUser[];
}

const GenericRolePage: React.FC<GenericRolePageProps> = ({
  role_label,
  role_type,
  userList,
}: GenericRolePageProps) => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  const [requestedUsers, setRequestedUsers] = useState<IClassroomUser[]>(userList.filter(user => user.status === ClassroomUserStatus.REQUESTED));
  const [invitedUsers, setInvitedUsers] = useState<IClassroomUser[]>(userList.filter(user => user.status === ClassroomUserStatus.ORG_INVITED));
  const [activeUsers, setActiveUsers] = useState<IClassroomUser[]>(userList.filter(user => user.status === ClassroomUserStatus.ACTIVE));

  const handleInviteAll = async () => {
    await sendOrganizationInvitesToRequestedUsers(selectedClassroom!.id, role_type)
      .then((data: IClassroomInvitedUsersListResponse) => {
        const { invited_users, requested_users } = data;
        setInvitedUsers(invited_users);
        setRequestedUsers(requested_users);
      })
      .catch((error) => {
        console.error("Error inviting all users:", error);
      });
  };

  const handleInviteUser = async (userId: number) => {
    await sendOrganizationInviteToUser(selectedClassroom!.id, role_type, userId)
      .then((data: IClassroomUserResponse) => {
        const { user } = data;
        setInvitedUsers([...invitedUsers, user]);
        setRequestedUsers(requestedUsers.filter(user => user.id !== userId));
      })
      .catch((error) => {
        console.error("Error inviting user:", error);
      });
  };

  const handleRevokeInvite = async (userId: number) => {
    await revokeOrganizationInvite(selectedClassroom!.id, userId)
      .then(() => {
        setInvitedUsers(invitedUsers.filter(user => user.id !== userId));
      })
      .catch((error) => {
        console.error("Error revoking invite:", error);
      });
  };

  const handleRemoveUser = async (userId: number) => {
    await removeUserFromClassroom(selectedClassroom!.id, userId)
      .then(() => {
        setActiveUsers(activeUsers.filter(user => user.id !== userId));
      })
      .catch((error) => {
        console.error("Error removing user:", error);
      });
  };

  return (
    <div>
      <h1>{role_label}s in {selectedClassroom?.org_name}</h1>
      <div>
        <h2>Active Users</h2>
        <ul>
          {activeUsers.length > 0 ? (
            activeUsers.map((user) => (
              <li key={user.id}>
                {user.first_name} {user.last_name}
                <button onClick={() => handleRemoveUser(user.id)}>Remove User</button>
              </li>
            ))
          ) : (
            <li>No active users</li>
          )}
        </ul>
      </div>

      <div>
        <h2>Invited Users</h2>
        <ul>
          {invitedUsers.length > 0 ? (
            invitedUsers.map((user) => (
              <li key={user.id}>
                {user.first_name} {user.last_name}
                <button onClick={() => handleRevokeInvite(user.id)}>Revoke Invitation</button>
              </li>
            ))
          ) : (
            <li>No invited users</li>
          )}
        </ul>
      </div>


      <div>
        <div style={{ display: 'flex', alignItems: 'center', gap: '1rem' }}>
          <h2>Requested Users</h2>
          <button onClick={handleInviteAll}>Invite All Requested Users</button>
        </div>
        <ul>
          {requestedUsers.length > 0 ? (
            requestedUsers.map((user) => (
              <li key={user.id}>
                {user.first_name} {user.last_name}
                <button onClick={() => handleInviteUser(user.id)}>Invite User</button>
              </li>
            ))
          ) : (
            <li>No requested users</li>
          )}
        </ul>
      </div>

    </div>
  );
};

export default GenericRolePage;
