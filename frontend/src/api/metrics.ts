const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export async function getAssignmentAcceptanceMetrics(
  classroomID: number,
  assignmentID: number
): Promise<IAssignmentAcceptanceMetrics> {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/progress-status`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  const resp: IAssignmentAcceptanceMetricsResponse = await response.json();
  return resp.status;
}

export async function getAssignmentGradedMetrics(
  classroomID: number,
  assignmentID: number
): Promise<IAssignmentGradedMetrics> {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/grading-status`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!response.ok) {
    throw new Error(response.statusText);
  }

  const resp: IAssignmentGradedMetricsResponse = await response.json();
  return resp.status;
}
