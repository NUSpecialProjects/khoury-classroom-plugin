package appclient

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/sharedclient"
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
func (api *AppAPI) GetJWT(ctx context.Context) (string, error) {
	token, err := api.appTokenSource.Token()
	if err != nil {
		return "", fmt.Errorf("error getting token: %v", err)
	}
	return token.AccessToken, nil
}

func (api *AppAPI) GetClientWithJWTAuth(ctx context.Context) (*github.Client, error) {
	// Create a new OAuth2 client with the JWT
	token, err := api.GetJWT(ctx)
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
	client, err := api.GetClientWithJWTAuth(ctx)
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

func (api *AppAPI) GetGitTree(owner string, repo string) ([]github.TreeEntry, error) {
	// Get the reference to the branch
	ref, _, err := api.Client.Git.GetRef(context.Background(), owner, repo, "heads/main")
	if err != nil {
		return nil, fmt.Errorf("error fetching branch ref: %v", err)
	}

	// Get the commit from the ref
	commitSHA := ref.Object.GetSHA()
	commit, _, err := api.Client.Git.GetCommit(context.Background(), owner, repo, commitSHA)
	if err != nil {
		return nil, fmt.Errorf("error fetching commit: %v", err)
	}

	treeSHA := commit.Tree.GetSHA()
	tree, _, err := api.Client.Git.GetTree(context.Background(), owner, repo, treeSHA, true)
	if err != nil {
		return nil, fmt.Errorf("error fetching tree: %v", err)
	}

	return tree.Entries, nil
}

func (api *AppAPI) GetGitBlob(owner string, repo string, sha string) ([]byte, error) {
	contents, _, err := api.Client.Git.GetBlobRaw(context.Background(), owner, repo, sha)
	if err != nil {
		return nil, fmt.Errorf("error fetching contents: %v", err)
	}
	return contents, nil
}
