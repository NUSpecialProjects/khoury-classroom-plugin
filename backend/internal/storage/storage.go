package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	Works
	Test
	Session
}

type Works interface {
	GetWorks(ctx context.Context, classroomID int, assignmentID int) ([]models.FormattedStudentWork, error)
	GetWork(ctx context.Context, classroomID int, assignmentID int, studentWorkID int) (*models.FormattedPaginatedStudentWork, error)
}

type Test interface {
	GetTests(ctx context.Context) ([]models.Test, error)
}

type Session interface {
	CreateSession(ctx context.Context, sessionData models.Session) error
	GetSession(ctx context.Context, gitHubUserID int64) (models.Session, error)
	DeleteSession(ctx context.Context, gitHubUserID int64) error
}
