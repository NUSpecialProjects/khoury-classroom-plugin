<<<<<<< HEAD
import { createToken } from "@/api/users";
=======
import { postClassroomToken } from "@/api/classrooms";
>>>>>>> main
import React, { useState } from "react";

interface CreateTokenProps {
  role_type: string;
<<<<<<< HEAD
=======
  role_label: string;
>>>>>>> main
  classroom: IClassroom | null;
}

const LinkGenerator: React.FC<CreateTokenProps> = ({
  role_type,
<<<<<<< HEAD
=======
  role_label,
>>>>>>> main
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
<<<<<<< HEAD
    await createToken(role_type, classroom)
=======
    await postClassroomToken(classroom.id, role_type, duration)
>>>>>>> main
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
<<<<<<< HEAD
      <button onClick={handleCreateToken} disabled={!classroom}>
        Create {role_type} Link
=======
      <h3>Create {role_label} Invite Link</h3>
      <select
        value={duration === undefined ? "" : duration}
        onChange={(e) =>
          setDuration(
            e.target.value === "" ? 10080 : Number(e.target.value)
          )
        }
      >
        {expirationOptions.map((option) => (
          <option key={option.label} value={option.value ?? ""}>
            {option.label}
          </option>
        ))}
      </select>
      <button onClick={handleCreateToken} disabled={!classroom}>
        Generate {role_label} Link
>>>>>>> main
      </button>
      {message && <p>{message}</p>}
    </div>
  );
};

export default LinkGenerator;
