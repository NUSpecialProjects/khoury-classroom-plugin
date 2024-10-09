package assignments

import (
  "github.com/CamPlume1/khoury-classroom/internal/storage"
)


type UsersService struct {
  store storage.Storage
}


func newUsersService(store storage.Storage) *UsersService {
  return &UsersService{store: store}

}
