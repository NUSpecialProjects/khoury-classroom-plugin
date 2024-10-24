const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getAssignments = async (
  classroomId: number
): Promise<IAssignment[]> => {
  const result = await fetch(`${base_url}/assignments/${classroomId}`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!result.ok) {
    throw new Error("Network response was not ok");
  }

  const data: IAssignment[] = (await result.json()) as IAssignment[];
  const formatted = data.map((assignment: IAssignment) => ({
    ...assignment,
    main_due_date: assignment.main_due_date
      ? new Date(assignment.main_due_date)
      : null,
  }));
  return formatted;
};
