import { Outlet } from "react-router-dom";
import LeftNav from "./LeftNav";
import TopNav from "./TopNav";

import "./styles.css";

const Layout: React.FC = () => {
  return (
    <div className="Layout">
      <div className="Layout__left">
        <LeftNav />
      </div>

      <div className="Layout__right">
        <div className="Layout__top">
          <TopNav />
        </div>
        <div className="Layout__content">
          <Outlet />
        </div>
      </div>

    </div>

  );
};

export default Layout;
