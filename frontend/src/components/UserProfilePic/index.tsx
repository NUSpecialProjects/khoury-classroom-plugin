import React, { useEffect, useState } from "react";
import "./styles.css";
import { fetchCurrentUser } from "@/api/users";
import { IoLogoOctocat } from "react-icons/io";

const UserProfilePic: React.FC = () => {
  const [user, setUser] = useState<IGitHubUser | null>(null);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const currentUser = (await fetchCurrentUser()) as IGitHubUser;
        setUser(currentUser);
      } catch (_) {
        // do nothing
      }
    };

    void fetchUser();
  }, []);

  return (
    <div className="User">
      {user && <img src={user.avatar_url} alt={user.login} className="User__avatar" />}
    </div>
  );
};

export default UserProfilePic;
