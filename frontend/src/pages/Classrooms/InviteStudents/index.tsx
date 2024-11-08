import Panel from "@/components/Panel";
import Button from "@/components/Button";
import CopyLink from "@/components/CopyLink";
import { useState, useContext, useEffect } from "react";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useNavigate } from "react-router-dom";
import { postClassroomToken } from "@/api/classrooms";


import "../styles.css";

const InviteStudents: React.FC = () => {
const navigate = useNavigate();
const { selectedClassroom } = useContext(SelectedClassroomContext);
const [link, setLink] = useState<string>("");

useEffect(() => {
    const handleCreateToken = async () => {
        if (!selectedClassroom) {
            return;
        }
        await postClassroomToken(selectedClassroom.id, "TA")
            .then((data: ITokenResponse) => {
                const url = "http://localhost:3000/app/token/apply?token=" + data.token;
                setLink(url);
            })
            .catch((error) => {
                console.log("Error creating token: " + error);
            });
    };

    if (selectedClassroom) {
        handleCreateToken();
    }
}, [selectedClassroom])

    return (
        <Panel title="Add Students" logo={true}>
            <div className="Invite">
                <div className="Invite__ContentWrapper">
                    <div className="Invite__TextWrapper">
                        <h2>Use the link below to invite students</h2>
                        <div>{"To add students to your classroom, invite them using this link!"}</div>
                    </div>
                    <CopyLink link={link} name="invite-students"></CopyLink>
                </div>
                <div className="ButtonWrapper">
                    <Button variant="primary" onClick={() => navigate("/app/classroom/success")}>Continue</Button>
                </div>
            </div>
        </Panel>
    );
};

export default InviteStudents;