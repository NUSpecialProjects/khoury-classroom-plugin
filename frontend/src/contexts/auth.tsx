import { getCurrentUser, logout as logoutApi } from "@/api/auth";
import { useState, createContext, useLayoutEffect } from "react";

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

  useLayoutEffect(() => {
    getCurrentUser()
      .then((ok) => {
        setIsLoggedIn(ok);
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
