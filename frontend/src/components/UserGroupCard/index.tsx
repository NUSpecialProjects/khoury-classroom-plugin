import React, { useEffect, useState } from "react";
import "./styles.css";
import { fetchUsersWithRole } from "@/api/users";

interface IUserGroupCardProps {
  label: string;
  role_type: string;
  semester: ISemester;
  onClick?: () => void;
}

const UserGroupCard: React.FC<IUserGroupCardProps> = ({
  label,
  role_type,
  semester,
  onClick,
}) => {
  const [numUsers, setNumUsers] = useState<number>(0);

  useEffect(() => {
    const getUsers = async () => {
      try {
        const users = await fetchUsersWithRole(role_type, semester);
        setNumUsers(users.length);
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };

    void getUsers();
  }, [role_type, semester]);

  let userIcons = [];

  if (numUsers > 3) {
    // Cap placeholders at 3, add overlap starting from the second icon
    userIcons = [1, 2, 3].map((_, index) => (
      <div
        className={`UserGroupCard__icon ${index > 0 ? "UserGroupCard__icon-overlap" : ""}`}
        key={index}
      />
    ));
  } else {
    // Render placeholders equal to the number, add overlap starting from the second icon
    userIcons = Array.from({ length: numUsers }).map((_, index) => (
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
          <p className="UserGroupCard__number">{numUsers}</p>
        </div>
      </div>
    </div>
  );
};

export default UserGroupCard;