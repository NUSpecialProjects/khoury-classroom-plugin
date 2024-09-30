package api

import (
	"context"
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v65/github"
)

type API struct {
	client *github.Client
}

func New(config *config.GitHub) (*API, error) {
	tr, err := ghinstallation.New(http.DefaultTransport, config.AppID, config.InstallationID, []byte(config.Key))
	if err != nil {
		println(err.Error())
	}
	client := github.NewClient(&http.Client{Transport: tr})

	return &API{
		client,
	}, nil

}

// Stub, does not necessarily need to be implemented
func (api *API) Ping() error {
	println(api.client.RateLimit.Get(context.Background()))
	return nil
}

func (api *API) ListRepos() []string {
	repos, _, repoErr := api.client.Repositories.ListByOrg(context.Background(), "NUSpecialProjects", nil)
	if repoErr != nil {
		// Handle error.
		println(repoErr.Error())
		return nil
	} else {
		repoNames := []string{}
		for _, repo := range repos {
			repoNames = append(repoNames, *repo.Name)
		}
		return repoNames
	}
}
