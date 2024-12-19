import React, { useContext, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useClassroomToken } from "@/api/classrooms";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import TokenApplyPage from "../Generic";
import EmptyDataBanner from "@/components/EmptyDataBanner";
import { useClassroomUser } from "@/hooks/useClassroomUser";
import { ClassroomRole } from "@/types/enums";

const ClassroomTokenApply: React.FC = () => {
  const navigate = useNavigate();
  const { selectedClassroom, setSelectedClassroom } = useContext(SelectedClassroomContext);
  const { classroomUser, loading: loadingCurrentClassroomUser } = useClassroomUser(selectedClassroom?.id);

  useEffect(() => {
    if (!loadingCurrentClassroomUser && classroomUser && selectedClassroom) {
      if (classroomUser.classroom_role === ClassroomRole.STUDENT) {
        navigate("/app/classroom/landing", { replace: true });
      } else {
        navigate("/app/dashboard", { replace: true });
      }
    }
  }, [loadingCurrentClassroomUser, classroomUser, selectedClassroom, navigate]);

  return (
    <EmptyDataBanner>
    <TokenApplyPage<IClassroomJoinResponse>
      useTokenFunction={async (token: string) => {
        return await useClassroomToken(token);
      }}
      successCallback={(data: IClassroomJoinResponse) => {
        setSelectedClassroom(data.classroom);
      }}
      loadingMessage="Joining classroom..."
      successMessage={(response: IClassroomJoinResponse) => response.message}
      />
    </EmptyDataBanner>
  );
};

export default ClassroomTokenApply;