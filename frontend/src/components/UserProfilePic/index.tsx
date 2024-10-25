import React, { useEffect, useState } from "react";
import "./styles.css";
import { fetchCurrentUser } from "@/api/users";
import { IoLogoOctocat } from "react-icons/io";

const UserProfilePic: React.FC = () => {
  const [user, setUser] = useState<IGitHubUser | null>(null);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const currentUser = await fetchCurrentUser() as IGitHubUser;
        setUser(currentUser);
      } catch (error) {
        console.error("Error fetching user data:", error);
      }
    };

    void fetchUser();
  }, []);

  return (
    <div className="User">
      {user? (
        <img src={user.avatar_url} alt={user.login} className="User__avatar" />
      ) : (
        // <div className="UserGroupCard__iconUser__avatar"> </div>
        <div>
          <IoLogoOctocat />
        </div>
      )}
    </div>
  );
};

export default UserProfilePic;
