package github

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
)

type GitHubAppClient interface { // All methods in the app client
	GitHubBaseClient
	// App only functionalities
}

type GitHubUserClient interface { // All methods in the OAUTH client
	GitHubBaseClient
	ListClassrooms(ctx context.Context) ([]models.Classroom, error)
	ListAssignmentsForClassroom(ctx context.Context, classroomID int64) ([]models.ClassroomAssignment, error)
	GetAcceptedAssignments(ctx context.Context, assignmentID int64) ([]models.ClassroomAcceptedAssignment, error)
	GitHubCallback(code string, clientCfg config.GitHubUserClient, authCfg config.AuthHandler) (string, error)
}

type GitHubBaseClient interface { //All methods in the shared client
	Ping(ctx context.Context) (string, error)
	ListRepositories(ctx context.Context) ([]*github.Repository, error)
	ListCommits(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, error)
	GetBranch(ctx context.Context, owner_name string, repo_name string, branch_name string) (*github.Branch, error)
	CreateBranch(ctx context.Context, owner string, repo string, baseBranch string, newBranchName string) error
	CreatePullRequest(ctx context.Context, owner string, repo string, baseBranch string, headBranch string, title string, body string) (*github.PullRequest, error)
	// CreateInlinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, line int, side string, commentBody string) (*github.PullRequestComment, error)
	// CreateMultilinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, startLine int, endLine int, side string, commentBody string) (*github.PullRequestComment, error)
	CreateFilePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, commentBody string) (*github.PullRequestComment, error)
	CreateRegularPRComment(ctx context.Context, owner string, repo string, pullNumber int, commentBody string) (*github.IssueComment, error)
}
