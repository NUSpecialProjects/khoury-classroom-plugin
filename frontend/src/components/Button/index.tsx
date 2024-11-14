import { Link } from "react-router-dom";

import "./styles.css";

interface IButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  href?: string;
  variant?: "primary" | "secondary";
}

const ButtonWrapper: React.FC<IButtonProps> = ({ children, href }) => {
  return href ? (
    <Link to={href} target="_blank">
      {children}
    </Link>
  ) : (
    <>{children}</>
  );
};

const Button: React.FC<IButtonProps> = ({
  children,
  href,
  variant = "primary",
  ...props
}) => {
  return (
    <ButtonWrapper href={href}>
      <button className={`Button Button--${variant}`} {...props}>
        {children}
      </button>
    </ButtonWrapper>
  );
};

export default Button;
