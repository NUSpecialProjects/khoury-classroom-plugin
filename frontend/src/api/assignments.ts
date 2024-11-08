 const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getAssignments = async (
  classroomId: number
): Promise<IAssignmentOutline[]> => {
  const result = await fetch(
    `${base_url}/assignments/classrooms/classroom/${classroomId}/all`,
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

  const data = (await result.json())

  return data.assignment_outlines as IAssignmentOutline[]
};


export const getAssignmentIndirectNav = async (
  classroomid: number, assignmentID: number
): Promise<IAssignmentOutline> => {
  const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;
  const result = await fetch(`${base_url}/assignments/classrooms/classroom/${classroomid}/assignment/${assignmentID}`, {
    method: 'GET',
    credentials: 'include',
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!result.ok) {
    throw new Error('Network response was not ok');
  }

  const data: IAssignmentOutline = (await result.json() as IAssignmentOutlineResponse).assignment_outline
  return data


};

export const AcceptAssignment = async (orgName: string, repoName: string, classroomID: number) => {
  const result = await fetch(
      `${base_url}/assignments/classrooms/classroom/${classroomID}/accept`,
      {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          org_name: orgName,
          repo_name: repoName,
          assignment_name: "Cam Test Assignment: POC",
          assignment_id: 1,
          org_id: 182810684
        }),
      }
    );

    if (!result.ok) {
        console.log(result)
    }

}
