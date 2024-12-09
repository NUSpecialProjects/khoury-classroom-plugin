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


export const getStudentWorkCommitsPerDay = async (
  classroomID: number,
  assignmentID: number,
  studentWorkID: number
) : Promise<Map<Date, number>> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/works/work/${studentWorkID}/commits-per-day`,
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
  const resp = (await response.json() as ICommitsPerDayResponse);
  
  console.log(resp)
  const commitsMap = new Map<Date, number>(
    Object.entries(resp.dated_commits).map(([key, value]) => [(new Date(key)), value])
  );
  console.log(commitsMap)

  return commitsMap;
}
