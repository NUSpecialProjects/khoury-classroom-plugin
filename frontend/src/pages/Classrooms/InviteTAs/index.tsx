import Panel from "@/components/Panel";
import Button from "@/components/Button";
import { useNavigate, Link } from "react-router-dom";

import "../styles.css";

const InviteTAs: React.FC = () => {
const navigate = useNavigate();

    return (
        <Panel title="Add Teaching Assistants" logo={true}>
            <div className="Invite">
                <div className="Invite__ContentWrapper">
                    <div className="Invite__TextWrapper">
                        <h2>Use the link below to invite TAs to your Classroom</h2>
                        <div>To add TA’s to your classroom, invite them using this link!</div>
                    </div>
                </div>
                <div className="ButtonWrapper">
                    <Button variant="primary" onClick={() => navigate("/app/classroom/invite-students")}>Continue</Button>
                </div>
            </div>
        </Panel>
    );
};

export default InviteTAs;