package test

import "github.com/CamPlume1/khoury-classroom/internal/storage"

type TestService struct {
	store storage.Storage
}


func newTestService(store storage.Storage) *TestService {
	return &TestService{store: store}
}