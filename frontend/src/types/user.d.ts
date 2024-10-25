interface IGitHubUser {
  login: string;
  id: number;
  node_id: string;
  avatar_url: string;
  url: string;
  name: string | null;
  email: string | null;
}