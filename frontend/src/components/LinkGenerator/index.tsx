import React, { useState } from "react";

interface CreateTokenProps {
    role_type: string;
    semester: ISemester | null;
}

const LinkGenerator: React.FC<CreateTokenProps> = ({ role_type, semester }) => {
    const [message, setMessage] = useState<string>("");

    const handleCreateToken = async () => {
        try {
            const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
            const response = await fetch(
                `${base_url}/github/role-token/create`,
                {
                    method: "POST",
                    credentials: "include",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        semester: semester,
                        role_type: role_type,
                    }),
                }
            );

            if (response.ok) {
                const data = await response.json();
                const url = "http://localhost:3000/app/token/apply?token=" + data.token;
                setMessage("Link created! " + url);
                navigator.clipboard.writeText(url);
            } else {
                setMessage("Failed to create token");
            }
        } catch (error) {
            console.error("Error creating token:", error);
            setMessage("Error creating token");
        }
    };

    return (
        <div>
            <button onClick={handleCreateToken} disabled={!semester || !semester.active}>Create {role_type} Link</button>
            {message && <p>{message}</p>}
        </div>
    );
};

export default LinkGenerator;