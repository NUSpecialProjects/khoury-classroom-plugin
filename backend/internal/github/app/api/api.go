package github

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	gh "github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/google/go-github/v52/github"
	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"
)

type API struct {
	gh.GitHubBase
	client *github.Client
}

func New(cfg *config.GitHubApp) (*API, error) {
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

	return &API{
		client: githubClient,
	}, nil
}
