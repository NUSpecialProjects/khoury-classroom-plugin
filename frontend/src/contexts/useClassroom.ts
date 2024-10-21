/* eslint-disable */

import { useState, useEffect } from "react";
import Cookies from "js-cookie";
import { Semester } from "@/types/semester";

const COOKIE_NAME = "selectedSemester";

const useSelectedSemester = () => {
    const [selectedSemester, setSelectedSemesterState] = useState<Semester | null>(null);

    useEffect(() => {
        const cookieValue = Cookies.get(COOKIE_NAME);
        if (cookieValue) {
            setSelectedSemesterState(JSON.parse(cookieValue));
        }
    }, []);

    const setSelectedSemester = (semester: Semester) => {
        setSelectedSemesterState(semester);
        Cookies.set(COOKIE_NAME, JSON.stringify(semester), { expires: 30, sameSite: 'Strict' });
    };

    return [selectedSemester, setSelectedSemester] as const;
};

export default useSelectedSemester;