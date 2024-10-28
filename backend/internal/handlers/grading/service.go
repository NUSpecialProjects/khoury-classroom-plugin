package grading

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type GradingService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func newGradingService(store storage.Storage, githubappclient github.GitHubAppClient) *GradingService {
	return &GradingService{store: store, githubappclient: githubappclient}
}
