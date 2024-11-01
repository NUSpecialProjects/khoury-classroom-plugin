
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ClipLoader } from "react-spinners";
import useUrlParameter from "@/hooks/useUrlParameter";

const TokenApplyPage: React.FC = () => {
  const inputToken = useUrlParameter("token", "/app/role/apply");
  const [message, setMessage] = useState<string>("Loading...");
  const [loading, setLoading] = useState<boolean>(true);
  const navigate = useNavigate();

  useEffect(() => {
    if (inputToken) {
      handleUseToken();
    }
  }, [inputToken]);

  const handleUseToken = async () => {
    try {
      setMessage("Applying role...");
      const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
      const response = await fetch(`${base_url}/github/role-token/use`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ token: inputToken }),
      });

      if (response.ok) {
        const data = await response.json();
        console.log("Token used:", data);
        setMessage(data.message + " Redirecting...");
        navigate("/app/dashboard", { replace: true }); //TODO: this will redirect to whatever their last selected semester is, but maybe should redirect ALWAYS to this role's semester
      } else {
        setMessage("Failed to use token: " + response.statusText);
      }
    } catch (error) {
      console.error("Error using token:", error);
      setMessage("Error using token: " + error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
        flexDirection: "column",
      }}
    >
      <ClipLoader size={50} color={"#123abc"} loading={loading} />
      <p>{message}</p>
    </div>
  );
};

export default TokenApplyPage;

// import React, { useEffect, useState } from "react";
// import { useNavigate } from "react-router-dom";
// import { ClipLoader } from "react-spinners";

// const TokenApplyPage: React.FC = () => {
//   const [inputToken, setInputToken] = useState<string>("");
//   const [message, setMessage] = useState<string>("Loading...");
//   const [loading, setLoading] = useState<boolean>(true);
//   const navigate = useNavigate();

//   useEffect(() => {
//     const params = new URLSearchParams(location.search);
//     const token = params.get("token");
//     if (token) {
//       setInputToken(token);
//       setMessage("Token received!");
//     } else {
//       setMessage("No token received. Please click the link again.");
//       setLoading(false);
//     }
//   }, [location.search]);

//   useEffect(() => {
//     if (inputToken) {
//       handleUseToken();
//     }
//   }, [inputToken]);

//   const handleUseToken = async () => {
//     try {
//       setMessage("Applying role...");
//       const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
//       const response = await fetch(`${base_url}/github/role-token/use`, {
//         method: "POST",
//         credentials: "include",
//         headers: {
//           "Content-Type": "application/json",
//         },
//         body: JSON.stringify({ token: inputToken }),
//       });

//       if (response.ok) {
//         const data = await response.json();
//         console.log("Token used:", data);
//         setMessage(data.message + " Redirecting...");
//         navigate("/app/dashboard", { replace: true }); //TODO: this will redirect to whatever their last selected semester is, but maybe should redirect ALWAYS to this role's semester
//       } else {
//         setMessage("Failed to use token: " + response.statusText);
//       }
//     } catch (error) {
//       console.error("Error using token:", error);
//       setMessage("Error using token: " + error);
//     } finally {
//       setLoading(false);
//     }
//   };

//   return (
//     <div
//       style={{
//         display: "flex",
//         justifyContent: "center",
//         alignItems: "center",
//         height: "100vh",
//         flexDirection: "column",
//       }}
//     >
//       <ClipLoader size={50} color={"#123abc"} loading={loading} />
//       <p>{message}</p>
//     </div>
//   );
// };

// export default TokenApplyPage;
