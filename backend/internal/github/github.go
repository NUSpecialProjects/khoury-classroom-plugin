package github

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
)

type GitHubAppClient interface { // All methods in the app client
	GitHubBaseClient
	GetWebhookSecret() string
	ListInstallations(ctx context.Context) ([]*github.Installation, error)
	GetFileTree(owner string, repo string, pullNumber int) ([]models.FileTreeNode, error)
	GetFileBlob(owner string, repo string, sha string) ([]byte, error)
}

type GitHubUserClient interface { // All methods in the OAUTH client
	GitHubBaseClient
	GetOrg(ctx context.Context, org_name string) (*github.Organization, error)
	GetCurrentUser(ctx context.Context) (models.GitHubUser, error)
	GetUserOrgs(ctx context.Context) ([]models.Organization, error)
	GitHubCallback(code string, clientCfg config.GitHubUserClient) (string, error)
	CreateTeam(ctx context.Context, org_name, team_name string) (*github.Team, error)
	AddTeamMember(ctx context.Context, team_id int64, user_name string, opt *github.TeamAddTeamMembershipOptions) error
	AssignPermissionToTeam(ctx context.Context, team_id int64, owner_name string, repo_name string, permission string) error
}

type GitHubBaseClient interface { //All methods in the shared client
	Ping(ctx context.Context) (string, error)
	ListRepositoriesByOrg(ctx context.Context, orgName string) ([]*github.Repository, error)
	ListCommits(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, error)
	GetBranch(ctx context.Context, owner_name string, repo_name string, branch_name string) (*github.Branch, error)
	CreateBranch(ctx context.Context, owner string, repo string, baseBranch string, newBranchName string) error
	GetPullRequest(ctx context.Context, owner string, repo string, pullNumber int) (*github.PullRequest, error)
	GetPullRequestDiff(ctx context.Context, owner string, repo string, pullNumber int) (string, error)
	CreatePullRequest(ctx context.Context, owner string, repo string, baseBranch string, headBranch string, title string, body string) (*github.PullRequest, error)
	// temp unused
	// CreateFilePRComment(ctx context.Context, owner string, repo string, commitSha string, filePath string, comment string) (*github.PullRequestComment, error)
	CreateLinePRComment(ctx context.Context, owner string, repo string, commitSha string, filePath string, line int64, comment string) (*github.PullRequestComment, error)
}
