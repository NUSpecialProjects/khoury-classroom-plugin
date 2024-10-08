import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import { Assignments, Grading, Settings, Dashboard, LoginStub } from "./pages";
import Layout from "./components/Layout";

import "./global.css";

export function App(): React.JSX.Element {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route path="assignments" element={<Assignments />} />
          <Route path="grading" element={<Grading />} />
          <Route path="settings" element={<Settings />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="oauth/callback" element={<LoginStub />}/>
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
