import React, { createContext, useContext, useState } from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import * as Pages from "./pages";
import Layout from "./components/Layout";

import "./global.css";

interface AuthContextProps {
  isLoggedIn: boolean;
  login: () => void;
}

// Handle Auth State- Vulnerable to XSS?
export const AuthContext: React.Context<AuthContextProps> = createContext<AuthContextProps>({
  isLoggedIn: false,
  login: () => {},
});
 

//If not logged in, route to login
const PrivateRoute = ({ element }: { element: React.JSX.Element }) => {
  const { isLoggedIn } = useContext(AuthContext);
  return isLoggedIn ? element : <Navigate to="/" />;
};



export function App(): React.JSX.Element {
  //Handle loggedin state
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  
  const login = () => {
    setIsLoggedIn(true);
  }

  
  return (
    <AuthContext.Provider value={{isLoggedIn, login}}>
    <Router>
      <Routes>
      <Route path="" element={<Pages.Login />}/>
      <Route path="oauth/callback" element={<Pages.Callback />} />
        <Route path="/app/" element={<PrivateRoute element={<Layout />} />}>
            <Route path="assignments" element={<Pages.Assignments />} />
            <Route path="assignments/:id" element={<Pages.AssignmentDetails />} />
            <Route path="grading" element={<Pages.Grading />}/>
            <Route path="settings" element={<Pages.Settings />} />
            <Route path="dashboard" element={<Pages.Dashboard />}/>
        </Route>
      </Routes>
    </Router>
    </AuthContext.Provider>
  );
}

// Safely handle the root element -> Enforced by eslint
const rootElement = document.getElementById("root");
if (!rootElement) {
  throw new Error("Root element not found. Unable to render React app.");
}

ReactDOM.createRoot(rootElement).render(<App />);