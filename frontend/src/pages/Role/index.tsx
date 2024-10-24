// import useSelectedSemester from "@/contexts/useSelectedSemester";
// import React, { useEffect, useState } from "react";
// import { useNavigate } from "react-router-dom";

// const RoleTokenPage: React.FC = () => {
//     const [createdToken, setCreatedToken] = useState<string | null>(null);
//     const [inputToken, setInputToken] = useState<string>("");
//     const [message, setMessage] = useState<string>("");
//     const [role_type, setRoleType] = useState<string>("Student");

//     const { selectedSemester } = useSelectedSemester();
//     const navigate = useNavigate();


//      useEffect(() => {
//         const params = new URLSearchParams(location.search);
//         const token = params.get("token");
//         if (token) {
//             setInputToken(token);
//             navigate("/app/role-token", { replace: true });
//         }
//     }, [location.search]);


//     const handleCreateToken = async () => {
//         try {
//             const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
//             const response = await fetch(
//                 `${base_url}/github/role-token/create`,
//                 {
//                 method: "POST",
//                 credentials: "include",
//                 headers: {
//                     "Content-Type": "application/json",
//                 },
//                 body: JSON.stringify({
//                     semester: selectedSemester,
//                     role_type: role_type,
//                 }),
//             });

//             if (response.ok) {
//                 const data = await response.json();
//                 setCreatedToken(data.token);
//                 const url = "http://localhost:3000/role-token?token=" + data.token;
//                 setMessage("Link created! " + url);
//             } else {
//                 setMessage("Failed to create token");
//             }
//         } catch (error) {
//             console.error("Error creating token:", error);
//             setMessage("Error creating token");
//         }
//     };

//     const handleUseToken = async () => {
//         try {
//             const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
//             const response = await fetch(
//                 `${base_url}/github/role-token/use`,
//                 {
//                 method: "POST",
//                 credentials: "include",
//                 headers: {
//                     "Content-Type": "application/json",
//                 },
//                 body: JSON.stringify({ token: inputToken }),
//             });

//             if (response.ok) {
//                 const data = await response.json();
//                 setMessage(data.message);
//             } else {
//                 setMessage("Failed to use token");
//             }
//         } catch (error) {
//             console.error("Error using token:", error);
//             setMessage("Error using token");
//         }
//     };

//     return (
//         <div>
//             <h1>Role Token Management</h1>
//             <div>
//                 <p>Select Role Type:</p>
//                 <button onClick={() => setRoleType("Student")}>Student</button>
//                 <button onClick={() => setRoleType("TA")}>Teaching Assistant</button>
//             </div>
            
//             <button onClick={handleCreateToken}>Create {role_type} Token</button>
//             {createdToken && (
//                 <div>
//                     <p>Created Token: {createdToken}</p>
//                 </div>
//             )}
//             <div>
//                 <input
//                     type="text"
//                     value={inputToken}
//                     onChange={(e) => setInputToken(e.target.value)}
//                     placeholder="Enter token"
//                 />
//                 <button onClick={handleUseToken}>Use Token</button>
//             </div>
//             {message && <p>{message}</p>}
//         </div>
//     );
// };

// export default RoleTokenPage;