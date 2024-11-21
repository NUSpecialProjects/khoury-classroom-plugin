import React, { useEffect, useState } from "react";
import "./styles.css";
import { fetchCurrentUser } from "@/api/users";

const UserProfilePic: React.FC = () => {
  const [user, setUser] = useState<IGitHubUser | null>(null);

  useEffect(() => {
    const fetchUser = async () => {
      await fetchCurrentUser()
        .then((user: IGitHubUser | null) => {
          setUser(user);
        })
        .catch((error: unknown) => {
          console.error("Error fetching user data:", error);
        });
    };

    void fetchUser();
  }, []);

  return (
    <div className="User">
      {user ? (
        <img src={user.avatar_url} alt={user.login} className="User__avatar" />
      ) : (
        <div className="User__avatar"> </div>
      )}
    </div>
  );
};

export default UserProfilePic;
