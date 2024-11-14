import React, { useEffect, useState } from "react";
// import { useNavigate } from "react-router-dom";
import { ClipLoader } from "react-spinners";
import useUrlParameter from "@/hooks/useUrlParameter";
import { useAssignmentToken } from "@/api/assignments";

const AcceptAssignmentPage: React.FC = () => {
  const inputToken = useUrlParameter("token", "/app/assignments/accept");
  const [message, setMessage] = useState<string>("Loading...");
  const [loading, setLoading] = useState<boolean>(true);
//   const navigate = useNavigate();

  useEffect(() => {
    if (inputToken) {
      handleAcceptAssignment();
    }
  }, [inputToken]);

  const handleAcceptAssignment = async () => {
    setMessage("Accepting assignment...");
    await useAssignmentToken(inputToken)
      .then((_: IAssignmentAcceptResponse) => {
        setMessage("Assignment accepted successfully!");
        // navigate("/app/assignments", { replace: true });
      })
      .catch((error) => {
        setMessage("Error accepting assignment: " + error);
      })
      .finally(() => {
        setLoading(false);
      });
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

export default AcceptAssignmentPage;
