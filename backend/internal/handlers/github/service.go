package github

import (
	githubapp "github.com/CamPlume1/khoury-classroom/internal/github/app"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type Service struct {
	store  storage.Storage
	github githubapp.GitHubApp
}

func newService(store storage.Storage, github githubapp.GitHubApp) *Service {
	return &Service{store: store, github: github}
}
