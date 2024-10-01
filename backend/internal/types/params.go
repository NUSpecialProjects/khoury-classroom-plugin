package types

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type Params struct {
	AuthHandler       config.AuthHandler
	GithubAuthHandler config.GithubAuthHandler
	Store             storage.Storage
	Github            github.Github
}
