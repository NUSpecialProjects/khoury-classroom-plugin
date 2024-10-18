// import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import "./styles.css";


const SemesterSelection: React.FC = () => {
    const navigate = useNavigate();

    return <div>
        <h1>Select a Semester</h1>
        <button onClick={() => { navigate("/app/dashboard") }}> Go to Dashboard Page</button>
    </div>;
}

export default SemesterSelection;