
import { useState, useEffect } from "react";
import Cookies from "js-cookie";

const COOKIE_NAME = "selectedSemester";

interface ISelectedSemester {
    selectedSemester: ISemester | null;
    setSelectedSemester: (semester: ISemester) => void;
}

const useSelectedSemester: () => ISelectedSemester = () => {
    const [selectedSemester, setSelectedSemesterState] = useState<ISemester | null>(null);

    useEffect(() => {
        const cookieValue = Cookies.get(COOKIE_NAME);
        if (cookieValue) {
            
            try {            
                const parsedValue = JSON.parse(cookieValue);
                const sem:ISemester = {
                    classroom_id: parsedValue?.classroom_id,
                    org_id: parsedValue?.org_id,
                    classroom_name: parsedValue?.classroom_name,
                    org_name: parsedValue?.name,
                    active: parsedValue?.active
                }
                setSelectedSemesterState(sem)

            } catch (error: unknown) {
                console.log("Error parsing json: ", error)
            }
        }
    }, []);

    const setSelectedSemester = (semester: ISemester) => {
        setSelectedSemesterState(semester);
        Cookies.set(COOKIE_NAME, JSON.stringify(semester), { expires: 30, sameSite: 'Strict' });
    };

    return {selectedSemester, setSelectedSemester};
};

export default useSelectedSemester;
