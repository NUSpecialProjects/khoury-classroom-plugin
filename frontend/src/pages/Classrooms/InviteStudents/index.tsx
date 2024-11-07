import Panel from "@/components/Panel";
import Button from "@/components/Button";
import { useNavigate, Link } from "react-router-dom";

import "../styles.css";

const InviteStudents: React.FC = () => {
const navigate = useNavigate();

    return (
        <Panel title="Add Students" logo={true}>
            <div className="Invite">
                <div className="Invite__ContentWrapper">
                    <div className="Invite__TextWrapper">
                        <h2>Use the link below to invite students</h2>
                        <div>To add students to your classroom, invite them using this link!</div>
                    </div>
                </div>
                <div className="Invite__ButtonWrapper">
                    <Button variant="primary" onClick={() => navigate("/app/classroom/success")}>Continue</Button>
                </div>
            </div>
        </Panel>
    );
};

export default InviteStudents;