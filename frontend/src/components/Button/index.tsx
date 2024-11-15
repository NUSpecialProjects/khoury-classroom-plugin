import { Link } from "react-router-dom";

import "./styles.css";

interface IButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  href?: string;
  variant?: "primary" | "secondary";
  size?: "default" | "small";
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
  onClick,
  variant = "primary",
  size = "default"
}) => {
  const className = `Button--${variant}${size === "small" ? ` ${size}` : ""}`;

  return (
    <ButtonWrapper href={href}>
      <button className={className} onClick={onClick} >
        {children}
      </button>
    </ButtonWrapper>
  );
};

export default Button;
