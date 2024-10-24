import { useState, createContext } from "react";

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
  //Handle loggedin state
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);

  const login = () => {
    setIsLoggedIn(true);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, login }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
