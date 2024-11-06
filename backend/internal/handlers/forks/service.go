package forks

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type ForkService struct {
	userCfg         *config.GitHubUserClient
	store storage.Storage
}

func NewForkService(userCfg  *config.GitHubUserClient, store storage.Storage) *ForkService {
	return &ForkService{userCfg: userCfg, store: store}
}
