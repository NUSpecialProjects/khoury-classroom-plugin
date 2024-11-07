import React from "react";
import "./styles.css";

interface IInputProps {
  label: string;
  name: string;
  placeholder: string;
}

const Input: React.FC<IInputProps> = ({
  label,
  name,
  placeholder,
}) => {
  return (
    <div className="Input__wrapper">
      <label className="Input__label" htmlFor={name}>
        {label}
      </label>
      <input id={name} name={name} placeholder={placeholder} className="Input"></input>
    </div>
  );
};

export default Input;
