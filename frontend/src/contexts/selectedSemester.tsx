import { useState, useLayoutEffect, createContext } from "react";
import Cookies from "js-cookie";

const COOKIE_NAME = "selectedSemester";

interface ISelectedSemesterContext {
  selectedSemester: ISemester | null;
  setSelectedSemester: (semester: ISemester) => void;
}

export const SelectedSemesterContext: React.Context<ISelectedSemesterContext> =
  createContext<ISelectedSemesterContext>({
    selectedSemester: null,
    setSelectedSemester: (_: ISemester) => {},
  });

const SelectedSemesterProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [selectedSemester, setSelectedSemesterState] =
    useState<ISemester | null>(null);
  const [loading, setLoading] = useState(true);

  useLayoutEffect(() => {
    const cookieValue = Cookies.get(COOKIE_NAME);
    if (cookieValue) {
      try {
        const parsedValue = JSON.parse(cookieValue) as ISemester;
        setSelectedSemesterState(parsedValue);
      } catch (error: unknown) {
        console.log("Error parsing semester cookie: ", error);
      }
    }
    setLoading(false);
  }, []);

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

  return (
    !loading && (
      <SelectedSemesterContext.Provider
        value={{ selectedSemester, setSelectedSemester }}
      >
        {children}
      </SelectedSemesterContext.Provider>
    )
  );
};

export default SelectedSemesterProvider;
