const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export function getClassroomsInOrg(
  orgId: number
): Promise<IClassroomListResponse> {
  console.log("Using mocked API call for org: ", orgId);
  return Promise.resolve({
    classrooms: [
      {
        id: 1,
        name: "Spring 2024",
        org_id: orgId,
        org_name: "Organization One",
      },
      {
        id: 2,
        name: "Fall 2023",
        org_id: orgId,
        org_name: "Organization One",
      },
    ],
  });
}

export async function postClassroom(
  classroom: Omit<IClassroom, "id">
): Promise<IClassroom> {
  const response = await fetch(`${base_url}/classrooms`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(classroom),
  });
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const resp = (await response.json()) as IClassroom;
  return resp;
}
