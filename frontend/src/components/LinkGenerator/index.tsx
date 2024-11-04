import { createToken } from "@/api/users";
import React, { useState } from "react";

interface CreateTokenProps {
  role_type: string;
  classroom: IClassroom | null;
}

const LinkGenerator: React.FC<CreateTokenProps> = ({
  role_type,
  classroom,
}) => {
  const [message, setMessage] = useState<string>("");

  const handleCreateToken = async () => {
    if (!classroom) {
      setMessage("No classroom selected");
      return;
    }
    await createToken(role_type, classroom)
      .then((data: ITokenResponse) => {
        const url = "http://localhost:3000/app/token/apply?token=" + data.token;
        setMessage("Link created! " + url);
        navigator.clipboard.writeText(url);
      })
      .catch((error) => {
        setMessage("Error creating token: " + error);
      });
  };

  return (
    <div>
      <button onClick={handleCreateToken} disabled={!classroom}>
        Create {role_type} Link
      </button>
      {message && <p>{message}</p>}
    </div>
  );
};

export default LinkGenerator;
