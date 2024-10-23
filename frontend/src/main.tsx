import React, { useContext, useEffect } from "react";
import ReactDOM from "react-dom/client";

import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
} from "react-router-dom";
import * as Pages from "./pages";
import Layout from "./components/Layout";
import AuthProvider, { AuthContext } from "./contexts/auth";

import "./global.css";
import SelectedSemesterProvider, {
  SelectedSemesterContext,
} from "./contexts/selectedSemester";

// If not logged in, route to login
const PrivateRoute = ({ element }: { element: React.JSX.Element }) => {
  const { isLoggedIn } = useContext(AuthContext);
  return isLoggedIn ? element : <Navigate to="/" />;
};

// An inner-app route that requires user to have selected a classroom
const AppRoute = ({ element }: { element: React.JSX.Element }) => {
  const { selectedSemester } = useContext(SelectedSemesterContext);
  useEffect(() => {
    console.log(selectedSemester);
  }, []);
  return selectedSemester ? element : <Navigate to="/class-selection" />;
};

export default function App(): React.JSX.Element {
  return (
    <AuthProvider>
      <SelectedSemesterProvider>
        <Router>
          <Routes>
            <Route path="" element={<Pages.Login />} />
            <Route path="oauth/callback" element={<Pages.Callback />} />

            <Route
              path="/app/"
              element={
                <PrivateRoute element={<AppRoute element={<Layout />} />} />
              }
            >
              <Route path="assignments" element={<Pages.Assignments />} />
              <Route path="assignments/:id" element={<Pages.Assignment />} />
              <Route path="grading" element={<Pages.Grading />} />
              <Route
                path="grading/assignment/:assignmentId/student/:studentAssignmentId"
                element={<Pages.Grader />}
              />
              <Route path="settings" element={<Pages.Settings />} />
              <Route path="dashboard" element={<Pages.Dashboard />} />
            </Route>

            <Route
              path="class-creation"
              element={<PrivateRoute element={<Pages.SemesterCreation />} />}
            />
            <Route
              path="class-selection"
              element={<PrivateRoute element={<Pages.SemesterSelection />} />}
            />
          </Routes>
        </Router>
      </SelectedSemesterProvider>
    </AuthProvider>
  );
}

// Safely handle the root element -> Enforced by eslint
const rootElement = document.getElementById("root");
if (!rootElement) {
  throw new Error("Root element not found. Unable to render React app.");
}

ReactDOM.createRoot(rootElement).render(<App />);
