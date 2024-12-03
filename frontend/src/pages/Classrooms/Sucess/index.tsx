import Panel from "@/components/Panel";
import Button from "@/components/Button";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useContext } from "react";

import "../styles.css";

const Success: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  return (
    <Panel title={selectedClassroom?.name + " Classroom Created!"} logo={true}>
      <div className="ButtonWrapper">
        <Button href="/app/dashboard">View my classroom</Button>
      </div>
    </Panel>
  );
};

export default Success;
