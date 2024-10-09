import React, { createContext, useContext, useState } from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";

import { Assignments, Grading, Settings, Dashboard, LoginStub, Splash } from "./pages";
import Layout from "./components/Layout";

import "./global.css";


// Handle Auth State- Vulnerable to XSS?
const AuthContext = createContext<{ isLoggedIn: boolean; login: () => void }>({
  isLoggedIn: false,
  login: () => { },
});

//If not logged in, route to login
const PrivateRoute = ({ element }: { element: JSX.Element }) => {
  const { isLoggedIn } = useContext(AuthContext);
  return isLoggedIn ? element : <Navigate to="" />;
};



export function App(): React.JSX.Element {

  //Handle loggedin state
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  const login = () => setIsLoggedIn(true);



  return (
    <AuthContext.Provider value={{ isLoggedIn, login }}>
      <Router>
        <Routes>
          <Route path="" element={<Splash />} />
          <Route path="/app/" element={<Layout />}>
            <Route path="assignments" element={<PrivateRoute element={<Assignments />} />} />
            <Route path="grading" element={<PrivateRoute element={<Grading />} />} />
            <Route path="settings" element={<PrivateRoute element={<Settings />} />} />
            <Route path="dashboard" element={<PrivateRoute element={<Dashboard />} />} />
            <Route path="oauth/callback" element={<LoginStub />} />
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
