import { Banner, NavStack } from "../components"
import { Link, Outlet } from "react-router-dom"

const Layout: React.FC = () => {



  const clientId: string = import.meta.env.VITE_GITHUB_CLIENT_ID as string;

    return (
    <div className="app">
    <Banner />
    <div className="body">
      <Link to={`https://github.com/login/oauth/authorize?client_id=${clientId}&scope=repo,read:org,classroom&allow_signup=false`}>Sign In</Link>
        <NavStack/>
        <Outlet />
      </div>     
    </div>
    )
  }

  export default Layout;