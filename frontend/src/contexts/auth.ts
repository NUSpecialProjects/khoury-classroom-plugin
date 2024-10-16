import { createContext } from "react";

interface AuthContextProps {
  isLoggedIn: boolean;
  login: () => void;
}

// Handle Auth State- Vulnerable to XSS?
export const AuthContext: React.Context<AuthContextProps> =
  createContext<AuthContextProps>({
    isLoggedIn: false,
    login: () => {},
  });
