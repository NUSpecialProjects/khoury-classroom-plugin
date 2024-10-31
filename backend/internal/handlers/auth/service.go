package auth

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type AuthService struct {
	store   storage.Storage
	userCfg *config.GitHubUserClient
}

func newAuthService(
	store storage.Storage,
	userCfg *config.GitHubUserClient,
) *AuthService {
	return &AuthService{
		store:   store,
		userCfg: userCfg,
	}
}
