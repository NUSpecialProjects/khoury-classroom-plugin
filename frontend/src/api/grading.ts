const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const createPRComment = async (
  classroomID: number,
  assignmentID: number,
  studentWorkID: number,
  comments: IGradingComment[]
) => {
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
        comments,
      }),
    }
  );
  if (!response.ok) {
    console.log(response);
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
  node: IFileTreeNode
): Promise<IGraderFile> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works/work/${studentWorkID}/blob/${node.sha}`,
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
  const content = await response.text();
  const file: IGraderFile = { content, name: node.name };
  return file;
};
