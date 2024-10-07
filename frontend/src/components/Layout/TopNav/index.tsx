import { Link } from "react-router-dom";

import "./styles.css";

const TopNav: React.FC = () => {
  return (
    <div className="TopNav">
      <div className="TopNav__left">Khoury Classroom</div>
      <div className="TopNav__right">Some Menu Here</div>
    </div>
  );
};

export default TopNav;
