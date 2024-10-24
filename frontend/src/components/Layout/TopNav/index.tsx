import React from "react";
import UserProfilePic from "../../UserProfilePic";

import "./styles.css";

const TopNav: React.FC = () => {
  return (
    <div className="TopNav">
      <div className="TopNav__right">
        <UserProfilePic />
      </div>
    </div>
  );
};

export default TopNav;
