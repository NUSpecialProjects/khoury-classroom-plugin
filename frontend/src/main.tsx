import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'; // Correct imports
import Layout from './Layout';
import { Assignments, Grading, Settings, Dashboard} from './pages';
import "./index.css"

function App(): JSX.Element {
  return (
    <Router basename="/vite-template"> 
      <Routes>
        <Route path="/" element={<Layout />} >
        <Route path="assignments" element={<Assignments/>} />
        <Route path="grading" element={<Grading />} />
        <Route path="settings" element={<Settings />} />
        <Route path="dashboard" element={<Dashboard />} />
        </Route>
      </Routes>
    </Router>
  );
}

ReactDOM.createRoot(document.getElementById("root")!).render(<App />);
