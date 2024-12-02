package webhooks

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type WebHookService struct {
	store     storage.Storage
	appClient github.GitHubAppClient
}

func newWebHookService(
	store storage.Storage,
	appClient github.GitHubAppClient,
) *WebHookService {
	return &WebHookService{
		store:     store,
		appClient: appClient,
	}
}
