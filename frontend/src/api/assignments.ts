const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getAssignments = async (
  classroomId: number
): Promise<IAssignmentOutline[]> => {
  const result = await fetch(
    `${base_url}/classrooms/classroom/${classroomId}/assignments`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data = await result.json();

  return data.assignment_outlines as IAssignmentOutline[];
};

export const getAssignmentIndirectNav = async (
  classroomID: number,
  assignmentID: number
): Promise<IAssignmentOutline> => {
  const result = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data: IAssignmentOutline = (
    (await result.json()) as IAssignmentOutlineResponse
  ).assignment_outline;

  return data;
};


export const setAssignmentRubric = async (
  rubric_id: number,
  classroomID: number,
  assignmentID: number,
): Promise<IAssignmentOutline> => {
  const response = await fetch(
    `${base_url}/classrooms/classroom/${classroomID}/assignments/assignment/${assignmentID}/rubric`,
    {
      method: "PUT",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(rubric_id),
    }
  );
  
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  const data: IAssignmentOutline = (await response.json() as IAssignmentOutlineResponse).assignment_outline 
  console.log(data)

  return data

};
