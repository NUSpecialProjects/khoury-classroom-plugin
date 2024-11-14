const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const createPRComment = async (
  classroomID: number,
  assignmentID: number,
  studentWorkID: number,
  comments: IGradingCommentMap
) => {
  console.log(comments);
  const comments1D: IGradingComment[] = [];
  for (const pathComments of Object.values(comments)) {
    for (const lineComments of Object.values(pathComments)) {
      for (const comment of Object.values(lineComments)) {
        comments1D.push(comment);
      }
    }
  }

  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works/work/${studentWorkID}/grade`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        body: "",
        comments: comments1D,
      }),
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
};

export const getFileTree = async (
  classroomID: number,
  assignmentID: number,
  studentWorkID: number
): Promise<IGitTreeNode[]> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works/work/${studentWorkID}/tree`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const resp = (await response.json()) as IGitTreeResponse;
  return resp.tree;
};

export const getFileBlob = async (
  classroomID: number,
  assignmentID: number,
  studentWorkID: number,
  sha: string
): Promise<string> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works/work/${studentWorkID}/blob/${sha}`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  return await response.text();
};
