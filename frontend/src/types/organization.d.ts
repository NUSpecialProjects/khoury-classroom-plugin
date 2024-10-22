interface IOrganization {
  login: string;
  id: number;
  html_url: string;
  name: string;
  avatar_url: string;
}

interface IOrganizationsResponse {
  orgs_with_app: Organization[];
  orgs_without_app: Organization[];
}
