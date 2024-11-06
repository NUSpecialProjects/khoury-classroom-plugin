import { CreateAssignment } from "@/api/assignments";

const Settings: React.FC = () => {
  return (
    <div>
      <h1>Settings</h1>
      <div><button onClick={CreateAssignment}>Create Assignment</button></div>
    </div>
  );
};

export default Settings;
