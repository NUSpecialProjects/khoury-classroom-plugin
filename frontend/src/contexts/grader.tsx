import { createContext, useContext, useState, useEffect } from "react";

import { SelectedClassroomContext } from "./selectedClassroom";
import { getPaginatedStudentWork } from "@/api/student_works";
import { createPRReview } from "@/api/grading";
import { useNavigate } from "react-router-dom";

interface IGraderContext {
  assignmentID: string | undefined;
  studentWorkID: string | undefined;
  studentWork: IPaginatedStudentWork | null;
  selectedFile: IFileTreeNode | null;
  feedback: IGradingFeedbackMap;
  stagedFeedback: IGradingFeedbackMap;
  setSelectedFile: React.Dispatch<React.SetStateAction<IFileTreeNode | null>>;
  addFeedback: (feedback: IGradingFeedback) => number;
  editFeedback: (feedbackID: number, feedback: IGradingFeedback) => void;
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
  const [nextFeedbackID, setNextFeedbackID] = useState(0);
  const { selectedClassroom } = useContext(SelectedClassroomContext);
  const [feedback, setFeedback] = useState<IGradingFeedbackMap>({});
  const [stagedFeedback, setStagedFeedback] = useState<IGradingFeedbackMap>({});
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
      })
      .catch((err: unknown) => {
        console.log(err);
        navigate("/404", { replace: true });
      });
  }, [studentWorkID]);

  const getNextFeedbackID = () => {
    const tmp = nextFeedbackID;
    setNextFeedbackID(nextFeedbackID + 1);
    return tmp;
  };

  const addFeedback = (fb: IGradingFeedback) => {
    const id = getNextFeedbackID();
    setStagedFeedback((prevFeedback) => ({
      ...prevFeedback,
      [id]: fb,
    }));
    return id;
  };

  const editFeedback = (feedbackID: number, feedback: IGradingFeedback) => {};

  const removeFeedback = (feedbackID: number) => {};

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

  useEffect(() => {
    console.log(feedback);
    setNextFeedbackID(feedback ? Object.keys(feedback).length : 0);
  }, [feedback]);

  useEffect(() => {
    console.log(stagedFeedback);
  }, [stagedFeedback]);

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
