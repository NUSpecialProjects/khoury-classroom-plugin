import React from "react";
import "./styles.css";

interface IInputProps extends React.InputHTMLAttributes<HTMLInputElement> {
    label: string;
    name: string;
    placeholder?: string;
    caption?: string;
}

const Input: React.FC<IInputProps> = ({
    label,
    name,
    placeholder,
    caption,
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
                className="Input">
            </input>
            <p className="Input__caption">{caption}</p>
        </div>
    );
};

export default Input;
