import { ClassroomResponse } from "@/types/classroom";
import { Organization, OrganizationsResponse } from "@/types/organization";

const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getOrganizations = async (): Promise<OrganizationsResponse> => {
    const response = await fetch(`${base_url}/github/user/orgs`, {
        method: "GET",
        credentials: 'include',
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Network response was not ok");
    }
    return response.json() as Promise<OrganizationsResponse>;
};

export const getClassrooms = async (orgId: number): Promise<ClassroomResponse> => {
    const response = await fetch(`${base_url}/github/user/orgs/${orgId}/classrooms`, {
        method: "GET",
        credentials: 'include',
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Network response was not ok");
    }
    return response.json() as Promise<ClassroomResponse>;
};

export const getOrganizationDetails = async (login: string): Promise<Organization> => {
    const response = await fetch(`${base_url}/github/orgs/${login}`, {
        method: "GET",
        credentials: 'include',
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Network response was not ok");
    }
    const resp = await response.json() as { org: Organization };
    return resp.org;
};

//TODO: Change the post endpoint to return the new semester
export const postSemester = async (orgId: number, classroomId: number, name: string): Promise<void> => {
    const response = await fetch(`${base_url}/github/semesters`, {
        method: "POST",
        credentials: 'include',
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            org_id: orgId,
            classroom_id: classroomId,
            name: name,
        }),
    });
    if (!response.ok) {
        throw new Error("Network response was not ok");
    }
};