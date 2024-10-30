// const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const fetchCurrentUser = async (): Promise<IGitHubUser> => {
  // const response = await fetch(`${base_url}/github/user`, {
  //   method: "GET",
  //   credentials: "include",
  //   headers: {
  //     "Content-Type": "application/json",
  //   },
  // });

  // if (!response.ok) {
  //   throw new Error("Network response was not ok");
  // }

  // return response.json();
  console.log("Using mocked API call for current user: ");
  return Promise.resolve(
    {
      login: "user1",
      id: 1,
      node_id: "node1",
      avatar_url: "https://avatars.githubusercontent.com/u/1?v=4",
      url: "https://api.github.com/users/user1",
      name: "User One",
      email: null,
    })
};

export const fetchUsersWithRole = async (
  role_type: string,
  semester: ISemester
): Promise<IGitHubUser[]> => {
  console.log(
    "Using mocked API call for role: ",
    role_type,
    "semester: ",
    semester
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
