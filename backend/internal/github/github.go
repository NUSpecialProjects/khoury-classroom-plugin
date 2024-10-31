package github

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
)

type GitHubAppClient interface { // All methods in the APP client
	GitHubBaseClient

	GetWebhookSecret() string

	// Get the installations of the github app
	ListInstallations(ctx context.Context) ([]*github.Installation, error)
	GetFileTree(owner string, repo string, pullNumber int) ([]models.FileTreeNode, error)
	GetFileBlob(owner string, repo string, sha string) ([]byte, error)

	// Create a new team in an organization
	CreateTeam(ctx context.Context, org_name, team_name string) (*github.Team, error)

	// Add a user to a team
	AddTeamMember(ctx context.Context, team_id int64, user_name string, opt *github.TeamAddTeamMembershipOptions) error

	// Add a repository permission to a team
	AssignPermissionToTeam(ctx context.Context, team_id int64, owner_name string, repo_name string, permission string) error

	// Add a repository permission to a user
	AssignPermissionToUser(ctx context.Context, ownerName string, repoName string, userName string, permission string) error
}
}

type GitHubUserClient interface { // All methods in the OAUTH client
	GitHubBaseClient

	// Get the details of an organization
	GetOrg(ctx context.Context, org_name string) (*github.Organization, error)

	// Get the current authenticated user
	GetCurrentUser(ctx context.Context) (models.GitHubUser, error)

	// Get the organizations the authenticated user is part of
	GetUserOrgs(ctx context.Context) ([]models.Organization, error)

	// Callback for the OAUTH flow
	GitHubCallback(code string, clientCfg config.GitHubUserClient) (string, error)
}

type GitHubBaseClient interface { //All methods in the SHARED client

	// Check the health of the API
	Ping(ctx context.Context) (string, error)

	// List the repositories in an organization
	ListRepositoriesByOrg(ctx context.Context, orgName string) ([]*github.Repository, error)

	// List the commits in a repository
	ListCommits(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, error)

	// Get the details of a branch
	GetBranch(ctx context.Context, owner_name string, repo_name string, branch_name string) (*github.Branch, error)

	// Create a new branch in a repository
	CreateBranch(ctx context.Context, owner string, repo string, baseBranch string, newBranchName string) error

	// Get the details of a pull request
	GetPullRequest(ctx context.Context, owner string, repo string, pullNumber int) (*github.PullRequest, error)

	// Get the diff of a pull request
	GetPullRequestDiff(ctx context.Context, owner string, repo string, pullNumber int) (string, error)

	// Create a new pull request in a repository
	CreatePullRequest(ctx context.Context, owner string, repo string, baseBranch string, headBranch string, title string, body string) (*github.PullRequest, error)
	//TODO: these two methods need to be fixed with API change
	// CreateInlinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, line int, side string, commentBody string) (*github.PullRequestComment, error)
	// CreateMultilinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, startLine int, endLine int, side string, commentBody string) (*github.PullRequestComment, error)
	CreateFilePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, commentBody string) (*github.PullRequestComment, error)
	CreateRegularPRComment(ctx context.Context, owner string, repo string, pullNumber int, commentBody string) (*github.IssueComment, error)
}
