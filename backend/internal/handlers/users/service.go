package users

import (
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type UserService struct {
	store storage.Storage
}

func newUserService(store storage.Storage) *UserService {
	return &UserService{store: store}
}
