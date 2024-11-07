interface IGitHubUser {
  login: string;
  id: number;
  node_id: string;
  avatar_url: string;
  url: string;
  name: string | null;
  email: string | null;
}

interface IGitHubUserResponse {
  user: IGitHubUser;
}

interface IClassroomUser {
  id: number;
  first_name: string;
  last_name: string;
  github_username: string;
  github_user_id: number;
  classroom_role: string;
}
