import React, { useEffect, useState } from "react";
import "./styles.css";

interface GitHubUser {
    login: string;
    id: number;
    node_id: string;
    avatar_url: string;
    url: string;
    name: string | null;
    email: string | null;
}

const UserProfilePic: React.FC = () => {
    const [user, setUser] = useState<GitHubUser | null>(null);

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
                const response = await fetch(`${base_url}/github/user`, {
                    method: "GET",
                    credentials: 'include',
                    headers: {
                        "Content-Type": "application/json",
                    },
                });
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
                const data: GitHubUser = await response.json() as GitHubUser;
                setUser(data);
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        };

        void fetchUser();
    }, []);

    if (!user) {
        return <div>Loading...</div>;
    }

    return (
        <div className="User">
            <img src={user.avatar_url} alt={user.login} className="User__avatar" />
        </div>
    );
};

export default UserProfilePic;