package api

import (
	"context"
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/google/go-github/v65/github"
	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"
)

type API struct {
	client *github.Client
}

func New(config *config.GitHub) (*API, error) {
	privateKey := []byte(config.Key)

	appTokenSource, err := githubauth.NewApplicationTokenSource(1112, privateKey)
	if err != nil {
		fmt.Println("Error creating application token source:", err)
	}

	installationTokenSource := githubauth.NewInstallationTokenSource(1113, appTokenSource)
	httpClient := oauth2.NewClient(context.Background(), installationTokenSource)
	client := github.NewClient(httpClient)

	return &API{
		client,
	}, nil

}

// Stub, does not necessarily need to be implemented
func (api *API) Ping() error {
	println(api.client.RateLimit.Get(context.Background()))
	return nil
}
