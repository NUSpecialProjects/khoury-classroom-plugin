import { Link } from "react-router-dom";

import "./styles.css";

interface IButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  href?: string;
  variant?: "primary" | "secondary";
  size?: "default" | "small";
  newTab?: boolean;
}

const ButtonWrapper: React.FC<IButtonProps> = ({
  children,
  href,
  newTab = false,
}) => {
  if (!href) {
    return <>{children}</>;
  }

  // For external URLs or when newTab is explicitly set
  if (newTab) {
    return (
      <a href={href} target="_blank" rel="noopener noreferrer">
        {children}
      </a>
    );
  }

  // For internal routing
  if (!href.startsWith('http')) {
    return <Link to={href}>{children}</Link>;
  }

  // For external URLs without new tab
  return <a href={href}>{children}</a>;
};

const Button: React.FC<IButtonProps> = ({
  className,
  children,
  href,
  variant = "primary",
  size = "default",
  newTab,
  ...props
}) => {
  return (
    <ButtonWrapper href={href} newTab={newTab}>
      <button
        className={`Button Button--${variant} Button--${size} ${className ?? ""}`}
        {...props}
      >
        {children}
      </button>
    </ButtonWrapper>
  );
};

export default Button;
