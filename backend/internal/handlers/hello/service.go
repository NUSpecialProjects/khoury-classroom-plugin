package hello

import "github.com/CamPlume1/khoury-classroom/internal/storage"

type Service struct {
	store storage.Storage
}

func newService(store storage.Storage) *Service {
	return &Service{store: store}
}
