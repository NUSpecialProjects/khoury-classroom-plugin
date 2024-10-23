
import { useState, useEffect } from "react";
import Cookies from "js-cookie";
import { useNavigate } from "react-router-dom";

const COOKIE_NAME = "selectedSemester";
interface ISelectedSemester {
  selectedSemester: ISemester | null;
  setSelectedSemester: (semester: ISemester) => void;
}

const useSelectedSemester = (): ISelectedSemester => {
  const [selectedSemester, setSelectedSemesterState] =
    useState<ISemester | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const cookieValue = Cookies.get(COOKIE_NAME);
    if (cookieValue) {
      try {
        const parsedValue = JSON.parse(cookieValue) as ISemester;
        setSelectedSemesterState(parsedValue);
      } catch (error: unknown) {
        console.log("Error parsing semester cookie: ", error);
      }
    } else {
      navigate("/class-selection");
    }
  }, [navigate]);

  const setSelectedSemester = (semester: ISemester | null) => {
    if (!semester) {
      Cookies.remove(COOKIE_NAME);
      setSelectedSemesterState(null);
    } else {
      setSelectedSemesterState(semester);
      Cookies.set(COOKIE_NAME, JSON.stringify(semester), {
        expires: 30,
        SameSite: "Strict",
      });
    }
  };

  return { selectedSemester, setSelectedSemester };
};

export default useSelectedSemester;
