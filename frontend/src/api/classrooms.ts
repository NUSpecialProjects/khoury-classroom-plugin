const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export async function getClassroomsInOrg(
  orgId: number
): Promise<IClassroomListResponse> {
  const response = await fetch(`${base_url}/orgs/org/${orgId}/classrooms`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const resp: IClassroomListResponse = await response.json();
  return resp;
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
  const resp: IClassroom = (await response.json());
  return resp;
}
