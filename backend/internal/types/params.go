package types

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type Params struct {
	AuthHandler        config.AuthHandler
	GitHubAppConfig    config.GitHubApp
	GitHubClientConfig config.GitHubClient
	Store              storage.Storage
	GitHubApp          github.GitHubApp
	GitHubClient       github.GitHubClient
}
