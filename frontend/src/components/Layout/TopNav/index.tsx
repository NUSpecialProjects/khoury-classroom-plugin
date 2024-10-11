import React from "react";
import { Link } from "react-router-dom";
import UserProfilePic from "../../UserProfilePic";

import "./styles.css";

const TopNav: React.FC = () => {
  return (
    <div className="TopNav">
      <div className="TopNav__left">
        <Link to="/">Khoury Classroom</Link>
      </div>

      <div className="TopNav__right">
        <UserProfilePic />
      </div>
    </div>
  );
};

export default TopNav;