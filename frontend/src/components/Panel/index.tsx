import React from "react";
import './styles.css'

interface IPanel {
    title: string;
    children: React.ReactNode;
}

const Panel: React.FC<IPanel> = ({ title, children }) => {
    return (
        <div className="Panel">
            <div className="Panel__wrapper">
                <h1 className="Panel__title">{title}</h1>
                <div className="Panel__content">
                    {children}
                </div>
            </div>
        </div>
    );
}

export default Panel;