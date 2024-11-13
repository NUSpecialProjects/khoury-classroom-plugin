import React from "react";
import './styles.css';
import '../styles.css';
import { IPageHeader } from '../types';
import { FaChevronLeft } from "react-icons/fa6";

interface ISubPageHeader extends IPageHeader {
    children : React.ReactNode;
}

const SubPageHeader: React.FC<ISubPageHeader> = ({ pageTitle, children }) => {
    return (
        <div className="PageHeader__wrapper">
            <h1 className="SubPageHeader__pageTitle">
                <div className="SubPageHeader__leftSideContents">
                    <FaChevronLeft/>
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