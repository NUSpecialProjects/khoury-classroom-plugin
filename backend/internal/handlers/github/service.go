package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type GitHubService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
	sessionmanager  *session.Store
}

func newGitHubService(
	store storage.Storage,
	githubappclient github.GitHubAppClient,
	sessionmanager *session.Store,
) *GitHubService {
	return &GitHubService{
		store:           store,
		githubappclient: githubappclient,
		sessionmanager:  sessionmanager,
	}
}
