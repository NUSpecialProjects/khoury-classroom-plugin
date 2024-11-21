import React from "react";
import "./styles.css";

interface ICheckbox extends React.InputHTMLAttributes<HTMLInputElement> {
    label: string;
    name: string;
    caption?: string;
}

const Checkbox: React.FC<ICheckbox> = ({
    label,
    name,
    caption,
    ...props
}) => {
    return (
        <div className="Checkbox__wrapper">
            <div className="Checkbox__checkWrapper">
                <input
                    id={name}
                    name={name}
                    type="checkbox"
                    {...props}
                    className="Checkbox">
                </input>
                <label className="Checkbox__label" htmlFor={name}>
                    {label}
                </label>
            </div>
            {caption &&
                <p className="Input__caption">{caption}</p>
            }
        </div>
    );
};

export default Checkbox;
