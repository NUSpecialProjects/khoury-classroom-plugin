import React from "react";
import './styles.css';
import '../styles.css';
import { Link } from "react-router-dom";
import { FaChevronLeft } from "react-icons/fa6";

const SubPageHeader: React.FC<ISubPageHeader> = ({ pageTitle, chevronLink, children }) => {
    return (
        <div className="PageHeader__wrapper">
            <h1 className="SubPageHeader__pageTitle">
                <div className="SubPageHeader__leftSideContents">
                    <Link to={chevronLink}>
                        <FaChevronLeft />
                    </Link>
                    {pageTitle}
                </div>
                <div className="SubPageHeader__rightSideContents">
                    {children}
                </div>
            </h1>
        </div>
    )
}

export default SubPageHeader;