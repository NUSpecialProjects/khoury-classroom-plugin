import { createContext, useState } from "react";

interface IGraderContext {
  assignmentID: string | undefined;
  studentWorkID: string | undefined;
  comments: IGradingCommentMap;
  addComment: (comment: IGradingComment) => number;
  editComment: (commentID: number, comment: IGradingComment) => void;
  removeComment: (commentID: number) => void;
}

export const GraderContext: React.Context<IGraderContext> =
  createContext<IGraderContext>({
    assignmentID: undefined,
    studentWorkID: undefined,
    comments: {},
    addComment: () => 0,
    editComment: () => {},
    removeComment: () => {},
  });

export const GraderProvider: React.FC<{
  assignmentID: string | undefined;
  studentWorkID: string | undefined;
  children: React.ReactNode;
}> = ({ assignmentID, studentWorkID, children }) => {
  const [comments, setComments] = useState<IGradingCommentMap>({});

  const addComment = (comment: IGradingComment) => {
    let newCommentID = 0;
    if (comments[comment.path] && comments[comment.path][comment.line]) {
      newCommentID = Object.keys(comments[comment.path][comment.line]).length;
    }
    setComments((prevComments) => ({
      ...prevComments,
      [comment.path]: {
        ...prevComments[comment.path],
        [comment.line]: {
          ...(prevComments[comment.path]?.[comment.line] || {}),
          [newCommentID]: comment,
        },
      },
    }));
    console.log(comments);
    return newCommentID;
  };

  const editComment = (commentID: number, comment: IGradingComment) => {};

  const removeComment = (commentID: number) => {};

  return (
    <GraderContext.Provider
      value={{
        assignmentID,
        studentWorkID,
        comments,
        addComment,
        editComment,
        removeComment,
      }}
    >
      {children}
    </GraderContext.Provider>
  );
};
