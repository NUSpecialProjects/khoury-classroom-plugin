package types

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/CamPlume1/khoury-classroom/internal/github"
)

type Params struct {
	AuthHandler config.AuthHandler
	Store    storage.Storage
	Github   github.Github
}
