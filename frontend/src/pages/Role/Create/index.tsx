import { SelectedSemesterContext } from "@/contexts/selectedSemester";
import React, { useContext, useState } from "react";

const TokenCreatePage: React.FC = () => {
    // const [createdToken, setCreatedToken] = useState<string | null>(null);
    const [message, setMessage] = useState<string>("");
    const [role_type, setRoleType] = useState<string>("Student");

  const { selectedSemester } = useContext(SelectedSemesterContext);

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
                    semester: selectedSemester,
                    role_type: role_type,
                }),
            });

            if (response.ok) {
                const data = await response.json();
                // setCreatedToken(data.token);
                const url = "http://localhost:3000/app/token/apply?token=" + data.token;
                setMessage("Link created! " + url);
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
            <h1>Create Role Token</h1>
            <div>
                <p>Select Role Type:</p>
                <button onClick={() => setRoleType("Student")}>Student</button>
                <button onClick={() => setRoleType("TA")}>Teaching Assistant</button>
            </div>
            
            <button onClick={handleCreateToken}>Create {role_type} Token</button>
            {/* {createdToken && (
                <div>
                    <p>Created Token: {createdToken}</p>
                </div>
            )} */}
            {message && <p>{message}</p>}
        </div>
    );
};

export default TokenCreatePage;