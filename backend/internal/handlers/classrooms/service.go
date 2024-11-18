package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type ClassroomService struct {
	store           storage.Storage
	githubappclient github.GitHubAppClient
	userCfg         *config.GitHubUserClient
}

func newClassroomService(
	store storage.Storage,
	githubappclient github.GitHubAppClient,
	userCfg *config.GitHubUserClient,
) *ClassroomService {
	return &ClassroomService{store: store, githubappclient: githubappclient, userCfg: userCfg}
}
