import React, { useEffect, useState } from "react";
import "./styles.css";

interface IPill {
    label: string,
    variant?: string
}

const Pill: React.FC<IPill> = (
    {
        label, 
        variant = "default"
    }) => {
    return (
        <div className={`Pill Pill--${variant}`}>
            <div className="Pill__label">{label}</div>
        </div>
    );
}

export default Pill;