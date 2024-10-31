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

export const getUserOrgsAndClassrooms = async (): Promise<IUserOrgsAndClassroomsResponse> => {
    const response = await fetch(`${base_url}/github/user/orgs-and-classrooms`, {
        method: "GET",
        credentials: "include",
        headers: {
            "Content-Type": "application/json",
        },
    });
    if (!response.ok) {
        throw new Error("Network response was not ok");
    }
    return response.json() as Promise<IUserOrgsAndClassroomsResponse>;
}