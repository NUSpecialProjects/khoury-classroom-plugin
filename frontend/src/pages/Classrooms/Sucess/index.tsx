import Panel from "@/components/Panel";
import Button from "@/components/Button";
import { useNavigate } from "react-router-dom";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useContext } from "react";

import "../styles.css";

const Success: React.FC = () => {
    const navigate = useNavigate();
    const { selectedClassroom } = useContext(SelectedClassroomContext);

    return (
        <Panel title={selectedClassroom?.name + " Classroom Created!"} logo={true}>
            <div className="ButtonWrapper">
                <Button variant="primary" onClick={() => navigate("/app/dashboard")}>View my classroom</Button>
            </div>
        </Panel>
    );
};

export default Success;