package postgres

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/google/go-github/v65/github"	
)

type API struct {
	client *github.Client
}

func New(config config.GitHub) (*API, error) {
	client:= github.NewClient()
	//@ Nick and Kenny implement here

}
