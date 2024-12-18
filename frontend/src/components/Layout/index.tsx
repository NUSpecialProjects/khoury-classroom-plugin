import { Outlet, Navigate, useNavigate } from "react-router-dom";
import { useContext, useEffect } from "react";
import ClipLoader from "react-spinners/ClipLoader";
import SimpleBar from "simplebar-react";
import "simplebar-react/dist/simplebar.min.css";

import { SelectedClassroomContext } from "@/contexts/selectedClassroom";
import { useClassroomUser } from "@/hooks/useClassroomUser";

import LeftNav from "./LeftNav";
import TopNav from "./TopNav";
import Button from "@/components/Button";

import "./styles.css";

const Layout: React.FC = () => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const {
    classroomUser,
    error: classroomUserError,
    loading: loadingCurrentClassroomUser,
  } = useClassroomUser(selectedClassroom?.id);

  const navigate = useNavigate();

  useEffect(() => {
    if (
      !loadingCurrentClassroomUser &&
      (classroomUserError || !classroomUser)
    ) {
      navigate(`/app/organization/select`);
    }
  }, [
    loadingCurrentClassroomUser,
    classroomUserError,
    classroomUser,
    selectedClassroom?.org_id,
    navigate,
  ]);

  if (loadingCurrentClassroomUser) {
    return (
      <div className="Layout__loading">
        <ClipLoader size={50} color={"#123abc"} loading={true} />
      </div>
    );
  }

  if (classroomUser?.classroom_role === "STUDENT") {
    return (
      <div className="Dashboard__unauthorized">
        <h2>Access Denied</h2>
        <p>
          You do not have permission to view the classroom management dashboard.
        </p>
        <p>Please contact your professor if you believe this is an error.</p>
        <Button
          variant="primary"
          href={`/app/classroom/select?org_id=${selectedClassroom?.org_id}`}
        >
          Return to Classroom Selection
        </Button>
      </div>
    );
  }

  return selectedClassroom ? (
    <div className="Layout">
      <div className="Layout__left">
        <LeftNav />
      </div>

      <SimpleBar className="Layout__right">
        <div className="Layout__top">
          <TopNav />
        </div>
        <div className="Layout__content">
          <Outlet />
        </div>
      </SimpleBar>
    </div>
  ) : (
    <Navigate to="/app/organization/select" />
  );
};

export default Layout;
