import React from 'react';
import './styles.css';
import { FaRegCopy } from "react-icons/fa";

interface ICopyLinkProps {
    link: string;
    name: string;
}

const CopyLink: React.FC<ICopyLinkProps> = ({
    link,
    name
}) => {
    const copyLink = async() => {
        try {
            await navigator.clipboard.writeText(link);
            alert('Link copied to clipboard!');
        } catch (_) {
            alert('Error copying link, please try again.');
        }
    }

    return (
        <div className="CopyLink">
            <input id={name} type="text" value={link} className="CopyLink__link" readOnly></input>
            <button onClick={copyLink} className="CopyLink__copyButton"><FaRegCopy /></button>
        </div>
    );
};

export default CopyLink;