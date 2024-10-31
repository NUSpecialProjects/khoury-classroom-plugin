package appclient

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/sharedclient"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/google/go-github/github"
	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"
)

type AppAPI struct { //app API
	sharedclient.CommonAPI
	webhooksecret  string
	appTokenSource oauth2.TokenSource
}

func New(cfg *config.GitHubAppClient) (*AppAPI, error) {
	// Read private key and IDs from the config
	privateKey := []byte(cfg.Key)
	appID := cfg.AppID
	installationID := cfg.InstallationID

	// Create an Application Token Source
	appTokenSource, err := githubauth.NewApplicationTokenSource(appID, privateKey)
	if err != nil {
		return nil, fmt.Errorf("error creating application token source: %v", err)
	}

	// Create an Installation Token Source
	installationTokenSource := githubauth.NewInstallationTokenSource(installationID, appTokenSource)

	// Create an OAuth2 HTTP client
	httpClient := oauth2.NewClient(context.Background(), installationTokenSource)

	// Create the GitHub client
	githubClient := github.NewClient(httpClient)

	return &AppAPI{
		CommonAPI: sharedclient.CommonAPI{
			Client: githubClient,
		},
		webhooksecret:  cfg.WebhookSecret,
		appTokenSource: appTokenSource,
	}, nil
}

func (api *AppAPI) GetWebhookSecret() string {
	return api.webhooksecret
}

// GetJWT generates a JWT from the appTokenSource
func (api *AppAPI) getJWT() (string, error) {
	token, err := api.appTokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("error getting token: %v", err)
	}
	return token.AccessToken, nil
}

func (api *AppAPI) getClientWithJWTAuth(ctx context.Context) (*github.Client, error) {
	// Create a new OAuth2 client with the JWT
	token, err := api.getJWT()
	if err != nil {
		return nil, fmt.Errorf("error getting app JWT: %v", err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	// Create a new GitHub client with the authenticated HTTP client
	return github.NewClient(tc), nil
}

func (api *AppAPI) ListInstallations(ctx context.Context) ([]*github.Installation, error) {
	client, err := api.getClientWithJWTAuth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting github client with JWT auth: %v", err)
	}

	// List installations
	installations, _, err := client.Apps.ListInstallations(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing installations: %v", err)
	}

	return installations, nil
}

// Helper function to create a map of file paths to modified status
func createStatusMap(modifiedFiles []*github.CommitFile) map[string]string {
	statusMap := make(map[string]string)
	for _, file := range modifiedFiles {
		statusMap[file.GetFilename()] = file.GetStatus()
	}
	return statusMap
}

func (api *AppAPI) GetFileTree(owner string, repo string, pullNumber int) ([]models.FileTreeNode, error) {
	// Get the reference to the branch
	ref, _, err := api.Client.Git.GetRef(context.Background(), owner, repo, "heads/main")
	if err != nil {
		return nil, fmt.Errorf("error fetching branch ref: %v", err)
	}

	// Get the latest commit from the ref
	commitSha := ref.Object.GetSHA()
	commit, _, err := api.Client.Git.GetCommit(context.Background(), owner, repo, commitSha)
	if err != nil {
		return nil, fmt.Errorf("error fetching commit: %v", err)
	}

	// Get the git tree from latest commit
	treeSHA := commit.Tree.GetSHA()
	gitTree, _, err := api.Client.Git.GetTree(context.Background(), owner, repo, treeSHA, true)
	if err != nil {
		return nil, fmt.Errorf("error fetching tree: %v", err)
	}

	// Get the touched files from the PR
	touched, _, err := api.Client.PullRequests.ListFiles(context.Background(), owner, repo, pullNumber, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching touched files: %v", err)
	}

	// Merge the touched files list with the git tree to yield final desired tree
	statuses := createStatusMap(touched)
	var tree []models.FileTreeNode
	for _, entry := range gitTree.Entries {
		status := statuses[entry.GetPath()]
		if status == "" {
			status = "unmodified"
		}

		tree = append(tree, models.FileTreeNode{
			Status: status,
			Entry:  entry,
		})
	}

	return tree, nil
}

func (api *AppAPI) GetFileBlob(owner string, repo string, sha string) ([]byte, error) {
	contents, _, err := api.Client.Git.GetBlobRaw(context.Background(), owner, repo, sha)
	if err != nil {
		return nil, fmt.Errorf("error fetching contents: %v", err)
	}
	return contents, nil
}

func (api *AppAPI) CreateTeam(ctx context.Context, org_name, team_name string) (*github.Team, error) {
	team := &github.NewTeam{
		Name: team_name,
	}

	createdTeam, _, err := api.Client.Teams.CreateTeam(ctx, org_name, *team)
	if err != nil {
		return nil, fmt.Errorf("error creating team: %v", err)
	}

	return createdTeam, nil
}

func (api *AppAPI) AddTeamMember(ctx context.Context, team_id int64, user_name string, opt *github.TeamAddTeamMembershipOptions) error {
	_, _, err := api.Client.Teams.AddTeamMembership(ctx, team_id, user_name, opt)
	if err != nil {
		return fmt.Errorf("error adding member to team: %v", err)
	}

	return nil
}

func (api *AppAPI) AssignPermissionToTeam(ctx context.Context, team_id int64, owner_name string, repo_name string, permission string) error {
	opt := &github.TeamAddTeamRepoOptions{
		Permission: permission,
	}

	_, err := api.Client.Teams.AddTeamRepo(ctx, team_id, owner_name, repo_name, opt)
	if err != nil {
		return fmt.Errorf("error assigning permission to team: %v", err)
	}

	return nil
}

func (api *AppAPI) AssignPermissionToUser(ctx context.Context, ownerName string, repoName string, userName string, permission string) error {
	opt := &github.RepositoryAddCollaboratorOptions{
		Permission: permission,
	}

	_, err := api.Client.Repositories.AddCollaborator(ctx, ownerName, repoName, userName, opt)
	if err != nil {
		return fmt.Errorf("error assigning permission to user: %v", err)
	}

	return nil
}
