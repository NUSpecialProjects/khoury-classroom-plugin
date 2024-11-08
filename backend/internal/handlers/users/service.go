package users

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type UserService struct {
	store   storage.Storage
	userCfg *config.GitHubUserClient
}

func newUserService(store storage.Storage, userCfg *config.GitHubUserClient) *UserService {
	return &UserService{store: store, userCfg: userCfg}
}
