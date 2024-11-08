import { useState, useLayoutEffect, createContext } from "react";
import Cookies from "js-cookie";

const COOKIE_NAME = "selectedClassroom";

interface ISelectedClassroomContext {
  selectedClassroom: IClassroom | null;
  setSelectedClassroom: (classroom: IClassroom) => void;
}

export const SelectedClassroomContext: React.Context<ISelectedClassroomContext> =
  createContext<ISelectedClassroomContext>({
    selectedClassroom: null,
    setSelectedClassroom: (_: IClassroom) => {},
  });

const SelectedClassroomProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [selectedClassroom, setSelectedClassroomState] =
    useState<IClassroom | null>(null);
  const [loading, setLoading] = useState(true);

  useLayoutEffect(() => {
    const cookieValue = Cookies.get(COOKIE_NAME);
    if (cookieValue) {
      try {
        const parsedValue = JSON.parse(cookieValue) as IClassroom;
        setSelectedClassroomState(parsedValue);
      } catch (_: unknown) {
        // do nothing
      }
    }
    setLoading(false);
  }, []);

  const setSelectedClassroom = (classroom: IClassroom | null) => {
    if (!classroom) {
      Cookies.remove(COOKIE_NAME);
      setSelectedClassroomState(null);
    } else {
      setSelectedClassroomState(classroom);
      Cookies.set(COOKIE_NAME, JSON.stringify(classroom), {
        expires: 30,
        SameSite: "Strict",
      });
    }
  };

  return (
    !loading && (
      <SelectedClassroomContext.Provider
        value={{ selectedClassroom: selectedClassroom, setSelectedClassroom }}
      >
        {children}
      </SelectedClassroomContext.Provider>
    )
  );
};

export default SelectedClassroomProvider;
