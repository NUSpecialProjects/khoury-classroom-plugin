import { useContext } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import Button from "@/components/Button";
import "./styles.css";

const AccessDenied: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  return (
    <div className="AccessDenied">
      <div className="AccessDenied__content">
        <h2>Access Denied</h2>
        <p>
          You do not have permission to view this page.
        </p>
        <p>Please contact your professor if you believe this is an error.</p>
        <Button
          variant="primary"
          href={`/app/classroom/select?org_id=${selectedClassroom?.org_id}`}
        >
          Return to Classroom Selection
        </Button>
      </div>
    </div>
  );
};

export default AccessDenied;
