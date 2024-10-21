package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	Test
	Assignments
	Semesters
}

type Test interface {
	GetTests(ctx context.Context) ([]models.Test, error)
}

type Assignments interface {
	GetAllAssignments(ctx context.Context) ([]models.Assignment, error)
	CreateRubric(ctx context.Context, rubricData models.Rubric) error
	CreateAssignment(x context.Context, assignmentData models.Assignment) error
	CreateStudentAssignment(ctx context.Context, studentAssignmentData models.StudentAssignment) error
	CreateDueDate(ctx context.Context, dueDateData models.DueDate) error
	CreateRegrade(ctx context.Context, regradeData models.Regrade) error
	CreateSession(ctx context.Context, sessionData models.Session) error
	GetSession(ctx context.Context, gitHubUserID int64) (models.Session, error)
	DeleteSession(ctx context.Context, gitHubUserID int64) error
  GetAssignmentIDs(ctx context.Context) ([]models.Assignment_CR_ID, error)
  
}

type Semesters interface {
	ListSemesters(ctx context.Context, orgID int64) ([]models.Semester, error)
	CreateSemester(ctx context.Context, semesterData models.Semester) error
	GetSemester(ctx context.Context, semesterID int64) (models.Semester, error)
	DeactivateSemester(ctx context.Context, semesterID int64) error
	ActivateSemester(ctx context.Context, semesterID int64) error
}
