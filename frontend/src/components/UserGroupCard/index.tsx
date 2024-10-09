import React from "react";
import "./styles.css";

interface IUserGroupCardProps {
    label: string;
    number: number;
}

const UserGroupCard: React.FC<IUserGroupCardProps> = ({ label, number }) => {
    let userIcons = [];

    if (number > 3) {
        // Cap placeholders at 3
        userIcons = [1, 2, 3].map((_, index) => (
            <div className="UserGroupCard__icon" key={index} />
        ));
    } else {
        // Render placeholders equal to the number
        userIcons = Array.from({ length: number }).map((_, index) => (
            <div className="UserGroupCard__icon" key={index} />
        ));
    }

    return (
        <div className="UserGroupCard">
            <div className="UserGroupCard__content">
                <h3 className="UserGroupCard__label">{label}</h3>
                <div className="UserGroupCard__detailsWrapper">
                    <div className="UserGroupCard__icons">
                        {userIcons}
                    </div>
                    <p className="UserGroupCard__number">{number}</p>
                </div>
            </div>
        </div>
    );
}

export default UserGroupCard;
