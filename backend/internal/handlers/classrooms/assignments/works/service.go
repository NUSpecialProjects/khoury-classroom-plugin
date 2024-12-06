package works

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type WorkService struct {
	store     storage.Storage
	userCfg   *config.GitHubUserClient
	appClient github.GitHubAppClient
	middleware.RoleChecker[WorkService]
}

func NewWorkService(store storage.Storage, userCfg *config.GitHubUserClient, appClient github.GitHubAppClient) *WorkService {
	return &WorkService{store: store, userCfg: userCfg, appClient: appClient}
}

// Getter for store field
func (s *WorkService) GetStore() storage.Storage {
	return s.store
}

// Getter for userCfg field
func (s *WorkService) GetUserCfg() *config.GitHubUserClient {
	return s.userCfg
}

// Getter for appClient field
func (s *WorkService) GetAppClient() github.GitHubAppClient {
	return s.appClient
}
