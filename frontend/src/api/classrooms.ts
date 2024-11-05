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
  const resp: IClassroom = await response.json();
  return resp;
}

export async function postClassroomToken(
  classroomId: number,
  role: string,
  duration?: number // Duration is optional
): Promise<IClassroomToken> {
  const response = await fetch(`${base_url}/classrooms/${classroomId}/token`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      classroom_role: role,
      duration: duration,
      })
  });

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  const resp: IClassroomToken = await response.json();
  return resp;
}

export async function useClassroomToken(
  token: string
): Promise<IMessageResponse> {
  const response = await fetch(`${base_url}/classrooms/token/${token}`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  const resp: IMessageResponse = await response.json();
  return resp;
}
