import { SelectedSemesterContext } from "@/contexts/selectedSemester";
import React, { useContext, useState } from "react";
import LinkGenerator from "@/components/LinkGenerator";

const TokenCreatePage: React.FC = () => {
  const [role_type, setRoleType] = useState<string>("Student");
  const { selectedSemester } = useContext(SelectedSemesterContext);

  return (
    <div>
      <h1>Create Role Token</h1>
      <div>
        <p>Select Role Type:</p>
        <button onClick={() => setRoleType("Student")}>Student</button>
        <button onClick={() => setRoleType("TA")}>Teaching Assistant</button>
      </div>

      <LinkGenerator role_type={role_type} semester={selectedSemester} />
    </div>
  );
};

export default TokenCreatePage;
