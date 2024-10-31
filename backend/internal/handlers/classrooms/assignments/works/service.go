package works

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type WorkService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func NewWorkService(store storage.Storage, githubappclient github.GitHubAppClient) *WorkService {
	return &WorkService{store: store, githubappclient: githubappclient}
}
