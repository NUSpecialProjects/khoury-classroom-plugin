package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type Service struct {
	store  storage.Storage
	github github.GitHubApp
}

func newService(store storage.Storage, github github.GitHubApp) *Service {
	return &Service{store: store, github: github}
}
