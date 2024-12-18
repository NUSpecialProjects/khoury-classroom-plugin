const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const gradeWork = async (
  classroomID: number,
  assignmentID: number,
  studentWorkID: number,
  feedback: IGraderFeedbackMap
) => {
  const feedback1D: IGraderFeedback[] = [];
  for (const fb of Object.values(feedback)) {
    feedback1D.push(fb);
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
        comments: feedback1D,
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
