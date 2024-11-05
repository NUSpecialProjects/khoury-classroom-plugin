package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type ClassroomService struct {
	store   storage.Storage
	userCfg *config.GitHubUserClient
}

func newClassroomService(store storage.Storage, userCfg *config.GitHubUserClient) *ClassroomService {
	return &ClassroomService{store: store, userCfg: userCfg}
}
