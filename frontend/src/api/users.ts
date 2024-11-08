const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const fetchCurrentUser = async (): Promise<IGitHubUser> => {
  const response = await fetch(`${base_url}/user`, {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  const data: IGitHubUserResponse = await response.json();
  return data.user;
};

export const fetchUser = async (
  user_name: string
): Promise<IGitHubUserResponse> => {
  const response = await fetch(`${base_url}/users/user/${user_name}`, {
    method: "GET",
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  return response.json();
};