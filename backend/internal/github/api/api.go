package api

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/google/go-github/v65/github"
)

type API struct {
	client *github.Client
}

func New(config *config.GitHub) (*API, error) {
	//client:= github.NewClient()
	//fmt.Println(client)
	//@ Nick and Kenny implement here
	return &API{
		client: nil,
	}, nil

}

//Stub, does not necessarily need to be implemented
func (api *API) Ping() error {
	return nil
}
