const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getStudentAssignment = async (
  semesterID: number,
  assignmentID: number,
  studentAssignmentID: number
): Promise<IStudentAssignment> => {
  const response = await fetch(
    `${base_url}/semesters/${semesterID}/assignments/${assignmentID}/student-assignments/${studentAssignmentID}`,
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
  const resp = (await response.json()) as IStudentAssignment;
  return resp;
};

export const getStudentAssignments = async (
  semesterID: number,
  assignmentID: number
): Promise<IStudentAssignment[]> => {
  const response = await fetch(
    `${base_url}/semesters/${semesterID}/assignments/${assignmentID}/student-assignments`,
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
  const resp = (await response.json()) as IStudentAssignment[];
  return resp;
};

export const getTotalStudentAssignments = async (
  semesterID: number,
  assignmentID: number
): Promise<number> => {
  const response = await fetch(
    `${base_url}/semesters/${semesterID}/assignments/${assignmentID}/student-assignments/total`,
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
  const resp = (await response.json()) as number;
  return resp;
};
