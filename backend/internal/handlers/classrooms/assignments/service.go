package assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type AssignmentService struct {
	store storage.Storage
}

func NewAssignmentService(store storage.Storage) *AssignmentService {
	return &AssignmentService{store: store}

}
