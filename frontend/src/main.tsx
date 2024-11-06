import React, { useContext } from "react";
import ReactDOM from "react-dom/client";
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Navigate,
  Outlet,
} from "react-router-dom";

import * as Pages from "./pages";
import Layout from "./components/Layout";
import AuthProvider, { AuthContext } from "./contexts/auth";
import SelectedSemesterProvider from "./contexts/selectedClassroom";

import "./global.css";

// If not logged in, route to login
const PrivateRoute = () => {
  const { isLoggedIn } = useContext(AuthContext);
  // return isLoggedIn ? <Outlet /> : <Navigate to="/" />;

  return isLoggedIn ? <Outlet /> : <Navigate to="/" />;
};

export default function App(): React.JSX.Element {
  return (
    <AuthProvider>
      <SelectedSemesterProvider>
        <Router>
          <Routes>
            {/******* LANDING PAGE & OAUTH CALLBACK *******/}
            <Route path="" element={<Pages.Login />} />
            <Route path="oauth/callback" element={<Pages.Callback />} />

            {/******* APP ROUTES: AUTHENTICATED USER *******/}
            <Route path="/app" element={<PrivateRoute />}>
              {/******* CLASS SELECTION: PRE-APP ACCESS STEP *******/}

              <Route path="classroom">
                <Route path="select" element={<Pages.ClassroomSelectPage />} />
                <Route path="create" element={<Pages.ClassroomCreatePage />} />
              </Route>
              <Route path="organization">
                <Route
                  path="select"
                  element={<Pages.OrganizationSelectPage />}
                />
              </Route>

              {/******* CLASS SELECTED: INNER APP *******/}
              <Route path="" element={<Layout />}>
                <Route path="assignments" element={<Pages.Assignments />} />
                <Route path="assignments/:id" element={<Pages.Assignment />} />
                <Route path="grading" element={<Pages.Grading />} />
                <Route path="settings" element={<Pages.Settings />} />
                <Route path="token/apply" element={<Pages.RoleApply />} />
                <Route path="token/create" element={<Pages.RoleCreation />} />
                <Route path="students" element={<Pages.StudentListPage />} />
                <Route path="tas" element={<Pages.TAListPage />} />
                <Route
                  path="professors"
                  element={<Pages.ProfessorListPage />}
                />
                <Route
                  path="grading/assignment/:assignmentId/student/:studentAssignmentId"
                  element={<Pages.Grader />}
                />
                <Route path="settings" element={<Pages.Settings />} />
                <Route path="dashboard" element={<Pages.Dashboard />} />
              </Route>
            </Route>

            {/******* 404 CATCH ALL *******/}
            <Route path="404" element={<Pages.PageNotFound />} />
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
