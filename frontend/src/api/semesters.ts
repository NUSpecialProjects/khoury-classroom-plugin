// const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getOrganizations = async (): Promise<IOrganizationsResponse> => {
  // const response = await fetch(`${base_url}/orgs/installations`, {
  //   method: "GET",
  //   credentials: "include",
  //   headers: {
  //     "Content-Type": "application/json",
  //   },
  // });
  // if (!response.ok) {
  //   throw new Error("Network response was not ok");
  // }
  // return response.json() as Promise<IOrganizationsResponse>;
  console.log("Using mocked API call for organizations");
  return Promise.resolve({
    orgs_with_app: [
      {
        login: "CS2200",
        id: 1,
        html_url: "nicktietje.com",
        name: "Organization One",
        avatar_url: "https://avatars.githubusercontent.com/u/1?v=4",
      },
      {
        login: "CS3500",
        id: 2,
        html_url: "nicktietje.com",
        name: "Organization Two",
        avatar_url: "https://avatars.githubusercontent.com/u/2?v=4",
      },
    ],
    orgs_without_app: [
      {
        login: "CHME4500",
        id: 3,
        html_url: "nicktietje.com",
        name: "Organization Three",
        avatar_url: "https://avatars.githubusercontent.com/u/3?v=4",
      },
      {
        login: "CS5600",
        id: 4,
        html_url: "nicktietje.com",
        name: "Organization Four",
        avatar_url: "https://avatars.githubusercontent.com/u/4?v=4",
      },
    ],
  });
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
        org_name: "CS2200",
        classroom_name: "Spring 2024",
        active: true,
      },
    ],
    inactive_semesters: [
      {
        org_id: 2,
        classroom_id: 2,
        org_name: "CS3500",
        classroom_name: "Fall 2023",
        active: false,
      },
      {
        org_id: 2,
        classroom_id: 3,
        org_name: "CS3500",
        classroom_name: "Spring 2022",
        active: false,
      },
      {
        org_id: 1,
        classroom_id: 4,
        org_name: "CS3500",
        classroom_name: "Spring 2021",
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
  console.log("Using mocked API call for org: ", orgId);
  return Promise.resolve({
    semesters: [
      {
        org_id: 1,
        classroom_id: 1,
        org_name: "CS2200",
        classroom_name: "Spring 2024",
        active: true,
      },
      {
        org_id: 1,
        classroom_id: 4,
        org_name: "CS3500",
        classroom_name: "Spring 2021",
        active: false,
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
  console.log(
    "Using mocked API call for classroom: ",
    classroomId,
    "activate: ",
    activate
  );
  return Promise.resolve({
    org_id: 1,
    classroom_id: 1,
    org_name: "CS2200",
    classroom_name: "Spring 2024",
    active: true,
  });
};
