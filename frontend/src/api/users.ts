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

  return response.json();
};

export const fetchUsersWithRole = async (
  role_type: string,
  classroom: IClassroom
): Promise<IGitHubUser[]> => {
  console.log(
    "Using mocked API call for role: ",
    role_type,
    "semester: ",
    classroom
  );
  return Promise.resolve([
    {
      login: "user1",
      id: 1,
      node_id: "node1",
      avatar_url: "https://avatars.githubusercontent.com/u/1?v=4",
      url: "https://api.github.com/users/user1",
      name: "User One",
      email: null,
    },
    {
      login: "user2",
      id: 2,
      node_id: "node2",
      avatar_url: "https://avatars.githubusercontent.com/u/2?v=4",
      url: "https://api.github.com/users/user2",
      name: "User Two",
      email: "",
    },
    {
      login: "user3",
      id: 3,
      node_id: "node3",
      avatar_url: "https://avatars.githubusercontent.com/u/3?v=4",
      url: "https://api.github.com/users/user3",
      name: null,
      email: "",
    },
  ]);
};

export const useRoleToken = async (token: string): Promise<IMessageResponse> => {
      const response = await fetch(`${base_url}/github/role-token/use`, { //TODO: this url needs to change
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ token: token }),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      return response.json() as Promise<IMessageResponse>;
}

export const createToken = async (role_type: string, classroom: IClassroom): Promise<ITokenResponse> => {

      const response = await fetch(`${base_url}/github/role-token/create`, { //TODO: this url will change
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          classroom: classroom,
          role_type: role_type,
        }),
      });

      if (!response.ok) {
        throw new Error("Network response was not ok");
      }

      return response.json() as Promise<ITokenResponse>;
  };
