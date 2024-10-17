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
	webhooksecret string
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
		webhooksecret: cfg.WebhookSecret,
	}, nil
}

// Any APP specific implementations can go here
func (api *AppAPI) GetWebhookSecret() string {
	return api.webhooksecret
}

func (api *AppAPI) GetStudentAssignmentFiles(assignmentID int32, studentAssignmentID int32, path string) ([]models.StudentAssignmentFiles, error) {
	endpoint := fmt.Sprintf("/repos/%s/%s/contents/%s", owner, repo, path)

	var files []models.StudentAssignmentFiles

	api.Client.Repositories.GetContents()

	// Create a new GET request
	req, err := api.Client.NewRequest("GET", endpoint, nil)
	if err != nil {
		return files, fmt.Errorf("error creating request: %v", err)
	}

	// Make the API call
	_, err = api.Client.Do(context.Background(), req, &files)
	if err != nil {
		return files, fmt.Errorf("error fetching assignment files: %v", err)
	}

	return files, nil
}
