package github

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
)

type GitHubAppClient interface { // All methods in the app client
	GitHubBaseClient
	ListInstallations(ctx context.Context) ([]*github.Installation, error)
	GetWebhookSecret() string
}

type GitHubUserClient interface { // All methods in the OAUTH client
	GitHubBaseClient
	GetOrg(ctx context.Context, org_name string) (*github.Organization, error)
	GetCurrentUser(ctx context.Context) (models.GitHubUser, error)
	GetUserOrgs(ctx context.Context) ([]models.Organization, error)
	GetUserClassrooms(ctx context.Context) ([]models.Classroom, error)
	ListAssignmentsForClassroom(ctx context.Context, classroom_id int64) ([]models.ClassroomAssignment, error)
	GetAcceptedAssignments(ctx context.Context, assignment_id int64) ([]models.ClassroomAcceptedAssignment, error)
	GitHubCallback(code string, clientCfg config.GitHubUserClient) (string, error)
	CreateTeam(ctx context.Context, org_name, team_name string) (*github.Team, error)
	AddTeamMember(ctx context.Context, team_id int64, user_name string, opt *github.TeamAddTeamMembershipOptions) error
	AssignPermissionToTeam(ctx context.Context, team_id int64, owner_name string, repo_name string, permission string) error
	CreateOrgRole(ctx context.Context, org_id int64, role_name string, desc string, permissions []string, base_role string) (*models.OrganizationRole, error)
	CreateOrgRoleFromTemplate(ctx context.Context, org_id int64, template_role models.OrganizationTemplateRole) (*models.OrganizationRole, error)
	DeleteOrgRole(ctx context.Context, org_id int64, role_id int64) error
	AssignOrgRoleToUser(ctx context.Context, org_id int64, user_name string, role_id int64) error
	RemoveOrgRoleFromUser(ctx context.Context, org_id int64, user_name string, role_id int64) error
	GetOrgRoles(ctx context.Context, org_id int64) ([]models.OrganizationRole, error)
	GetUsersAssignedToRole(ctx context.Context, org_id int64, role_id int64) ([]models.GitHubUser, error)
	GetUserRoles(ctx context.Context, org_id int64) ([]models.OrganizationRole, error)
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
	//TODO: these two methods need to be fixed with API change
	// CreateInlinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, line int, side string, commentBody string) (*github.PullRequestComment, error)
	// CreateMultilinePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, startLine int, endLine int, side string, commentBody string) (*github.PullRequestComment, error)
	CreateFilePRComment(ctx context.Context, owner string, repo string, pullNumber int, commitID string, path string, commentBody string) (*github.PullRequestComment, error)
	CreateRegularPRComment(ctx context.Context, owner string, repo string, pullNumber int, commentBody string) (*github.IssueComment, error)
}
