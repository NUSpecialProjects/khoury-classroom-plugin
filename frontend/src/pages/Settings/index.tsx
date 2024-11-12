import { AcceptAssignment } from "@/api/assignments";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useContext } from "react";

const Settings: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  const handleCreateAssignment = () => {
    // Call the function with the required arguments
    if (selectedClassroom!= null){
      AcceptAssignment("NUSpecialProjects", "practicum-take-home", selectedClassroom.id);
    }
    else {
      console.log("Context Error")
    }

  };

  
  return (
    <div>
      <h1>Settings</h1>
      <div><button onClick={handleCreateAssignment}>Accept Example Assignment</button></div>
    </div>
  );
};

export default Settings;
