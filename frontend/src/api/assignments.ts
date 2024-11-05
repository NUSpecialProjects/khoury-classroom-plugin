const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getAssignments = async (
  classroomId: number
): Promise<IAssignment[]> => {
  const result = await fetch(
    `${base_url}/classrooms/${classroomId}/assignments`,
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
  
  return data.assignment_outlines as IAssignment[] || []
};