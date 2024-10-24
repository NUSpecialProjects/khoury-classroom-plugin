const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getOrganizations = async (): Promise<IOrganizationsResponse> => {
  const response = await fetch(`${base_url}/github/user/orgs`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json() as Promise<IOrganizationsResponse>;
};

export const getClassrooms = async (
  orgId: number
): Promise<IClassroomResponse> => {
  const response = await fetch(
    `${base_url}/github/user/orgs/${orgId.toString()}/classrooms`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json() as Promise<IClassroomResponse>;
};

export const getOrganizationDetails = async (
  login: string
): Promise<IOrganization> => {
  const response = await fetch(`${base_url}/github/orgs/${login}`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const resp = (await response.json()) as { org: IOrganization };
  return resp.org;
};

export const postSemester = async (
  orgId: number,
  classroomId: number,
  OrgName: string,
  ClassroomName: string
): Promise<ISemester> => {
  const response = await fetch(`${base_url}/github/semesters`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      org_id: orgId,
      classroom_id: classroomId,
      org_name: OrgName,
      classroom_name: ClassroomName,
    }),
  });
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const data = (await response.json()) as { semester: ISemester };
  return data.semester;
};

export const getUserSemesters = async (): Promise<IUserSemestersResponse> => {
  const response = await fetch(`${base_url}/github/user/semesters`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json();
};

export const getOrgSemesters = async (
  orgId: number
): Promise<IOrgSemestersResponse> => {
  const response = await fetch(
    `${base_url}/github/orgs/${orgId.toString()}/semesters`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json();
};

export const activateSemester = async (
  classroomId: number
): Promise<ISemester> => {
  return modifySemester(classroomId, true);
};

export const deactivateSemester = async (
  classroomId: number
): Promise<ISemester> => {
  return modifySemester(classroomId, false);
};

const modifySemester = async (
  classroomId: number,
  activate: boolean
): Promise<ISemester> => {
  const response = await fetch(
    `${base_url}/github/semesters/${classroomId.toString()}`,
    {
      method: "PUT",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ activate }),
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const data = (await response.json()) as { semester: ISemester };
  return data.semester;
};
