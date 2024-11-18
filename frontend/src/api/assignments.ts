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
  classroomid: number,
  assignmentID: number
): Promise<IAssignmentOutline> => {
  const result = await fetch(`${base_url}/classrooms/classroom/${classroomid}/assignments/assignment/${assignmentID}`, {
    method: 'GET',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data: IAssignmentOutline = (await result.json() as IAssignmentOutlineResponse).assignment_outline

  return data
};

export const createAssignment = async (
  templateRepoID: number,
  assignment: IAssignmentFormData
): Promise<IAssignmentFormData> => {
  const result = await fetch(
    `${base_url}/classrooms/classroom/${assignment.classroomId}/assignments`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        template_id: templateRepoID,
        name: assignment.assignmentName,
        classroom_id: assignment.classroomId,
        group_assignment: assignment.groupAssignment,
        main_due_date: assignment.mainDueDate,
      })
    }
  );

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data = (await result.json())

  return data.assignment_outline as IAssignmentFormData
};
