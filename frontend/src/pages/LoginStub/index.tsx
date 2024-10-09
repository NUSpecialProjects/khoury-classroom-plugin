import { useEffect } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

const LoginStub: React.FC = () => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");
  const navigate = useNavigate();


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
                    navigate("")
                }
                console.log(response)
            }
            catch (err: unknown){
                //navigate to login page
                navigate("")
            }
            navigate("/app/")
        }
        sendCode()
    }
    else {
        navigate("")
    }
  })



  return (
    <div>
      <p>Loading...</p>
    </div>
  );
};



export default LoginStub;
