package webhooks

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type WebHookService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
}

func newWebHookService(
	store storage.Storage,
	githubappclient github.GitHubAppClient,
) *WebHookService {
	return &WebHookService{
		store:           store,
		githubappclient: githubappclient,
	}
}
