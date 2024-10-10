import { FiGithub, FiX } from "react-icons/fi";
import "./styles.css"


const Login: React.FC = () => {
    const clientId: string = import.meta.env.VITE_GITHUB_CLIENT_ID as string

    return(
        <div className="LandingPage">
            <div className="LogoBar">
            <FiGithub className="Icon"/>
            <FiX className="Icon"/>
            <img src="src/assets/icons/Northeastern_LVX.svg.png" className="Logo" />
            </div>
        <div className="LandingTitle">Khoury Classroom</div>
        <a className="SignInLink" href={`https://github.com/login/oauth/authorize?client_id=${clientId}&scope=repo,read:org,classroom&allow_signup=false`}>
            Log In With GitHub</a>
        </div>
    )
} 

export default Login;