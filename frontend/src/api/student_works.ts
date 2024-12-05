const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getPaginatedStudentWork = async (
  classroomID: number,
  assignmentID: number,
  studentWorkID: number
): Promise<IPaginatedStudentWorkResponse> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works/work/${studentWorkID}`,
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
  const resp = (await response.json()) as IPaginatedStudentWorkResponse;
  return resp;
};

export const getStudentWorks = async (
  classroomID: number,
  assignmentID: number
): Promise<IStudentWork[]> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works`,
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
  const resp = ((await response.json()) as IStudentWorkResponses).student_works;
  return resp;
};

export const getAssignmentFirstCommit = async (
  classroomID: number,
  assignmentID: number 
): Promise<Date> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/first-commit`,
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
  const resp = ((await response.json()) as IAssignmentCommitDate).first_commit_at;
  return resp;
};
