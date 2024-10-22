import { FaTachometerAlt, FaCog } from "react-icons/fa";
import { MdFactCheck, MdEditDocument } from "react-icons/md";
import { Link } from "react-router-dom";

import "./styles.css";

const LeftNav: React.FC = () => {
  const navItems = [
    { name: "Dashboard", dest: "/app/dashboard", Icon: FaTachometerAlt },
    { name: "Grading", dest: "/app/grading", Icon: MdFactCheck },
    { name: "Assignments", dest: "/app/assignments", Icon: MdEditDocument },
    { name: "Settings", dest: "/app/settings", Icon: FaCog },
    { name: "Kenny", dest: "/app/assignments/1", Icon: FaCog },
  ];

  const className = "CS 3200 Database Design";

  return (
    <div className="LeftNav">
      <div className="LeftNav__classNameHeader">{className}</div>
      {navItems.map((item, index) => (
        <Link key={index} to={item.dest} className="LeftNav__link">
          <item.Icon /> {item.name}
        </Link>
      ))}
    </div>
  );
};

export default LeftNav;
