export interface Organization {
    login: string;
    id: number;
    html_url: string;
    name: string;
    avatar_url: string;
}

export interface OrganizationsResponse {
    orgs_with_app: Organization[];
    orgs_without_app: Organization[];
}
