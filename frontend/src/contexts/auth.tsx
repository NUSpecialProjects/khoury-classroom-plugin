import { logout as logoutApi } from "@/api/auth";
import { fetchCurrentUser } from "@/api/users";
import { useState, createContext, useLayoutEffect, useContext } from "react";
import { SelectedClassroomContext } from "./selectedClassroom";

interface IAuthContext {
  isLoggedIn: boolean;
  login: () => void;
  logout: () => void;
}

export const AuthContext = createContext<IAuthContext>({
  isLoggedIn: false,
  login: () => {},
  logout: () => {},
});

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [loading, setLoading] = useState(true);
  const { setSelectedClassroom } = useContext(SelectedClassroomContext);

  useLayoutEffect(() => {
    fetchCurrentUser()
      .then((user: IGitHubUser | null) => {
        if (user) {
          setIsLoggedIn(true);
        } else {
          setIsLoggedIn(false);
        }
      })
      .catch((_: unknown) => {
        setIsLoggedIn(false);
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
      .catch((_: Error) => {
      });
  };

  if (loading) {
    return null;
  }

  return (
    <AuthContext.Provider value={{ isLoggedIn, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
