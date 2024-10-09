import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import * as Pages from "./pages";
import Layout from "./components/Layout";

import "./global.css";

export function App(): React.JSX.Element {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route path="assignments" element={<Pages.Assignments />} />
          <Route path="assignments/:id" element={<Pages.AssignmentDetails />} />
          <Route path="grading" element={<Pages.Grading />} />
          <Route path="settings" element={<Pages.Settings />} />
          <Route path="dashboard" element={<Pages.Dashboard />} />
        </Route>
      </Routes>
    </Router>
  );
}

// Safely handle the root element -> Enforced by eslint
const rootElement = document.getElementById("root");
if (!rootElement) {
  throw new Error("Root element not found. Unable to render React app.");
}

ReactDOM.createRoot(rootElement).render(<App />);
