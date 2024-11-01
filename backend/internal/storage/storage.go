package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	Test
	Session
    Classroom
}

type Test interface {
	GetTests(ctx context.Context) ([]models.Test, error)
}

type Session interface {
	CreateSession(ctx context.Context, sessionData models.Session) error
	GetSession(ctx context.Context, gitHubUserID int64) (models.Session, error)
	DeleteSession(ctx context.Context, gitHubUserID int64) error
}

type Classroom interface {
    CreateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error)
    UpdateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error)
    GetClassroomByID(ctx context.Context, classroomID int64) (models.Classroom, error)
}
