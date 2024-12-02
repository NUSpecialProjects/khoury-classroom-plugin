import { AuthContext } from "@/contexts/auth";
import { useContext } from "react";
import PageHeader from "@/components/PageHeader";
import Button from "@/components/Button";
import './styles.css';

const Settings: React.FC = () => {
  const { logout } = useContext(AuthContext);

  return (
    <div className="Settings">
      <PageHeader pageTitle="Settings"></PageHeader>
      <Button variant="primary" onClick={logout}>
        Logout
      </Button>
    </div>
  );
};

export default Settings;
