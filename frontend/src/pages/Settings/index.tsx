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
    // Call the function with the required arguments
    if (selectedClassroom != null){
      acceptAssignment("NUSpecialProjects", "practicum-take-home", selectedClassroom.id, "practicum-take-home");
    }
    else {
      console.log("Context Error")
    }

  };




  return (
    <div>
      <PageHeader pageTitle="Settings"></PageHeader>
      <Button variant="primary" onClick={logout}>Logout</Button>
      <Button variant="secondary" onClick={stubAssignmentAcceptEntry}>Accept Example Assignment</Button>
    </div>
  );
};

export default Settings;
