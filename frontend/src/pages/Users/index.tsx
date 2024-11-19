import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import CopyLink from "@/components/CopyLink";
import { useState, useContext, useEffect } from "react";
import SubPageHeader from "@/components/PageHeader/SubPageHeader";
import { Table, TableCell, TableRow } from "@/components/Table";
import EmptyDataBanner from "@/components/EmptyDataBanner";
import { postClassroomToken } from "@/api/classrooms";
import './styles.css';

interface GenericRolePageProps {
  role_label: string;
  role_type: string;
  userList: IClassroomUser[];
}

const GenericRolePage: React.FC<GenericRolePageProps> = ({
  role_label,
  role_type,
  userList,
}) => {

  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const [link, setLink] = useState<string>("");
  const base_url: string = import.meta.env.VITE_PUBLIC_FRONTEND_DOMAIN as string;
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const handleCreateToken = async () => {
      if (!selectedClassroom) {
        return;
      }
      await postClassroomToken(selectedClassroom.id, role_type)
        .then((data: ITokenResponse) => {
          const url = `${base_url}/app/token/apply?token=${data.token}`;
          setLink(url);
        })
        .catch((_) => {
          setError("Failed to generate invite URL. Please try again.");
        });
    };

    if (selectedClassroom) {
      handleCreateToken();
    }
  }, [selectedClassroom])

  return (
    <div>
      <SubPageHeader pageTitle={role_label + `s`} chevronLink="/app/dashboard/"></SubPageHeader>
      <div className="Users__inviteLinkWrapper">
        <div>
          <h2>Invite {role_label + `s`}</h2>
          <p>Share this link to invite and add students to {selectedClassroom?.name}.</p>
        </div>
        <CopyLink link={link} name="invite-tas"></CopyLink>
        {error && <p className="error">{error}</p>}
      </div>
      <div className="Users__tableWrapper">
        <Table cols={1}>
          <TableRow style={{ borderTop: "none" }}>
            <TableCell>{role_label + ` Name`}</TableCell>
          </TableRow>
          {userList.length >= 1 ? (
            userList.map((user, i) => (
              <TableRow key={i}>
                <TableCell style={{ textDecoration: "none", cursor: "default" }}>{user.first_name} {user.last_name}</TableCell>
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
