import { Outlet, Navigate } from "react-router-dom";
import { useContext } from "react";
import SimpleBar from "simplebar-react";
import "simplebar-react/dist/simplebar.min.css";

import LeftNav from "./LeftNav";
import TopNav from "./TopNav";
import { SelectedClassroomContext } from "@/contexts/selectedClassroom";

import "./styles.css";
import { FaTachometerAlt } from "react-icons/fa";
import { MdEditDocument } from "react-icons/md";
import { MdFactCheck } from "react-icons/md";
import { FaGear } from "react-icons/fa6";
import { useClassroomUser } from "@/hooks/useClassroomUser";
import { ClassroomRole } from "@/types/users";

const Layout: React.FC = () => {
  const { selectedClassroom: selectedClassroom } = useContext(
    SelectedClassroomContext
  );
  const { classroomUser } = useClassroomUser(selectedClassroom?.id);

  const baseNavItems = [
    { name: "Dashboard", dest: "/app/dashboard", Icon: FaTachometerAlt },
    { name: "Grading", dest: "/app/grading", Icon: MdEditDocument },
  ];

  const professorNavItems = [
    { name: "Rubrics", dest: "/app/rubrics", Icon: MdFactCheck }
  ];

  const settingsNavItem = [
    { name: "Settings", dest: "/app/settings", Icon: FaGear }
  ];

  const navItems = [
    ...baseNavItems,
    ...(classroomUser?.classroom_role === ClassroomRole.PROFESSOR ? professorNavItems : []),
    ...settingsNavItem
  ];

  return selectedClassroom ? (
    <div className="Layout">
      <div className="Layout__left">
        <LeftNav navItems={navItems} />
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
