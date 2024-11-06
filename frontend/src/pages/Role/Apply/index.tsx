import React, { useContext, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ClipLoader } from "react-spinners";
import useUrlParameter from "@/hooks/useUrlParameter";
import { useClassroomToken } from "@/api/classrooms";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";

const TokenApplyPage: React.FC = () => {
  const inputToken = useUrlParameter("token", "/app/token/apply");
  const [message, setMessage] = useState<string>("Loading...");
  const [loading, setLoading] = useState<boolean>(true);
  const navigate = useNavigate();
  const {setSelectedClassroom} = useContext(SelectedClassroomContext);

  useEffect(() => {
    if (inputToken) {
      handleUseToken();
    }
  }, [inputToken]);

  const handleUseToken = async () => {
    setMessage("Applying role...");
    await useClassroomToken(inputToken)
      .then((data: IClassroomJoinResponse) => {
        setMessage(data.message + " Redirecting...");
        setSelectedClassroom(data.classroom);
        navigate("/app/dashboard", { replace: true });
      })
      .catch((error) => {
        setMessage("Error using token: " + error);
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

export default TokenApplyPage;
