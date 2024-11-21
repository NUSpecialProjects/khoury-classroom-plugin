package storage

import (
	"context"

	"github.com/CamPlume1/khoury-classroom/internal/models"
)

type Storage interface {
	Close(context.Context)
	FeedbackComment
	Works
	Test
	Session
	Classroom
	User
	AssignmentOutline
	Rubric
	AssignmentTemplate
}

type FeedbackComment interface {
	GetFeedbackOnWork(ctx context.Context, studentWorkID int) ([]models.PRReviewCommentResponse, error)
	CreateNewFeedbackComment(ctx context.Context, TAUserID int64, studentWorkID int, comment models.PRReviewCommentResponse) error
	AttachRubricItemToWork(ctx context.Context, studentWorkID int, path string, line int, rubricItemID int) error
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
	AddUserToClassroom(ctx context.Context, classroomID int64, classroomRole string, classroomStatus models.UserStatus, userID int64) (models.ClassroomUser, error)
	RemoveUserFromClassroom(ctx context.Context, classroomID int64, userID int64) error
	ModifyUserRole(ctx context.Context, classroomID int64, classroomRole string, userID int64) (models.ClassroomUser, error)
	ModifyUserStatus(ctx context.Context, classroomID int64, status models.UserStatus, userID int64) (models.ClassroomUser, error)
	GetUsersInClassroom(ctx context.Context, classroomID int64) ([]models.ClassroomUser, error)
	GetUserInClassroom(ctx context.Context, classroomID int64, userID int64) (models.ClassroomUser, error)
	GetClassroomsInOrg(ctx context.Context, orgID int64) ([]models.Classroom, error)
	GetUserClassroomsInOrg(ctx context.Context, orgID int64, userID int64) ([]models.Classroom, error)
	CreateClassroomToken(ctx context.Context, tokenData models.ClassroomToken) (models.ClassroomToken, error)
	GetClassroomToken(ctx context.Context, token string) (models.ClassroomToken, error)
}

type User interface {
	CreateUser(ctx context.Context, userToCreate models.User) (models.User, error)
	GetUserByGitHubID(ctx context.Context, githubUserID int64) (models.User, error)
	GetUserByID(ctx context.Context, userID int64) (models.User, error)
}

type AssignmentOutline interface {
	GetAssignmentsInClassroom(ctx context.Context, classroomID int64) ([]models.AssignmentOutline, error)
	GetAssignmentByID(ctx context.Context, assignmentID int64) (models.AssignmentOutline, error)
	CreateAssignment(ctx context.Context, assignmentData models.AssignmentOutline) (models.AssignmentOutline, error)
    UpdateAssignmentRubric(ctx context.Context, rubricID int64, assignmentID int64) (models.AssignmentOutline, error)
	GetAssignmentWithTemplateByAssignmentID(ctx context.Context, assignmentID int64) (models.AssignmentOutlineWithTemplate, error)
	GetAssignmentByToken(ctx context.Context, token string) (models.AssignmentOutline, error)
	CreateAssignmentToken(ctx context.Context, tokenData models.AssignmentToken) (models.AssignmentToken, error)
}

type AssignmentTemplate interface {
	AssignmentTemplateExists(ctx context.Context, templateID int64) (bool, error)
	CreateAssignmentTemplate(ctx context.Context, assignmentTemplateData models.AssignmentTemplate) (models.AssignmentTemplate, error)
	GetAssignmentTemplateByID(ctx context.Context, templateID int64) (models.AssignmentTemplate, error)
}

type Rubric interface {
	CreateRubric(ctx context.Context, rubricData models.Rubric) (models.Rubric, error)
	GetRubric(ctx context.Context, rubricID int64) (models.Rubric, error)
	AddItemToRubric(ctx context.Context, rubricItemData models.RubricItem) (models.RubricItem, error)
	GetRubricItems(ctx context.Context, rubricID int64) ([]models.RubricItem, error)
    UpdateRubric(ctx context.Context, rubricID int64, rubricData models.Rubric) (models.Rubric, error)
    UpdateRubricItem(ctx context.Context, rubricItemData models.RubricItem) (models.RubricItem, error)
}
