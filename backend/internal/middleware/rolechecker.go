package middleware

import (
	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/CamPlume1/khoury-classroom/internal/storage"
	"github.com/gofiber/fiber/v2"
)

type Checkable interface {
	GetStore() storage.Storage
	GetUserCfg() *config.GitHubUserClient
	GetAppClient() github.GitHubAppClient
}

type RoleChecker[T any] struct {
	Checkable
}

// Helper function to check if the user has at least the role specified
func (roleChecker *RoleChecker[T]) RequireAtLeastRole(c *fiber.Ctx, classroomID int64, role models.ClassroomRole) (models.ClassroomUser, error) {
	return roleChecker.checkRole(c, classroomID, role, func(userRole, requiredRole models.ClassroomRole) bool {
		return userRole.Compare(requiredRole) < 0
	})
}

func (roleChecker *RoleChecker[T]) RequireGreaterThanRole(c *fiber.Ctx, classroomID int64, role models.ClassroomRole) (models.ClassroomUser, error) {
	return roleChecker.checkRole(c, classroomID, role, func(userRole, requiredRole models.ClassroomRole) bool {
		return userRole.Compare(requiredRole) <= 0
	})
}

// Helper function containing shared role checking logic
func (roleChecker *RoleChecker[T]) checkRole(c *fiber.Ctx, classroomID int64, role models.ClassroomRole, failCheck func(models.ClassroomRole, models.ClassroomRole) bool) (models.ClassroomUser, error) {
	_, _, user, err := GetClientAndUser(c, roleChecker.GetStore(), roleChecker.GetUserCfg())
	if err != nil {
		return models.ClassroomUser{}, errs.AuthenticationError()
	}

	classroom, err := roleChecker.GetStore().GetClassroomByID(c.Context(), classroomID)
	if err != nil {
		return models.ClassroomUser{}, errs.InternalServerError()
	}

	classroomUser, err := roleChecker.GetStore().GetUserInClassroom(c.Context(), classroomID, *user.ID)
	if err != nil {
		return models.ClassroomUser{}, errs.InternalServerError()
	}

	// Check if user has sufficient role using provided comparison function
	if failCheck(classroomUser.Role, role) {
		return models.ClassroomUser{}, errs.InsufficientPermissionsError()
	}

	// if the user is removed or not in the org, they are not in the classroom
	if classroomUser.Status == models.UserStatusRemoved || classroomUser.Status == models.UserStatusNotInOrg {
		return models.ClassroomUser{}, errs.UserNotFoundInClassroomError()
	}

	// if the user is a student, check if they are in the student team
	if classroomUser.Role == models.Student {
		studentTeam, err := roleChecker.GetAppClient().GetTeamByName(c.Context(), classroom.OrgName, *classroom.StudentTeamName)
		if err != nil { // student team doesn't exist :(
			return models.ClassroomUser{}, errs.InternalServerError()
		} else { // student team exists, check if the user is in it
			var studentIsInStudentTeam = false
			studentTeamMembers, err := roleChecker.GetAppClient().GetTeamMembers(c.Context(), *studentTeam.ID)
			if err != nil {
				return models.ClassroomUser{}, errs.InternalServerError()
			}
			for _, member := range studentTeamMembers {
				if *member.Login == user.GithubUsername {
					studentIsInStudentTeam = true
				}
			}
			if !studentIsInStudentTeam {
				return models.ClassroomUser{}, errs.StudentNotInStudentTeamError()
			}
		}
	}

	return classroomUser, nil
}
