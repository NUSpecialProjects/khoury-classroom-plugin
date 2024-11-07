package assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type AssignmentService struct {
	userCfg         *config.GitHubUserClient
	store storage.Storage
}

func NewAssignmentService(store storage.Storage, userCfg *config.GitHubUserClient) *AssignmentService {
	return &AssignmentService{store: store, userCfg: userCfg}
}
