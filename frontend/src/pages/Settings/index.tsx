import { acceptAssignment } from "@/api/assignments";
import { AuthContext } from "@/contexts/auth";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useContext } from "react";

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
      <h1>Settings</h1>
      <button onClick={logout}>Logout</button>
      <button onClick={stubAssignmentAcceptEntry}>Accept Example Assignment</button>
    </div>
  );
};

export default Settings;
