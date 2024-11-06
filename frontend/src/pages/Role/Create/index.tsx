import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import React, { useContext, useState } from "react";
import LinkGenerator from "@/components/LinkGenerator";

const TokenCreatePage: React.FC = () => {
  const [role_type, setRoleType] = useState<string>("Student");
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );

  return (
    <div>
      <h1>Create Role Token</h1>
      <div>
        <p>Select Role Type:</p>
        <button onClick={() => setRoleType("Student")}>Student</button>
        <button onClick={() => setRoleType("TA")}>Teaching Assistant</button>
      </div>

      <LinkGenerator role_type={role_type} classroom={selectedClassroom} />
    </div>
  );
};

export default TokenCreatePage;
