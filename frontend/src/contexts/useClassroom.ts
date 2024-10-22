import { useState, useEffect } from "react";
import Cookies from "js-cookie";
const COOKIE_NAME = "selectedSemester";
interface ISelectedSemester {
  selectedSemester: ISemester | null;
  setSelectedSemester: (semester: ISemester) => void;
}
const useSelectedSemester = (): ISelectedSemester => {
  const [selectedSemester, setSelectedSemesterState] =
    useState<ISemester | null>(null);

  useEffect(() => {
    const cookieValue = Cookies.get(COOKIE_NAME);
    if (cookieValue) {
      try {
        const parsedValue = JSON.parse(cookieValue) as ISemester;
        setSelectedSemesterState(parsedValue);
      } catch (error: unknown) {
        console.log("Error parsing semester cookie: ", error);
      }
    }
  }, []);

  const setSelectedSemester = (semester: ISemester) => {
    setSelectedSemesterState(semester);
    Cookies.set(COOKIE_NAME, JSON.stringify(semester), {
      expires: 30,
      sameSite: "Strict",
    });
  };

  return { selectedSemester, setSelectedSemester };
};

export default useSelectedSemester;
