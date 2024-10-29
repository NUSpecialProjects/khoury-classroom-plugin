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
  return Promise.resolve({
    available_classrooms: [
      {
        id: 1,
        name: "classroom1",
        url: "https://classroom1.com",
      },
      {
        id: 2,
        name: "classroom2",
        url: "https://classroom2.com",
      },
    ],
    unavailable_classrooms: [
      {
        id: 3,
        name: "classroom3",
        url: "https://classroom3.com",
      },
      {
        id: 4,
        name: "classroom4",
        url: "https://classroom4.com",
      },
    ],
  })
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
  return Promise.resolve({
    id: 1,
    org_id: orgId,
    classroom_id: classroomId,
    org_name: OrgName,
    classroom_name: ClassroomName,
    active: true,
  });
};

export const getUserSemesters = async (): Promise<IUserSemestersResponse> => {
  return Promise.resolve({
    active_semesters: [
      {
        org_id: 1,
        classroom_id: 1,
        org_name: "org1",
        classroom_name: "classroom1",
        active: true,
      },
      {
        org_id: 2,
        classroom_id: 2,
        org_name: "org2",
        classroom_name: "classroom2",
        active: false,
      },
    ],
    inactive_semesters: [
      {
        org_id: 2,
        classroom_id: 3,
        org_name: "org3",
        classroom_name: "classroom3",
        active: false,
      },
      {
        org_id: 1,
        classroom_id: 4,
        org_name: "org4",
        classroom_name: "classroom4",
        active: false,
      },
    ],
  });
};


interface ISemester {
  org_id: number;
  classroom_id: number;
  org_name: string;
  classroom_name: string;
  active: boolean;
}

interface IUserSemestersResponse {
  active_semesters: ISemester[];
  inactive_semesters: ISemester[];
}

interface IOrgSemestersResponse {
  semesters: ISemester[];
}


export const getOrgSemesters = async (
  orgId: number
): Promise<IOrgSemestersResponse> => {
  return Promise.resolve({
    semesters: [
      {
        org_id: 1,
        classroom_id: 1,
        org_name: "org1",
        classroom_name: "classroom1",
        active: true,
      },
      {
        org_id: 2,
        classroom_id: 2,
        org_name: "org2",
        classroom_name: "classroom2",
        active: true,
      },
    ],
  });
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
  return Promise.resolve({
    org_id: 1,
    classroom_id: classroomId,
    org_name: "org1",
    classroom_name: "classroom1",
    active: activate,
  });
};
