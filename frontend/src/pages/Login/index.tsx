import { FiGithub, FiX } from "react-icons/fi";
import "./styles.css";
import { Navigate, useLocation, useNavigate } from "react-router-dom";
import { useContext, useEffect, useMemo, useState } from "react";
import ErrorMessage from "@/components/Error";
import { getCallbackURL } from "@/api/login";

import { AuthContext } from "@/contexts/auth";

enum LoginStatus {
  LOADING = "LOADING",
  CALLBACK_ERRORED = "CALLBACK ERRORED",
  LOGIN_ERRORED = "LOGIN ERRORED",
  READY = "READY",
}

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
  const [status, setStatus] = useState(LoginStatus.LOADING);
  const [error, setError] = useState<string | null>(errorFromQuery);
  const [callbackURL, setCallbackURL] = useState<string | null>(null);

  useEffect(() => {
    setStatus(LoginStatus.LOADING);
    setError(null);
    const fetchCallbackURL = async () => {
      try {
        const url = await getCallbackURL();
        if (!url) {
          throw new Error("Callback URL is empty");
        }
        setCallbackURL(url);
        setStatus(LoginStatus.READY);
        setError(null);
      } catch (_) {
        setStatus(LoginStatus.LOGIN_ERRORED);
        setError("Error occurred while communicating with the server");
        console.log("Error occurred while fetching callback URL");
      }
    };
    fetchCallbackURL();
  }, []);

  useEffect(() => {
    if (errorFromQuery) {
      queryParams.delete("error");
      setError(errorFromQuery);
      navigate({ search: queryParams.toString() }, { replace: true });
      setStatus(LoginStatus.CALLBACK_ERRORED);
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

      {callbackURL && status !== LoginStatus.LOADING && (
        <a className="SignInLink" href={callbackURL}>
          Log In With GitHub
        </a>
      )}

      {status === LoginStatus.LOADING && (
        <div className="LoadingMessage">Loading...</div>
      )}

      {error && <ErrorMessage message={error} />}
    </div>
  );
};

export default Login;
