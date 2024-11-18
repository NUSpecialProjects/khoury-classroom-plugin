import { sendOrganizationInvitesToRequestedUsers, sendOrganizationInviteToUser, revokeOrganizationInvite, removeUserFromClassroom } from "@/api/classrooms";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { ClassroomRole, ClassroomUserStatus } from "@/types/users";
import React, { useContext } from "react";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { Table, TableCell, TableRow } from "@/components/Table";
import EmptyDataBanner from "@/components/EmptyDataBanner";
import './styles.css';
import Button from "@/components/Button";

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

  const removeUserFromList = (userId: number) => {
    userList = userList.filter(user => user.id !== userId);
  };

  const addUserToList = (user: IClassroomUser) => {
    userList = [...userList, user];
  };

  const handleInviteAll = async () => {
    await sendOrganizationInvitesToRequestedUsers(selectedClassroom!.id, role_type)
      .then((data: IClassroomInvitedUsersListResponse) => {
        userList = [...data.invited_users, ...data.requested_users];
      })
      .catch((error) => {
        console.error("Error inviting all users:", error);
      });
  };

  const handleInviteUser = async (userId: number) => {
    await sendOrganizationInviteToUser(selectedClassroom!.id, role_type, userId)
      .then((data: IClassroomUserResponse) => {
        removeUserFromList(userId);
        addUserToList(data.user);
      })
      .catch((error) => {
        console.error("Error inviting user:", error);
      });
  };

  const handleRevokeInvite = async (userId: number) => {
    await revokeOrganizationInvite(selectedClassroom!.id, userId)
      .then((_) => {
        removeUserFromList(userId);
      })
      .catch((error) => {
        console.error("Error revoking invite:", error);
      });
  };

  const handleRemoveUser = async (userId: number) => {
    await removeUserFromClassroom(selectedClassroom!.id, userId)
      .then(() => {
        removeUserFromList(userId);
      })
      .catch((error) => {
        console.error("Error removing user:", error);
      });
  };

  const getActionButton = (user: IClassroomUser) => {
    switch (user.status) {
      case ClassroomUserStatus.ACTIVE:
        return <Button size="small" onClick={() => handleRemoveUser(user.id)}>Remove User</Button>;
      case ClassroomUserStatus.ORG_INVITED:
        return <Button size="small" onClick={() => handleRevokeInvite(user.id)}>Revoke Invitation</Button>;
      case ClassroomUserStatus.REQUESTED:
        return <Button size="small" onClick={() => handleInviteUser(user.id)}>Invite User</Button>;
      default:
        return null;
    }
  };

  return (
    <div>
      <SubPageHeader pageTitle={role_label + `s`} chevronLink="/app/dashboard/"></SubPageHeader>
      
      {userList.filter(user => user.status === ClassroomUserStatus.REQUESTED).length > 0 && (
        <div className="Users__inviteAllWrapper">
          <button onClick={handleInviteAll}>Invite All Requested Users</button>
        </div>
      )}

      <div className="Users__tableWrapper">
        <Table cols={3}>
          <TableRow style={{ borderTop: "none" }}>
            <TableCell>{role_label} Name</TableCell>
            <TableCell>Status</TableCell>
            <TableCell>Actions</TableCell>
          </TableRow>
          {userList.length > 0 ? (
            userList.map((user, i) => (
              <TableRow key={i}>
                <TableCell>{user.first_name} {user.last_name}</TableCell>
                <TableCell>{user.status}</TableCell>
                <TableCell>{getActionButton(user)}</TableCell>
              </TableRow>
            ))
          ) : (
            <EmptyDataBanner>
              <p>There are currently no {role_label}s in this classroom.</p>
            </EmptyDataBanner>
          )}
        </Table>
      </div>
    </div>
  );
};

export default GenericRolePage;
