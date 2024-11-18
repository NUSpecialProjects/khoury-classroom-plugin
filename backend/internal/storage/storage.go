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
	AssignmentOutline
	AssignmentTemplate
}

type Works interface {
	GetWorks(ctx context.Context, classroomID int, assignmentID int) ([]*models.StudentWorkWithContributors, error)
	GetWork(ctx context.Context, classroomID int, assignmentID int, studentWorkID int) (*models.PaginatedStudentWorkWithContributors, error)
	CreateStudentWork(ctx context.Context, work *models.StudentWork, GHUserID int64) error
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
	AddUserToClassroom(ctx context.Context, classroomID int64, classroomRole string, userID int64) (int64, error)
	ModifyUserRole(ctx context.Context, classroomID int64, classroomRole string, userID int64) error
	GetUsersInClassroom(ctx context.Context, classroomID int64) ([]models.UserWithRole, error)
	GetUserInClassroom(ctx context.Context, classroomID int64, userID int64) (models.UserWithRole, error)
	GetClassroomsInOrg(ctx context.Context, orgID int64) ([]models.Classroom, error)
	CreateClassroomToken(ctx context.Context, tokenData models.ClassroomToken) (models.ClassroomToken, error)
	GetClassroomToken(ctx context.Context, token string) (models.ClassroomToken, error)
}

type User interface {
	CreateUser(ctx context.Context, userToCreate models.User) (models.User, error)
	GetUserByGitHubID(ctx context.Context, githubUserID int64) (models.User, error)
}

type AssignmentOutline interface {
	GetAssignmentsInClassroom(ctx context.Context, classroomID int64) ([]models.AssignmentOutline, error)
	GetAssignmentByID(ctx context.Context, assignmentID int64) (models.AssignmentOutline, error)
	CreateAssignment(ctx context.Context, assignmentData models.AssignmentOutline) (models.AssignmentOutline, error)

	GetAssignmentWithTemplateByAssignmentID(ctx context.Context, assignmentID int64) (models.AssignmentOutlineWithTemplate, error)
	GetAssignmentByToken(ctx context.Context, token string) (models.AssignmentOutline, error)
	CreateAssignmentToken(ctx context.Context, tokenData models.AssignmentToken) (models.AssignmentToken, error)
}

type AssignmentTemplate interface {
	AssignmentTemplateExists(ctx context.Context, templateID int64) (bool, error)
	CreateAssignmentTemplate(ctx context.Context, assignmentTemplateData models.AssignmentTemplate) (models.AssignmentTemplate, error)
	GetAssignmentTemplateByID(ctx context.Context, templateID int64) (models.AssignmentTemplate, error)
}
