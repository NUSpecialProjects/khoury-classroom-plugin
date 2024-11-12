import React from "react";
import "./styles.css";

interface IDropdownProps {
  label: string;
  name: string;
  placeholder: string;
  options: string[];
}

const Dropdown: React.FC<IDropdownProps> = ({
  label,
  name,
  placeholder,
  options,
}) => {
  return (
    <div className="Dropdown__wrapper">
      <label className="Dropdown__label" htmlFor={name}>
        {label}
      </label>
      <select id={name} name={name} className="Dropdown">
        <option className="Dropdown__option">{placeholder}</option>
        {options.map((option, index) => (
          <option className="Dropdown__option" key={index} value={option}>
            {option}
          </option>
        ))}
      </select>
    </div>
  );
};

export default Dropdown;
