package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type ClassroomService struct {
	store     storage.Storage
	appClient github.GitHubAppClient
	userCfg   *config.GitHubUserClient
	middleware.RoleChecker[ClassroomService]
}

func newClassroomService(
	store storage.Storage,
	appClient github.GitHubAppClient,
	userCfg *config.GitHubUserClient,
) *ClassroomService {
	service := &ClassroomService{store: store, appClient: appClient, userCfg: userCfg}
	service.RoleChecker = middleware.RoleChecker[ClassroomService]{Checkable: service}
	return service
}

// Getter for store field
func (s *ClassroomService) GetStore() storage.Storage {
	return s.store
}

// Getter for appClient field
func (s *ClassroomService) GetAppClient() github.GitHubAppClient {
	return s.appClient
}

// Getter for userCfg field
func (s *ClassroomService) GetUserCfg() *config.GitHubUserClient {
	return s.userCfg
}
