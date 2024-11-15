const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const getAppInstallations =
  async (): Promise<IOrganizationsResponse> => {
    const response = await fetch(`${base_url}/orgs/installations`, {
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

export const getOrganizationDetails = async (
  login: string
): Promise<IOrganization> => {
  const response = await fetch(`${base_url}/orgs/org/${login}`, {
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

export const getOrganizationTemplates = async (
  orgName: string,
  itemsPerPage: string = "",
  pageNum: string = ""
): Promise<ITemplatesResponse> => {
  const url = new URL(`${base_url}/orgs/org/${orgName}/templates`);
  url.searchParams.append("items_per_page", itemsPerPage);
  url.searchParams.append("page_num", pageNum);

  const response = await fetch(url.toString(), {
    method: "GET",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (!response.ok) {
    throw new Error("Network response was not ok");
  }

  return response.json() as Promise<ITemplatesResponse>;
};
