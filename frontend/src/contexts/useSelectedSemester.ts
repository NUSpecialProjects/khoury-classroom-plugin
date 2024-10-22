/* eslint-disable */

import { useState, useEffect } from "react";
import Cookies from "js-cookie";
import { Semester } from "@/types/semester";
import { useNavigate } from "react-router-dom";

const COOKIE_NAME = "selectedSemester";

const useSelectedSemester = () => {
    const [selectedSemester, setSelectedSemesterState] = useState<Semester | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const cookieValue = Cookies.get(COOKIE_NAME);
        if (cookieValue) {
            setSelectedSemesterState(JSON.parse(cookieValue));
        } else {
            navigate("/semester-selection");
        }
    }, [navigate]);

    const setSelectedSemester = (semester: Semester | null) => {
        if (!semester) {
            Cookies.remove(COOKIE_NAME);
            setSelectedSemesterState(null);
        } else {
            setSelectedSemesterState(semester);
            Cookies.set(COOKIE_NAME, JSON.stringify(semester), { expires: 30, SameSite: 'Strict' });
        }
    };

    return { selectedSemester, setSelectedSemester };
};

export default useSelectedSemester;