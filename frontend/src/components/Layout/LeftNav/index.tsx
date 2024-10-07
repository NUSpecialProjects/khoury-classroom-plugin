import { FaTachometerAlt, FaStream, FaUsers, FaCog } from "react-icons/fa";
import { Link } from "react-router-dom";

const LeftNav: React.FC = () => {
  const navItems = [
    { name: "Dashboard", dest: "/dashboard", Icon: FaTachometerAlt },
    { name: "Grading", dest: "/grading", Icon: FaStream },
    { name: "Assignments", dest: "/assignments", Icon: FaUsers },
    { name: "Settings", dest: "/settings", Icon: FaCog },
  ];

  return (
    <div className="LeftNav">
      {navItems.map((item, index) => (
        <Link key={index} to={item.dest} className="nav-link">
          <item.Icon /> {item.name}
        </Link>
      ))}
    </div>
  );
};

export default LeftNav;
