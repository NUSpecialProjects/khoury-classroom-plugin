import { AuthContext } from "@/contexts/auth";
import { useContext } from "react";
import PageHeader from "@/components/PageHeader";
import Button from "@/components/Button";

const Settings: React.FC = () => {
  const { logout } = useContext(AuthContext);

  return (
    <div>
      <PageHeader pageTitle="Settings"></PageHeader>
      <Button variant="primary" onClick={logout}>Logout</Button>
    </div>
  );
};

export default Settings;
