import React, { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { useClassroomToken } from "@/api/classrooms";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import TokenApplyPage from "../Generic";
import EmptyDataBanner from "@/components/EmptyDataBanner";

const ClassroomTokenApply: React.FC = () => {
  const navigate = useNavigate();
  const { setSelectedClassroom } = useContext(SelectedClassroomContext);

  return (
    <EmptyDataBanner>
    <TokenApplyPage<IClassroomJoinResponse>
      useTokenFunction={async (token: string) => {
        return await useClassroomToken(token);
      }}
      successCallback={(data: IClassroomJoinResponse) => {
        setSelectedClassroom(data.classroom);
        navigate("/app/dashboard", { replace: true });
      }}
      loadingMessage="Joining classroom..."
      successMessage={(response: IClassroomJoinResponse) => response.message}
      />
    </EmptyDataBanner>
  );
};

export default ClassroomTokenApply;