 const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getAssignments = async (
  classroomId: number
): Promise<IAssignment[]> => {
  // const result = await fetch(
  //   `${base_url}/semesters/${classroomId}/assignments`,
  //   {
  //     method: "GET",
  //     credentials: "include",
  //     headers: {
  //       "Content-Type": "application/json",
  //     },
  //   }
  // );

  // if (!result.ok) {
  //   throw new Error("Network response was not ok");
  // }

  // return (await result.json()) as IAssignment[];
  console.log("Using mocked API call for assignments in: ", classroomId);

  return Promise.resolve([
    {
      id: 1,
      rubric_id: 1,
      assignment_classroom_id: 1,
      semester_id: 1,
      name: "Assignment",
      inserted_date: null,
      main_due_date: null,
    },
  ]);
};


export const CreateAssignment = async () => {
  const result = await fetch(
      `${base_url}/forks/fork`,
      {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
      }
    );

    if (!result.ok) {
        console.log(result)
    }

}