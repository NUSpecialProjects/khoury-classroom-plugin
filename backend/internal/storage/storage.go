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
	GetStudentAssignment(ctx context.Context, classroomID int64, assignmentID int64, studentAssignmentID int64) (models.StudentAssignment, error)
	CreateDueDate(ctx context.Context, dueDateData models.DueDate) error
	CreateRegrade(ctx context.Context, regradeData models.Regrade) error
	CreateSession(ctx context.Context, sessionData models.Session) error
	GetSession(ctx context.Context, gitHubUserID int64) (models.Session, error)
	DeleteSession(ctx context.Context, gitHubUserID int64) error
}

type Semesters interface {
	ListSemestersByOrgList(ctx context.Context, orgIDs []int64) ([]models.Semester, error)
	ListSemestersByOrg(ctx context.Context, orgID int64) ([]models.Semester, error)
	CreateSemester(ctx context.Context, semesterData models.Semester) (models.Semester, error)
	GetSemester(ctx context.Context, classroomID int64) (models.Semester, error)
	DeactivateSemester(ctx context.Context, classroomID int64) (models.Semester, error)
	ActivateSemester(ctx context.Context, classroomID int64) (models.Semester, error)
}
