import React from "react";
import "./styles.css";

interface IInputProps extends React.InputHTMLAttributes<HTMLInputElement> {
    label: string;
    name: string;
    placeholder?: string;
    caption?: string;
    rightElement?: React.ReactNode;
}

const Input: React.FC<IInputProps> = ({
    label,
    name,
    placeholder,
    caption,
    rightElement,
    ...props
}) => {
    return (
        <div className="Input__wrapper">
            <label className="Input__label" htmlFor={name}>
                {label}
            </label>
            <div className="Input__container">
                <input
                    id={name}
                    name={name}
                    placeholder={placeholder}
                    {...props}
                    className="Input"
                />
                {rightElement && <div className="Input__rightElement">{rightElement}</div>}
            </div>
            {caption && 
            <p className="Input__caption">{caption}</p>
            }
        </div>
    );
};

export default Input;
