import { getCurrentUser } from "@/api/auth";
import { useState, createContext, useLayoutEffect } from "react";

interface IAuthContext {
  isLoggedIn: boolean;
  login: () => void;
}

// Handle Auth State- Vulnerable to XSS?
export const AuthContext: React.Context<IAuthContext> =
  createContext<IAuthContext>({
    isLoggedIn: false,
    login: () => {},
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
      .catch((err: unknown) => {
        console.log("Error fetching current user: ", err);
      })
      .finally(() => {
        setLoading(false);
      });
  }, []);

  const login = () => {
    setIsLoggedIn(true);
  };

  return (
    !loading && (
      <AuthContext.Provider value={{ isLoggedIn, login }}>
        {children}
      </AuthContext.Provider>
    )
  );
};

export default AuthProvider;
