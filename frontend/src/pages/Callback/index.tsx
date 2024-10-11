import { AuthContext } from "@/main";
import { useContext, useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

const Callback: React.FC = () => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const navigate = useNavigate();
// eslint-disable-next-line @typescript-eslint/no-unsafe-argument
const { login } = useContext(AuthContext);

  useEffect(()=> {
    //if code, good, else, route to home
    if (code) {
        console.log("Code: " + code)
        const sendCode = () => {
          const base_url : string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string
                 fetch(`${base_url}/login`, {
                    method: "POST",
                    credentials: 'include',
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body :JSON.stringify({code}),
                })
                .then(response => {
                  if (!response.ok){
                    //Navgate back to login page
                    navigate("/")
                    return
                }
                else {

                  //Successful login. Navigate to dashboard page and call login
                  // eslint-disable-next-line @typescript-eslint/no-unsafe-call
                  login()
                  navigate("/app/dashboard")
                }
                }

                )
                .catch((err: unknown) => {
                  navigate("/")
                  console.log("Error Occured: ", err)
                return});
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
