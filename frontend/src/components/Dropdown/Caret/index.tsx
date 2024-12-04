import React from 'react';
import './styles.css';

interface CaretProps {
    isUp?: boolean;
    className?: string;
    onClick?: () => void;
    clickable?: boolean;
}

const Caret: React.FC<CaretProps> = ({ 
    isUp = false, 
    className = '', 
    onClick,
    clickable = false 
}) => {
    const clickableClass = clickable ? 'Caret--clickable' : 'Caret--not-clickable';
    return (
        <span 
            className={`Caret ${isUp ? 'Caret--up' : ''} ${clickableClass} ${className}`} 
            onClick={clickable ? onClick : undefined} 
        />
    );
};

export default Caret;