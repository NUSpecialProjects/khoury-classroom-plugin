// Represents a GitHub repository.
interface IRepository {
  id: number;
  name: string;
  owner: IGitHubUser;
  private: boolean;
  description?: string | null;
  url: string;
  is_template: boolean;
  archived: boolean;
}

interface IRepositoryResponse {
  template_repos: IRepository[];
}
