import React from "react";
import './styles.css';

interface IEmptyDataBanner {
    children: React.ReactNode;
}

const EmptyDataBanner: React.FC<IEmptyDataBanner> = ({children}) => {
   return (
    <div className="emptyDataBanner">
        {children}
    </div>
   )
}

export default EmptyDataBanner;