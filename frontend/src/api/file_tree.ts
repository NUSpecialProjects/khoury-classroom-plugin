const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const createPRComment = async (
  orgName: string,
  repoName: string,
  filePath: string,
  line: number,
  comment: string
) => {
  const response = await fetch(
    `${base_url}/grading/org/${orgName}/repo/${repoName}/comment`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        file_path: filePath,
        line,
        comment,
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
  const resp = (await response.json()) as IGitTreeNode[];
  return resp;
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
