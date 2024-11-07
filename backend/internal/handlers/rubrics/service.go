package rubrics

import (
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type RubricService struct {
	store   storage.Storage
}

func newRubricService(store storage.Storage) *RubricService {
	return &RubricService{store: store}
}
