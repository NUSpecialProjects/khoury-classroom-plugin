package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Service struct {
	store          storage.Storage
	githubapp      github.GitHubAppClient
	sessionmanager *session.Store
}

func newService(store storage.Storage, githubapp github.GitHubAppClient, sessionmanager *session.Store) *Service {
	return &Service{store: store, githubapp: githubapp, sessionmanager: sessionmanager}
}
