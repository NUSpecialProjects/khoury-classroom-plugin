import { CreateAssignment } from "@/api/assignments";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useContext } from "react";

const Settings: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  const handleCreateAssignment = () => {
    // Call the function with the required arguments
    if (selectedClassroom!= null){
      CreateAssignment("NUSpecialProjects", "practicum-take-home", selectedClassroom.id);
    }
    else {
      console.log("Context Error")
    }

  };


  return (
    <div>
      <h1>Settings</h1>
      <div><button onClick={handleCreateAssignment}>Create Assignment</button></div>
    </div>
  );
};

export default Settings;
