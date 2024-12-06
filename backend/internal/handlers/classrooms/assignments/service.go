package assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type AssignmentService struct {
	store storage.Storage
	userCfg *config.GitHubUserClient
	appClient github.GitHubAppClient
}

func NewAssignmentService(store storage.Storage, userCfg *config.GitHubUserClient, appClient github.GitHubAppClient) *AssignmentService {
	return &AssignmentService{store: store, userCfg: userCfg, appClient: appClient}
}
