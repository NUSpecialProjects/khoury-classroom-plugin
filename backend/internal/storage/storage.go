package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	Test
	Assignments
	StudentAssignments
}

type Test interface {
	GetTests(ctx context.Context) ([]models.Test, error)
}

type StudentAssignments interface {
	GetTotalStudentAssignments(ctx context.Context, classroomID int64, assignmentID int64) (int64, error)
	GetStudentAssignments(ctx context.Context, classroomID int64, assignmentID int64) ([]models.StudentAssignment, error)
	GetStudentAssignment(ctx context.Context, classroomID int64, assignmentID int64, studentAssignmentID int64) (models.StudentAssignment, error)
}

type Assignments interface {
	GetAssignmentsInSemester(ctx context.Context, classroomID int64) ([]models.Assignment, error)
	GetAssignment(ctx context.Context, classroomID int64, localAssignmentID int64) (models.Assignment, error)
	CreateRubric(ctx context.Context, rubricData models.Rubric) error
	CreateAssignment(x context.Context, assignmentData models.Assignment) error
	CreateStudentAssignment(ctx context.Context, studentAssignmentData models.StudentAssignment) error
	CreateDueDate(ctx context.Context, dueDateData models.DueDate) error
	CreateRegrade(ctx context.Context, regradeData models.Regrade) error
	CreateSession(ctx context.Context, sessionData models.Session) error
	GetSession(ctx context.Context, gitHubUserID int64) (models.Session, error)
	DeleteSession(ctx context.Context, gitHubUserID int64) error
	GetAssignmentIDs(ctx context.Context) ([]models.Assignment_Classroom_ID, error)
}
