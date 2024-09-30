package types

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type Params struct {
	Supabase config.Database
	Store    storage.Storage
}
