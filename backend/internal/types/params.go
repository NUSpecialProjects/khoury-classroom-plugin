package types

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	githubapp "github.com/CamPlume1/khoury-classroom/internal/github/app"
	githubclient "github.com/CamPlume1/khoury-classroom/internal/github/client"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type Params struct {
	AuthHandler     config.AuthHandler
	GithubAppConfig config.GitHubApp
	Store           storage.Storage
	GithubApp       githubapp.GitHubApp
	GithubClient    githubclient.GitHubClient
}
