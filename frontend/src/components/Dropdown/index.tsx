import { ChangeEvent, useState } from "react";
import Caret from "./Caret";
import "./styles.css";

export interface IDropdownOption {
  value: string;
  label: string;
  disabled?: boolean;
}

interface IDropdownProps {
  options: IDropdownOption[];
  onChange?: (selected: string) => void;
  selectedOption: string | null;
  loading: boolean;
  labelText: string;
  captionText?: string;
  placeholder?: string;
  loadingText?: string;
  noOptionsText?: string;
}

const GenericDropdown = ({
  onChange,
  options,
  selectedOption,
  loading,
  labelText,
  captionText,
  placeholder = 'Select an option',
  loadingText = 'Loading options...',
  noOptionsText = 'No options available',
}: IDropdownProps) => {
  const [isOpen, setIsOpen] = useState(false);

  const handleChange = (event: ChangeEvent<HTMLSelectElement>) => {
      const selected = event.target.value;
      if (onChange) {
          onChange(selected);
      }
  };

  const handleClick = () => {
    setIsOpen(!isOpen);
  };

  const renderOptions = () => {
      if (loading) {
          return (
              <option value="" disabled>
                  {loadingText}
              </option>
          );
      } else if (options.length === 0) {
          return (
              <option value="" disabled>
                  {noOptionsText}
              </option>
          );
      }

      return options.map((option, index) => (
          <option key={index} value={option.value} disabled={option.disabled}>
              {option.label}
          </option>
      ));
  };

  return (
      <div className="Dropdown">
          <div className="Dropdown__wrapper">
              <label className="Dropdown__label" htmlFor="dropdown">
                  {labelText}
              </label>
              <div className="Dropdown__select-wrapper">
                  <select
                      className="Dropdown__select"
                      value={selectedOption || ''}
                      onChange={handleChange}
                      onClick={handleClick}
                      onBlur={() => setIsOpen(false)}
                  >
                      <option className="Dropdown__option" value="" disabled>
                          {placeholder}
                      </option>
                      {renderOptions()}
                  </select>
                  <Caret isUp={!isOpen} clickable={false}/>
              </div>
              {captionText && <div className='Dropdown__caption'>{captionText}</div>}
          </div>
      </div>
  );
};

export default GenericDropdown;