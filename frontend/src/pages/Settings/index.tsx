import { AuthContext } from "@/contexts/auth";
import { useContext } from "react";

const Settings: React.FC = () => {
  const { logout } = useContext(AuthContext);
  return (
    <div>
      <h1>Settings</h1>
      <button onClick={logout}>Logout</button>
    </div>
  );
};

export default Settings;
