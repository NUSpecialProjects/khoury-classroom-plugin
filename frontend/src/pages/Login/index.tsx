import { FiGithub, FiX } from "react-icons/fi";
import "./styles.css";
import { Navigate, useLocation, useNavigate } from "react-router-dom";
import { useContext, useEffect, useMemo, useState } from "react";
import ErrorMessage from "@/components/Error";
import { getCallbackURL } from "@/api/login";

import { AuthContext } from "@/contexts/auth";

const Login: React.FC = () => {
  const { isLoggedIn } = useContext(AuthContext);

  const publicUrl: string = import.meta.env.BASE_URL as string;

  const location = useLocation();
  const navigate = useNavigate();
  const queryParams = useMemo(
    () => new URLSearchParams(location.search),
    [location.search]
  );
  const errorFromQuery = queryParams.get("error");
  const [error, setError] = useState<string | null>(errorFromQuery);
  const [callbackURL, setCallbackURL] = useState<string | null>(null);

  useEffect(() => {
    const fetchCallbackURL = async () => {
      try {
        const url = await getCallbackURL();
        setCallbackURL(url);
      } catch (error) {
        console.error("Error fetching callback URL:", error);
      }
    };
    fetchCallbackURL();
  }, []);

  useEffect(() => {
    if (errorFromQuery) {
      queryParams.delete("error");
      setError(errorFromQuery);
      navigate({ search: queryParams.toString() }, { replace: true });
    }
  }, [errorFromQuery, navigate, queryParams]);

  return isLoggedIn ? (
    <Navigate to="/app/dashboard" />
  ) : (
    <div className="LandingPage">
      <div className="LogoBar">
        <FiGithub className="Icon" />
        <FiX className="Icon" />
        <img
          src={`${publicUrl}icons/Northeastern_LVX.svg.png`}
          className="Logo"
        />
      </div>
      <div className="LandingTitle">FonteMarks</div>
      {callbackURL && (
        <a className="SignInLink" href={callbackURL}>
          Log In With GitHub
        </a>
      )}

      {error && <ErrorMessage message={error} />}
    </div>
  );
};

export default Login;
