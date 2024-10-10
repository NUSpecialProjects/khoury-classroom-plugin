package types

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Params struct {
	UserCfg        config.GitHubUserClient
	SessionManager *session.Store
	Store          storage.Storage
	GitHubApp      github.GitHubAppClient
}
