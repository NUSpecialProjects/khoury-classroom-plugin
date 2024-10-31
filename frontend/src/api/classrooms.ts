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

export const getUserOrgsAndClassrooms =
  async (): Promise<IUserOrgsAndClassroomsResponse> => {
    // const response = await fetch(`${base_url}/github/user/orgs-and-classrooms`, {
    //     method: "GET",
    //     credentials: "include",
    //     headers: {
    //         "Content-Type": "application/json",
    //     },
    // });
    // if (!response.ok) {
    //     throw new Error("Network response was not ok");
    // }
    // return response.json() as Promise<IUserOrgsAndClassroomsResponse>;
    console.log("Using mocked API call for user orgs and classrooms");
    return Promise.resolve({
      orgs_and_classrooms: new Map([
        [
          {
            login: "CS2200",
            id: 1,
            html_url: "nicktietje.com",
            name: "Organization One",
            avatar_url: "https://avatars.githubusercontent.com/u/1?v=4",
          },
          [
            {
              id: 1,
              name: "Classroom One",
              org_id: 1,
              org_name: "Organization One",
            },
            {
              id: 2,
              name: "Classroom Two",
              org_id: 1,
              org_name: "Organization One",
            },
          ],
        ],
        [
          {
            login: "CS3500",
            id: 2,
            html_url: "nicktietje.com",
            name: "Organization Two",
            avatar_url: "https://avatars.githubusercontent.com/u/2?v=4",
          },
          [
            {
              id: 3,
              name: "Classroom Three",
              org_id: 2,
              org_name: "Organization Two",
            },
            {
              id: 4,
              name: "Classroom Four",
              org_id: 2,
              org_name: "Organization Two",
            },
          ],
        ],
      ]),
    });
  };
