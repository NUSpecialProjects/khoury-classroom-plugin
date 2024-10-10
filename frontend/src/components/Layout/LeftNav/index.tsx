import { FaTachometerAlt, FaStream, FaUsers, FaCog } from "react-icons/fa";
import { Link } from "react-router-dom";

import "./styles.css";

const LeftNav: React.FC = () => {
  const navItems = [
    { name: "Dashboard", dest: "/app/dashboard", Icon: FaTachometerAlt },
    { name: "Grading", dest: "/app/grading", Icon: FaStream },
    { name: "Assignments", dest: "/app/assignments", Icon: FaUsers },
    { name: "Settings", dest: "/app/settings", Icon: FaCog },
    { name: "Kenny", dest: "/app/assignments/1", Icon: FaCog },
  ];

  return (
    <div className="LeftNav">
      {navItems.map((item, index) => (
        <Link key={index} to={item.dest} className="LeftNav__link">
          <item.Icon /> {item.name}
        </Link>
      ))}
    </div>
  );
};

export default LeftNav;
