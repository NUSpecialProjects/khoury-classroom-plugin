const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const fetchCurrentUser = async (): Promise<IGitHubUser> => {
  const response = await fetch(`${base_url}/github/user`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  return response.json()
};

export const fetchUsersWithRole = async (role_type: string, semester: ISemester): Promise<IGitHubUser[]> => {
    const response = await fetch(`${base_url}/github/semesters/${semester.classroom_id.toString()}/roles/${role_type}/users`, {
        method: "GET",
        credentials: "include",
        headers: {
        "Content-Type": "application/json",
        },
    });
    
    if (!response.ok) {
        throw new Error("Network response was not ok");
    }
    
    return response.json()
    }
