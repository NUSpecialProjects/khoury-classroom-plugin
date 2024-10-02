package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type Service struct {
	store  storage.Storage
	github github.GithubApp
}

func newService(store storage.Storage, github github.GithubApp) *Service {
	return &Service{store: store, github: github}
}
