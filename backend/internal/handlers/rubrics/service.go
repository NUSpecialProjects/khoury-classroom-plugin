package rubrics

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type RubricService struct {
	store   storage.Storage
	userCfg *config.GitHubUserClient
}

func newRubricService(store storage.Storage, userCfg *config.GitHubUserClient) *RubricService {
    return &RubricService{store: store, userCfg: userCfg}
}
