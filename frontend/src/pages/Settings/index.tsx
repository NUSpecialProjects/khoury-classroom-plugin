import { acceptAssignment } from "@/api/assignments";
import { AuthContext } from "@/contexts/auth";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useContext } from "react";
import PageHeader from "@/components/PageHeader";
import Button from "@/components/Button";

const Settings: React.FC = () => {
  const { logout } = useContext(AuthContext);

  const { selectedClassroom } = useContext(SelectedClassroomContext);

  const stubAssignmentAcceptEntry = () => {
    if (!selectedClassroom) return;

    // Call the function with the required arguments
    acceptAssignment(
      "NUSpecialProjects",
      "practicum-take-home",
      selectedClassroom.id,
      "practicum-take-home"
    );
  };

  return (
    <div>
      <PageHeader pageTitle="Settings"></PageHeader>
      <Button onClick={logout}>Logout</Button>
      <Button variant="secondary" onClick={stubAssignmentAcceptEntry}>
        Accept Example Assignment
      </Button>
    </div>
  );
};

export default Settings;
