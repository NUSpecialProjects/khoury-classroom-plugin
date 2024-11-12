import React, { useEffect, useState } from "react";
import "./styles.css";
import { fetchUser } from "@/api/users";

interface IUserGroupCardProps {
  label: string;
<<<<<<< HEAD
  role_type: string;
  classroom: IClassroom;
=======
  givenUsersList: IClassroomUser[];
>>>>>>> main
  onClick?: () => void;
}

const UserGroupCard: React.FC<IUserGroupCardProps> = ({
  label,
<<<<<<< HEAD
  role_type,
  classroom,
=======
  givenUsersList,
>>>>>>> main
  onClick,
}) => {
  const [userMap, setUserMap] = useState<Map<IClassroomUser, IGitHubUser>>(
    new Map()
  );
  useEffect(() => {
<<<<<<< HEAD
    const getUsers = async () => {
      try {
        const users = await fetchUsersWithRole(role_type, classroom);
        setNumUsers(users.length);
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    };

    void getUsers();
  }, [role_type, classroom]);
=======
    const loadGitHubUsers = async () => {
      if (givenUsersList) {
        const newMap = new Map();
        await Promise.all(
          givenUsersList.map(async (classroomUser) => {
            await fetchUser(classroomUser.github_username)
              .then((userResponse: IGitHubUserResponse) => {
                newMap.set(classroomUser, userResponse.user);
              })
              .catch((_) => {
                // do nothing
              });
          })
        );
        setUserMap(newMap);
      }
    };

    void loadGitHubUsers();
  }, [givenUsersList]);
>>>>>>> main

  let userIcons: React.ReactNode[] = [];
  const MAX_USERS_TO_SHOW = 3;

  if (givenUsersList && givenUsersList.length > 0) {
    const usersToShow = givenUsersList.slice(0, MAX_USERS_TO_SHOW);
    userIcons = usersToShow.map((classroomUser, index) => {
      const githubUser = userMap.get(classroomUser);
      return (
        <div key={index}>
          {!githubUser ? (
            <div className="UserGroupCard__icon-placeholder" />
          ) : (
            <img
              className={`UserGroupCard__icon ${index > 0 ? "UserGroupCard__icon-overlap" : ""}`}
              src={githubUser.avatar_url}
              alt={`${githubUser.login}'s avatar`}
            />
          )}
        </div>
      );
    });
  }

  return (
    <div className="UserGroupCard" onClick={() => onClick && onClick()}>
      <div className="UserGroupCard__content">
        <h3 className="UserGroupCard__label">{label}</h3>
        <div className="UserGroupCard__detailsWrapper">
          <div className="UserGroupCard__icons">{userIcons}</div>
          <p className="UserGroupCard__number">{givenUsersList.length}</p>
        </div>
      </div>
    </div>
  );
};

export default UserGroupCard;
