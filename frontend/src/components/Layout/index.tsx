import { Outlet, Navigate } from "react-router-dom";
import { useContext } from "react";
import SimpleBar from "simplebar-react";
import "simplebar/dist/simplebar.min.css";

import LeftNav from "./LeftNav";
import TopNav from "./TopNav";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";

import "./styles.css";

const Layout: React.FC = () => {
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );
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
