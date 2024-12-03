import { createContext, useContext, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

import { SelectedClassroomContext } from "./selectedClassroom";
import { getPaginatedStudentWork } from "@/api/student_works";
import { createPRReview } from "@/api/grader";

interface IGraderContext {
  assignmentID: string | undefined;
  studentWorkID: string | undefined;
  studentWork: IPaginatedStudentWork | null;
  selectedFile: IFileTreeNode | null;
  feedback: IGraderFeedbackMap;
  stagedFeedback: IGraderFeedbackMap;
  setSelectedFile: React.Dispatch<React.SetStateAction<IFileTreeNode | null>>;
  addFeedback: (feedback: IGraderFeedback) => number;
  editFeedback: (feedbackID: number, feedback: IGraderFeedback) => void;
  removeFeedback: (feedbackID: number) => void;
  postFeedback: () => void;
}

export const GraderContext: React.Context<IGraderContext> =
  createContext<IGraderContext>({
    assignmentID: undefined,
    studentWorkID: undefined,
    studentWork: null,
    selectedFile: null,
    feedback: {},
    stagedFeedback: {},
    setSelectedFile: () => {},
    addFeedback: () => 0,
    editFeedback: () => {},
    removeFeedback: () => {},
    postFeedback: () => {},
  });

export const GraderProvider: React.FC<{
  assignmentID: string | undefined;
  studentWorkID: string | undefined;
  children: React.ReactNode;
}> = ({ assignmentID, studentWorkID, children }) => {
  const { selectedClassroom } = useContext(SelectedClassroomContext);

  const [nextFeedbackID, setNextFeedbackID] = useState(0);
  const [feedback, setFeedback] = useState<IGraderFeedbackMap>({});
  const [stagedFeedback, setStagedFeedback] = useState<IGraderFeedbackMap>({});
  const [studentWork, setStudentWork] = useState<IPaginatedStudentWork | null>(
    null
  );
  const [selectedFile, setSelectedFile] = useState<IFileTreeNode | null>(null);

  const navigate = useNavigate();

  // fetch requested student assignment
  useEffect(() => {
    // reset states
    setSelectedFile(null);
    setStudentWork(null);

    if (!selectedClassroom || !assignmentID || !studentWorkID) return;

    getPaginatedStudentWork(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID)
    )
      .then((resp) => {
        setStudentWork(resp.student_work);
        setFeedback(resp.feedback);
        setStagedFeedback({});
      })
      .catch((_: unknown) => {
        navigate("/404", { replace: true });
      });
  }, [studentWorkID]);

  const getNextFeedbackID = () => {
    const tmp = nextFeedbackID;
    setNextFeedbackID(nextFeedbackID + 1);
    return tmp;
  };

  const addFeedback = (fb: IGraderFeedback) => {
    const id = getNextFeedbackID();
    setStagedFeedback((prevFeedback) => ({
      ...prevFeedback,
      [id]: fb,
    }));
    return id;
  };

  const editFeedback = (_feedbackID: number, _feedback: IGraderFeedback) => {};

  const removeFeedback = (_feedbackID: number) => {};

  const postFeedback = () => {
    if (!selectedClassroom || !assignmentID || !studentWorkID) return;

    createPRReview(
      selectedClassroom.id,
      Number(assignmentID),
      Number(studentWorkID),
      stagedFeedback
    ).then(() => {
      setFeedback((prevFeedback) => ({
        ...prevFeedback,
        ...stagedFeedback,
      }));
      setStagedFeedback({});
    });
  };

  // once feedback is updated, reset id to its length
  // this is so when posting staged feedback, it will never overwrite existing feedback
  useEffect(() => {
    setNextFeedbackID(feedback ? Object.keys(feedback).length : 0);
  }, [feedback]);

  return (
    <GraderContext.Provider
      value={{
        assignmentID,
        studentWorkID,
        studentWork,
        selectedFile,
        feedback,
        stagedFeedback,
        setSelectedFile,
        addFeedback,
        editFeedback,
        removeFeedback,
        postFeedback,
      }}
    >
      {children}
    </GraderContext.Provider>
  );
};
