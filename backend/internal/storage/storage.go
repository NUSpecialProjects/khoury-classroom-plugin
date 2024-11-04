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
	Classroom
	User
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

type Classroom interface {
	CreateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error)
	UpdateClassroom(ctx context.Context, classroomData models.Classroom) (models.Classroom, error)
	GetClassroomByID(ctx context.Context, classroomID int64) (models.Classroom, error)
	AddUserToClassroom(ctx context.Context, classroomID int64, userID int64) (int64, error)
}

type User interface {
	CreateUser(ctx context.Context, userToCreate models.User) (models.User, error)
}
