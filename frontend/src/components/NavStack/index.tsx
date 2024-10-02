import { IconType } from "react-icons";
import { FaTachometerAlt, FaStream, FaUsers, FaCog } from "react-icons/fa";
import { Link } from "react-router-dom";

interface NavItemProps {
  buttonName: string;
  buttonDest: string;
  IconDest: IconType;
}

const NavItem: React.FC<NavItemProps> = ({ buttonName, buttonDest, IconDest }) => {
  return (
    <Link to={buttonDest} className="nav-link">
      <IconDest /> {buttonName}
    </Link>
  );
};

const NavStack = () => {
  const navItems: NavItemProps[] = [
    {
      buttonName: "Dashboard",
      buttonDest: "/dashboard",
      IconDest: FaTachometerAlt,
    },
    {
      buttonName: "Grading",
      buttonDest: "/grading",
      IconDest: FaStream,
    },
    {
      buttonName: "Assignments",
      buttonDest: "/assignments",
      IconDest: FaUsers,
    },
    {
      buttonName: "Settings",
      buttonDest: "/settings",
      IconDest: FaCog,
    },
  ];

  return (
    <div className="side-banner">
      {navItems.map((item, index) => (
        <NavItem
          key={index}
          buttonName={item.buttonName}
          buttonDest={item.buttonDest}
          IconDest={item.IconDest}
        />
      ))}
    </div>
  );
};

export default NavStack;
