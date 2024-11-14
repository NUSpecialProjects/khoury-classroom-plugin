import React from "react";
import './styles.css';
import '../styles.css';
import { MdChevronRight } from "react-icons/md"

const BreadcrumbPageHeader: React.FC<IBreadcrumbPageHeader> = ({ pageTitle, breadcrumbItems }) => {
    return (
        <div className="PageHeader__wrapper">
            <h1 className="BreadcrumbPageHeader__pageTitle">
                {pageTitle}
                {breadcrumbItems.map((item, index) => (
                    <React.Fragment key={index}>
                        {index < breadcrumbItems.length - 1 ? (
                            <>
                                <MdChevronRight />
                                <div className="BreadcrumbPageHeader__item">
                                    {item}
                                </div>
                            </>
                        ) : (
                            <>
                                <MdChevronRight />
                                <div className="BreadcrumbPageHeader__lastItem">
                                    {item}
                                </div>
                            </>
                        )}
                    </React.Fragment>
                ))}
            </h1>
        </div>
    )
}

export default BreadcrumbPageHeader;