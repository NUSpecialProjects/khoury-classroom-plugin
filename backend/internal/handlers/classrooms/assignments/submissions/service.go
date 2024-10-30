package submissions

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type SubmissionService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func NewSubmissionService(store storage.Storage, githubappclient github.GitHubAppClient) *SubmissionService {
	return &SubmissionService{store: store, githubappclient: githubappclient}
}
