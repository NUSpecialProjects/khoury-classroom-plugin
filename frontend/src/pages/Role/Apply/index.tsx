import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { ClipLoader } from "react-spinners";
import useUrlParameter from "@/hooks/useUrlParameter";
import { useRoleToken } from "@/api/users";

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
    setMessage("Applying role...");
    await useRoleToken(inputToken)
      .then((data: IMessageResponse) => {
        setMessage(data.message + " Redirecting...");
        navigate("/app/dashboard", { replace: true }); //TODO: this will redirect to whatever their last selected classroom is, but maybe should redirect ALWAYS to this role's semester
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
