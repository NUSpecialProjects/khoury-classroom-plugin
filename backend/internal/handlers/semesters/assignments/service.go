package assignments

import (
  "github.com/CamPlume1/khoury-classroom/internal/storage"
)


type AssignmentsService struct {
  store storage.Storage
}


func newAssignmentService(store storage.Storage) *AssignmentsService {
  return &AssignmentsService{store: store}

}
