package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/storage"
)

type ClassroomService struct {
	store storage.Storage
}

func newClassroomService(store storage.Storage) *ClassroomService {
	return &ClassroomService{store: store}

}
