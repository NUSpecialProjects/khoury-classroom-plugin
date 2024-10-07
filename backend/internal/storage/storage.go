package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	Test
	Classrooms
  Assignments
}

type Test interface {
	GetTests(ctx context.Context) ([]models.Test, error)
}

type Classrooms interface {
	UsersIn(ctx context.Context, classroomID string) ([]models.User, error)
	GetAllClassrooms(ctx context.Context) ([]models.Classroom, error)
	CreateClassroom(ctx context.Context, classData models.Classroom) error
	CreateRegrade(ctx context.Context, classData models.Regrade) error
}

type Assignments interface {
  GetAllAssignmentTemplates(ctx context.Context) ([]models.AssignmentTemplate, error)
}
