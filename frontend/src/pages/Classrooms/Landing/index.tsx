import Panel from "@/components/Panel";
import "./styles.css";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useContext } from "react";

const Landing = () => {

  const { selectedClassroom } = useContext(SelectedClassroomContext);

  return (
    <Panel title={`Successfully joined ${selectedClassroom?.name}`}>
      <p>You may now close this page.</p>
    </Panel>
  );
};

export default Landing;
