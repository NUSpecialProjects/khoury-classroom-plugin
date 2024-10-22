import React, { useContext, useState } from "react";
import ReactDOM from "react-dom/client";

import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import * as Pages from "./pages";
import Layout from "./components/Layout";
import { AuthContext } from "./contexts/auth";

import "./global.css";

//If not logged in, route to login
const PrivateRoute = ({ element }: { element: React.JSX.Element }) => {
  const { isLoggedIn } = useContext(AuthContext);
  return isLoggedIn ? element : <Navigate to="/" />;
};

export default function App(): React.JSX.Element {
  //Handle loggedin state
  const [isLoggedIn, setIsLoggedIn] = useState(
    import.meta.env.MODE == "development"
  );

  const login = () => {
    setIsLoggedIn(true);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, login }}>
      <Router>
        <Routes>
          <Route path="" element={<Pages.Login />} />
          <Route path="oauth/callback" element={<Pages.Callback />} />
          <Route path="/app/" element={<PrivateRoute element={<Layout />} />}>
            <Route path="assignments" element={<Pages.Assignments />} />
            <Route path="assignments/:id" element={<Pages.Assignment />} />
            <Route path="grading" element={<Pages.Grading />} />
            <Route path="settings" element={<Pages.Settings />} />
            <Route path="dashboard" element={<Pages.Dashboard />} />
          </Route>
          <Route
            path="semester-creation"
            element={<Pages.SemesterCreation />}
          />
          <Route
            path="semester-selection"
            element={<Pages.SemesterSelection />}
          />
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
