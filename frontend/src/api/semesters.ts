import { ClassroomResponse } from "@/types/classroom";
import { Organization, OrganizationsResponse } from "@/types/organization";
import {
  OrgSemestersResponse,
  Semester,
  UserSemestersResponse,
} from "@/types/semester";

const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getOrganizations = async (): Promise<OrganizationsResponse> => {
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
  return response.json() as Promise<OrganizationsResponse>;
};

export const getClassrooms = async (
  orgId: number
): Promise<ClassroomResponse> => {
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
  return response.json() as Promise<ClassroomResponse>;
};

export const getOrganizationDetails = async (
  login: string
): Promise<Organization> => {
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
  const resp = (await response.json()) as { org: Organization };
  return resp.org;
};

export const postSemester = async (
  orgId: number,
  classroomId: number,
  OrgName: string,
  ClassroomName: string
): Promise<Semester> => {
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
  const data = (await response.json()) as { semester: Semester };
  return data.semester;
};

export const getUserSemesters = async (): Promise<UserSemestersResponse> => {
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
): Promise<OrgSemestersResponse> => {
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
  orgId: number,
  classroomId: number
): Promise<Semester> => {
  return modifySemester(orgId, classroomId, true);
};

export const deactivateSemester = async (
  orgId: number,
  classroomId: number
): Promise<Semester> => {
  return modifySemester(orgId, classroomId, false);
};

const modifySemester = async (
  orgId: number,
  classroomId: number,
  activate: boolean
): Promise<Semester> => {
  const response = await fetch(
    `${base_url}/github/semesters/${orgId.toString()}/${classroomId.toString()}`,
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
  const data = (await response.json()) as { semester: Semester };
  return data.semester;
};
