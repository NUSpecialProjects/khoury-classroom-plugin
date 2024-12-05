import React from 'react';
import { FaChevronUp } from 'react-icons/fa6';
import './styles.css';

interface CaretProps {
    isUp?: boolean;
    className?: string;
    onClick?: () => void;
    clickable?: boolean;
}

const Caret: React.FC<CaretProps> = ({ 
    isUp = true, 
    className = '', 
    onClick,
    clickable = false 
}) => {
    const clickableClass = clickable ? 'Caret--clickable' : 'Caret--not-clickable';
    return (
        <span 
            className={`Caret ${!isUp ? 'Caret--rotated' : ''} ${clickableClass} ${className}`} 
            onClick={clickable ? onClick : undefined} 
        >
            <FaChevronUp />
        </span>
    );
};

export default Caret;