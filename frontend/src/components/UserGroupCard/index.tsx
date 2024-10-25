import React from "react";
import "./styles.css";

interface IUserGroupCardProps {
  label: string;
  number: number;
  onClick?: () => void;
}

const UserGroupCard: React.FC<IUserGroupCardProps> = ({
  label,
  number,
  onClick,
}) => {
  let userIcons = [];

  if (number > 3) {
    // Cap placeholders at 3, add overlap starting from the second icon
    userIcons = [1, 2, 3].map((_, index) => (
      <div
        className={`UserGroupCard__icon ${index > 0 ? "UserGroupCard__icon-overlap" : ""}`}
        key={index}
      />
    ));
  } else {
    // Render placeholders equal to the number, add overlap starting from the second icon
    userIcons = Array.from({ length: number }).map((_, index) => (
      <div
        className={`UserGroupCard__icon ${index > 0 ? "UserGroupCard__icon-overlap" : ""}`}
        key={index}
      />
    ));
  }

  return (
    <div className="UserGroupCard" onClick={() => onClick && onClick()}>
      <div className="UserGroupCard__content">
        <h3 className="UserGroupCard__label">{label}</h3>
        <div className="UserGroupCard__detailsWrapper">
          <div className="UserGroupCard__icons">{userIcons}</div>
          <p className="UserGroupCard__number">{number}</p>
        </div>
      </div>
    </div>
  );
};

export default UserGroupCard;
