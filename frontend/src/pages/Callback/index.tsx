import { AuthContext } from "@/contexts/auth";
import { useContext, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import ClipLoader from "react-spinners/ClipLoader";
import "./styles.css";
import { sendCode } from "@/api/auth";

const Callback: React.FC = () => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const navigate = useNavigate();
  const { login } = useContext(AuthContext);

  const handleSuccessfulLogin = () => {
    const redirectUrl = localStorage.getItem("redirectAfterLogin");
    login(); // set the user's login status to true
    if (redirectUrl) {
      localStorage.removeItem("redirectAfterLogin");
      navigate(redirectUrl); // redirect to the page that was requested before login
    } else {
      navigate("/app/organization/select"); // default redirect after login
    }
  };

  useEffect(() => {
    //if code, good, else, route to home
    if (code) {
      sendCode(code)
        .then((response) => {
          if (!response.ok) {
            // Navigate back to login page
            navigate("/");
            return;
          } else {
            //Successful login. Handle redirect
            handleSuccessfulLogin();
          }
        })
        .catch((err: unknown) => {
          // Navigate back to login page with an error message attached
          navigate(
            `/?error=${encodeURIComponent("An error occurred while logging in. Please try again.")}`
          );
          console.log("Error Occurred: ", err);
          return;
        });
    } else {
      navigate("/");
    }
  });

  return (
    <div className="callback-container">
      <ClipLoader size={50} color={"#123abc"} loading={true} />
      <p>Logging in...</p>
    </div>
  );
};

export default Callback;
