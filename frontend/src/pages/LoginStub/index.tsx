import { useSearchParams } from "react-router-dom";

const LoginStub: React.FC = () => {
  const [searchParams] = useSearchParams();
  const code = searchParams.get("code");

  return (
    <div>
      <p>OAuth Code: {code}</p>
    </div>
  );
};



export default LoginStub;