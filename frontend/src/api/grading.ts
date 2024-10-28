const base_url: string = import.meta.env.VITE_PUBLIC_API_DOMAIN as string;

export const createPRComment = async (
  orgName: string,
  repoName: string,
  commitSha: string,
  filePath: string,
  line: number,
  comment: string
) => {
  const response = await fetch(
    `${base_url}/grading/org/${orgName}/repo/${repoName}/comment`,
    {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        commit_sha: commitSha,
        file_path: filePath,
        line,
        comment,
      }),
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
};

export const getGitTree = async (
  orgName: string,
  repoName: string
): Promise<IGitTree> => {
  const response = await fetch(
    `${base_url}/grading/org/${orgName}/repo/${repoName}/tree`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const resp = (await response.json()) as IGitTree;
  return resp;
};

export const getGitBlob = async (
  orgName: string,
  repoName: string,
  node: IFileTreeNode
): Promise<IGraderFile> => {
  const response = await fetch(
    `${base_url}/grading/org/${orgName}/repo/${repoName}/blob/${node.sha}`,
    {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  const content = await response.text();
  const file: IGraderFile = { content, name: node.name };
  return file;
};
