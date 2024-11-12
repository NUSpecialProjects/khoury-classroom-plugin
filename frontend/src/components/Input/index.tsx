import React from "react";
import "./styles.css";

interface IInputProps extends React.InputHTMLAttributes<HTMLInputElement>{
    label: string;
    name: string;
    placeholder?: string;
}

const Input: React.FC<IInputProps> = ({
    label,
    name,
    placeholder,
    ...props
}) => {
    return (
        <div className="Input__wrapper">
            <label className="Input__label" htmlFor={name}>
                {label}
            </label>
            <input
                id={name}
                name={name}
                placeholder={placeholder}
                {...props}
                className="Input"></input>
        </div>
    );
};

export default Input;
