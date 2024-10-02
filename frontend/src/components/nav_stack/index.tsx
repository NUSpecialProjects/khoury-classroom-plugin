import { IconType } from "react-icons";
import { FaTachometerAlt, FaStream, FaUsers, FaCog } from "react-icons/fa";
import { useNavigate } from "react-router-dom";


interface NavItemProps {
  buttonName: string;
  buttonDest: string;
  IconDest: IconType;
}

const NavItem: React.FC<NavItemProps> = ({buttonName, buttonDest, IconDest}) => {
  const navigate = useNavigate()
  

  const PageRouteFunction = () => {
      navigate(buttonDest)
  }
  
  return (<>
  <button onClick={PageRouteFunction}>
    <IconDest/> {buttonName}
    </button>
  </>)
}



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
      {
        navItems.map(item => {
          return <NavItem buttonName={item.buttonName} buttonDest={item.buttonDest} IconDest={item.IconDest} />
})
      }
    </div>
  ); 
};

export default NavStack;