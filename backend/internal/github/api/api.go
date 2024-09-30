package api

import (
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v65/github"
)

type API struct {
	client *github.Client
}

func New(config *config.GitHub) (*API, error) {
	tr, err := ghinstallation.NewKeyFromFile(http.DefaultTransport, config.AppID, config.InstallationID, "private-key.pem")
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
	return nil
}
