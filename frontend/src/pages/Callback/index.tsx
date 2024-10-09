import { AuthContext } from "@/main";
import { useContext, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

const Callback: React.FC = () => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const navigate = useNavigate();
  const { login } = useContext(AuthContext);

  useEffect(()=> {
    //if code, good, else, route to home
    if (code) {
        console.log("Code: " + code)
        const sendCode = async () => {
            try {
                const response = await fetch(`${import.meta.env.VITE_PUBLIC_API_DOMAIN}/login`, {
                    method: "POST",
                    credentials: 'include',
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body :JSON.stringify({code}),
                });

                if (!response.ok){
                    //Navgate back to login page
                    navigate("/")
                }
                else {
                  login()
                  // This is a child component. Update this component to take in the login handler, and call it in this success case
                  navigate("/app/dashboard")
                }
                console.log(response)
            }
            catch (err: unknown){
                //navigate to login page
                navigate("/")
            }
        }
        sendCode()
    }
    else {
        navigate("/")
    }
  })



  return (
    <div>
      <p>Loading...</p>
    </div>
  );
};



export default Callback;
