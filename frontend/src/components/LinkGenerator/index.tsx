import { postClassroomToken } from "@/api/classrooms";
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
  const [duration, setDuration] = useState<number | undefined>(10080); // Default to 7 days

  const expirationOptions = [
    { label: "1 hour", value: 60 },
    { label: "6 hours", value: 360 },
    { label: "12 hours", value: 720 },
    { label: "1 day", value: 1440 },
    { label: "7 days", value: 10080 },
    { label: "Never", value: undefined },
  ];

  const handleCreateToken = async () => {
    if (!classroom) {
      setMessage("No classroom selected");
      return;
    }
    await postClassroomToken(classroom.id, role_type, duration)
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
      <select 
        value={duration === undefined ? "" : duration}
        onChange={(e) => setDuration(e.target.value === "" ? undefined : Number(e.target.value))}
      >
        {expirationOptions.map((option) => (
          <option key={option.label} value={option.value ?? ""}>
            {option.label}
          </option>
        ))}
      </select>
      <button onClick={handleCreateToken} disabled={!classroom}>
        Create {role_type} Link
      </button>
      {message && <p>{message}</p>}
    </div>
  );
};

export default LinkGenerator;
