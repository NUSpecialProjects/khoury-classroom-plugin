import { logout as logoutApi } from "@/api/auth";
import { useState, createContext, useLayoutEffect, useContext } from "react";
import { SelectedClassroomContext } from "./selectedClassroom";
import { fetchCurrentUser } from "@/api/users";

interface IAuthContext {
  currentUser: IGitHubUser | null;
  isLoggedIn: boolean;
  login: () => void;
  logout: () => void;
}

export const AuthContext = createContext<IAuthContext>({
  currentUser: null,
  isLoggedIn: false,
  login: () => {},
  logout: () => {},
});

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [currentUser, setCurrentUser] = useState<IGitHubUser | null>(null);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [loading, setLoading] = useState(true);
  const { setSelectedClassroom } = useContext(SelectedClassroomContext);

  useLayoutEffect(() => {
    fetchCurrentUser()
      .then((user) => {
        setIsLoggedIn(true);
        setCurrentUser(user);
      })
      .catch((_: unknown) => {
        setIsLoggedIn(false);
        setCurrentUser(null);
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  const login = () => {
    setIsLoggedIn(true);
  };

  const logout = () => {
    logoutApi()
      .then(() => {
        setSelectedClassroom(null);
        setIsLoggedIn(false);
      })
      .catch((_: Error) => {});
  };

  if (loading) {
    return null;
  }

  return (
    <AuthContext.Provider value={{ currentUser, isLoggedIn, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
