package classrooms

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/github"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/CamPlume1/khoury-classroom/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// Returns the classrooms the authenticated user is part of.
func (s *ClassroomService) getUserClassrooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Returns the details of a classroom.
func (s *ClassroomService) getClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		// Only allow TAs and Profs to get classroom info
		_, err = s.RequireAtLeastRole(c, classroomID, models.TA)
		if err != nil {
			return err
		}

		classroomData, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"classroom": classroomData})
	}
}

// Creates a new classroom.
func (s *ClassroomService) createClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, githubUser, user, err := middleware.GetClientAndUser(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		var classroomData models.Classroom
		err = c.BodyParser(&classroomData)
		if err != nil {
			return errs.InvalidRequestBody(models.Classroom{})
		}

		membership, err := client.GetUserOrgMembership(c.Context(), classroomData.OrgName, githubUser.Login)
		if err != nil || *membership.Role != "admin" {
			return errs.InsufficientPermissionsError()
		}

		// Determine the team name for the classroom
		studentTeamName := strings.ReplaceAll(strings.ToLower(classroomData.Name), " ", "-") + "-students"
		classroomData.StudentTeamName = &studentTeamName

		// Handle existing student team
		existingTeam, err := client.GetTeamByName(c.Context(), classroomData.OrgName, studentTeamName)
		if err == nil && existingTeam != nil {
			// Team exists - delete it first
			err = client.DeleteTeam(c.Context(), *existingTeam.ID)
			if err != nil {
				return errs.InternalServerError()
			}
		}

		// Create the student team
		description := "The students of " + classroomData.OrgName + " - " + classroomData.Name + ".\n\nAutomatically generated by Khoury Classroom."
		maintainers := []string{githubUser.Login}
		_, err = client.CreateTeam(c.Context(), classroomData.OrgName, *classroomData.StudentTeamName, &description, maintainers)
		if err != nil {
			return errs.InternalServerError()
		}

		// Create the classroom
		createdClassroom, err := s.store.CreateClassroom(c.Context(), classroomData)
		if err != nil {
			return errs.InternalServerError()
		}

		// Add the user as a professor to the classroom
		_, err = s.store.AddUserToClassroom(c.Context(), createdClassroom.ID, string(models.Professor), models.UserStatusActive, *user.ID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"classroom": createdClassroom})
	}
}

// Updates an existing classroom.
func (s *ClassroomService) updateClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		var classroomData models.Classroom
		error := c.BodyParser(&classroomData)
		if error != nil {
			return errs.InvalidRequestBody(models.Classroom{})
		}
		classroomData.ID = classroomID

		// Only allow professors to update classrooms
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		updatedClassroom, err := s.store.UpdateClassroom(c.Context(), classroomData)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"classroom": updatedClassroom})
	}
}

// Updates the name of a classroom
func (s *ClassroomService) updateClassroomName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		var classroomData models.Classroom
		error := c.BodyParser(&classroomData)
		if error != nil {
			return errs.InvalidRequestBody(models.Classroom{})
		}
		classroomData.ID = classroomID

		// Only allow professors to update classroom names
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		existingClassroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}
		existingClassroom.Name = classroomData.Name

		updatedClassroom, err := s.store.UpdateClassroom(c.Context(), existingClassroom)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"classroom": updatedClassroom})
	}
}

// Returns the users of a classroom.
func (s *ClassroomService) getClassroomUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		// Only allow TAs and Profs to get users in classrooms
		_, err = s.RequireAtLeastRole(c, classroomID, models.TA)
		if err != nil {
			return err
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		usersInClassroom, err := s.store.GetUsersInClassroom(c.Context(), classroomID)
		if err != nil {

			return errs.InternalServerError()
		}

		updatedUsersInClassroom := []models.ClassroomUser{}

		for _, classroomUser := range usersInClassroom {
			newClassroomUser, err := s.updateUserStatus(c.Context(), client, classroomUser.User, classroom)
			// don't include members who are not in the org
			if newClassroomUser.Status == models.UserStatusRemoved {
				continue
			}
			if err != nil { // failed to update their status, so just keep the old one
				updatedUsersInClassroom = append(updatedUsersInClassroom, classroomUser)
			} else { // add the updated user to the list
				updatedUsersInClassroom = append(updatedUsersInClassroom, newClassroomUser)
			}
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{"users": updatedUsersInClassroom})
	}
}
func (s *ClassroomService) getRubricsInClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		rubrics, err := s.store.GetRubricsInClassroom(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		var fullRubrics []models.FullRubric
		for _, rubric := range rubrics {
			var fullRubric models.FullRubric
			fullRubric.Rubric = rubric

			items, err := s.store.GetRubricItems(c.Context(), rubric.ID)
			if err != nil {
				return errs.InternalServerError()
			}
			fullRubric.RubricItems = items

			fullRubrics = append(fullRubrics, fullRubric)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"full_rubrics": fullRubrics})
	}
}

// Removes a user from a classroom.
func (s *ClassroomService) removeUserFromClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		// Only allow professors to remove users from classrooms
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		studentTeam, err := s.appClient.GetTeamByName(c.Context(), classroom.OrgName, *classroom.StudentTeamName)
		if err != nil {
			return errs.InternalServerError()
		}

		toBeRemovedUser, err := s.store.GetUserByID(c.Context(), userID)
		if err != nil {
			return errs.InternalServerError()
		}

		// remove the user from the student team
		err = s.appClient.RemoveTeamMember(c.Context(), classroom.OrgName, *studentTeam.ID, toBeRemovedUser.GithubUsername)
		if err != nil {
			log.Default().Println("Warning: Failed to remove user from student team", err)
			// do nothing, the user has already been removed from the team or they were never in the team
		}

		err = s.store.RemoveUserFromClassroom(c.Context(), classroomID, userID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.SendStatus(http.StatusOK)
	}
}

// Generates a token to join a classroom.
func (s *ClassroomService) generateClassroomToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := models.ClassroomRoleRequestBody{}

		if err := c.BodyParser(&body); err != nil {
			return errs.InvalidRequestBody(body)
		}

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		classroomRole, err := models.NewClassroomRole(body.ClassroomRole)
		if err != nil {
			return errs.BadRequest(err)
		}

		// Only allow professors to invite people to classrooms
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		token, err := utils.GenerateToken(16)
		if err != nil {
			return errs.InternalServerError()
		}

		tokenData := models.ClassroomToken{
			ClassroomID:   classroomID,
			ClassroomRole: classroomRole,
			BaseToken: models.BaseToken{
				Token: token,
			},
		}

		// Set ExpiresAt only if Duration is provided
		if body.Duration != nil {
			expiresAt := time.Now().Add(time.Duration(*body.Duration) * time.Minute)
			tokenData.ExpiresAt = &expiresAt
		}

		classroomToken, err := s.store.CreateClassroomToken(c.Context(), tokenData)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"token": classroomToken.Token})
	}
}

// Uses a classroom token to request to join a classroom.
func (s *ClassroomService) useClassroomToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, _, user, err := middleware.GetClientAndUser(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		token := c.Params("token")
		if token == "" {
			return errs.MissingAPIParamError("token")
		}

		// Go get the token from the DB
		classroomToken, err := s.store.GetClassroomToken(c.Context(), token)
		if err != nil {
			return errs.AuthenticationError()
		}

		// Check if the token is valid
		if classroomToken.ExpiresAt != nil && classroomToken.ExpiresAt.Before(time.Now()) {
			return errs.ExpiredTokenError()
		}

		// Get the classroom from the DB
		classroom, err := s.store.GetClassroomByID(c.Context(), classroomToken.ClassroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroomUser, err := s.store.GetUserInClassroom(c.Context(), classroomToken.ClassroomID, *user.ID)
		if err != nil {
			classroomUser, err = s.store.AddUserToClassroom(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), models.UserStatusRequested, *user.ID)
			if err != nil {
				return errs.InternalServerError()
			}
		}

		classroomUser, err = s.updateUserStatus(c.Context(), client, user, classroom)
		if err != nil {
			return errs.InternalServerError()
		}

		// don't do anything if the user has been removed from the classroom
		if classroomUser.Status == models.UserStatusRemoved {
			return errs.InsufficientPermissionsError()
		}

		// user is already in the classroom. If their role can be upgraded, do so. Don't downgrade them.
		roleComparison := classroomUser.Role.Compare(classroomToken.ClassroomRole)
		if roleComparison < 0 {
			// Upgrade the user's role in the classroom
			classroomUser, err = s.store.ModifyUserRole(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), *classroomUser.ID)
			if err != nil {
				return errs.InternalServerError()
			}
		}

		// Invite the user to the organization
		// classroomUser, err = s.inviteUserToOrganization(c.Context(), s.appClient, classroom.OrgName, classroomToken.ClassroomID, classroomToken.ClassroomRole, user)
		classroomUser, err = s.inviteUserToOrganization(c.Context(), s.appClient, classroom, classroomToken.ClassroomRole, user)
		if err != nil {
			return errs.InternalServerError()
		}

		// Accept the pending invitation to the organization
		err = s.acceptOrgInvitation(c.Context(), client, classroom.OrgName, classroomToken.ClassroomID, user)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":   "Token applied successfully",
			"user":      classroomUser,
			"classroom": classroom,
		})
	}
}

// Returns the user's status in the classroom, nil if not in the classroom
func (s *ClassroomService) getCurrentClassroomUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, _, user, err := middleware.GetClientAndUser(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroomUser, err := s.updateUserStatus(c.Context(), client, user, classroom)
		if err != nil {
			if err == errs.UserNotFoundInClassroomError() {
				// User not found in classroom, return null
				return c.Status(http.StatusOK).JSON(fiber.Map{"user": nil})
			} else {
				return errs.InternalServerError()
			}
		}

		if classroomUser.Status == models.UserStatusNotInOrg || classroomUser.Status == models.UserStatusRemoved {
			return errs.InconsistentOrgMembershipError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"user": classroomUser})
	}
}

// Updates the user's status in our DB to reflect their org membership, as of this moment
func (s *ClassroomService) updateUserStatus(ctx context.Context, client github.GitHubUserClient, user models.User, classroom models.Classroom) (models.ClassroomUser, error) {
	classroomUser, err := s.store.GetUserInClassroom(ctx, classroom.ID, *user.ID)
	if err != nil {
		return models.ClassroomUser{}, errs.UserNotFoundInClassroomError()
	}

	// if the user has been removed from the classroom, don't update their org membership
	if classroomUser.Status == models.UserStatusRemoved {
		return classroomUser, nil
	}

	membership, err := client.GetUserOrgMembership(ctx, classroom.OrgName, user.GithubUsername)
	if err != nil && classroomUser.Status != models.UserStatusRequested { // if the user is in the requested state, we don't want to change their status
		// user isn't in the org, set them to NOT IN ORG (this probably means they have been removed from the org OR they denied their invite)
		classroomUser, err = s.store.ModifyUserStatus(ctx, classroom.ID, models.UserStatusNotInOrg, *user.ID)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
		return classroomUser, nil
	} else if membership != nil && *membership.State == "active" { // user is in the org, set them to active
		classroomUser, err = s.store.ModifyUserStatus(ctx, classroom.ID, models.UserStatusActive, *user.ID)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
	} else if membership != nil && *membership.State == "pending" { // user has a pending invitation, set them to invited
		classroomUser, err = s.store.ModifyUserStatus(ctx, classroom.ID, models.UserStatusOrgInvited, *user.ID)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
	}
	return classroomUser, nil
}

// Sends invites to all users in the classroom who are in the requested state
func (s *ClassroomService) sendOrganizationInvitesToRequestedUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroomRole, err := models.NewClassroomRole(c.Params("classroom_role"))
		if err != nil {
			return errs.BadRequest(err)
		}

		// Only allow professors to invite people to org
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		classroomUsers, err := s.store.GetUsersInClassroom(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		stillRequestedUsers := []models.ClassroomUser{}
		invitedUsers := []models.ClassroomUser{}

		for _, classroomUser := range classroomUsers {
			if classroomUser.Status != models.UserStatusRequested {
				continue
			}
			//TODO: these are many content generating requests to the GitHub API, maybe need to delay between them
			// use the current user's client to invite the user to the organization
			modifiedClassroomUser, err := s.inviteUserToOrganization(c.Context(), s.appClient, classroom, classroomRole, classroomUser.User)
			if err != nil { // we failed to invite the user, but this is not a critical failure.
				stillRequestedUsers = append(stillRequestedUsers, classroomUser)
			} else {
				invitedUsers = append(invitedUsers, modifiedClassroomUser)
			}
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":         "Invites sent successfully",
			"invited_users":   invitedUsers,
			"requested_users": stillRequestedUsers,
		})
	}
}

// Sends an invite to a user to join the organization
func (s *ClassroomService) sendOrganizationInviteToUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroomRole, err := models.NewClassroomRole(c.Params("classroom_role"))
		if err != nil {
			return errs.BadRequest(err)
		}

		// Only allow professors to invite people to org
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		inviteeUserID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		invitee, err := s.store.GetUserInClassroom(c.Context(), classroomID, inviteeUserID)
		if err != nil {
			return errs.InternalServerError()
		}

		// use the current user's client to invite the user to the organization
		invitee, err = s.inviteUserToOrganization(c.Context(), s.appClient, classroom, classroomRole, invitee.User)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Invite sent successfully",
			"user":    invitee,
		})
	}
}

// Removes a user from the requested state
func (s *ClassroomService) denyRequestedUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		// Only allow professors to remove users from classrooms
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		err = s.store.RemoveUserFromClassroom(c.Context(), classroomID, userID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.SendStatus(http.StatusOK)
	}
}

// Revokes an invite to a user to join the organization
func (s *ClassroomService) revokeOrganizationInvite() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		targetUser, err := s.store.GetUserInClassroom(c.Context(), classroomID, userID)
		if err != nil {
			return errs.InternalServerError()
		}

		// Only allow professors to remove users from classrooms
		_, err = s.RequireAtLeastRole(c, classroomID, models.Professor)
		if err != nil {
			return err
		}

		err = s.store.RemoveUserFromClassroom(c.Context(), classroomID, userID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		err = s.appClient.CancelOrgInvitation(c.Context(), classroom.OrgName, targetUser.GithubUsername)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.SendStatus(http.StatusOK)
	}
}

// Helper function to invite a user to the organization (delegates based on the role supplied)
func (s *ClassroomService) inviteUserToOrganization(ctx context.Context, client github.GitHubBaseClient, classroom models.Classroom, classroomRole models.ClassroomRole, user models.User) (models.ClassroomUser, error) {
	var classroomUser models.ClassroomUser
	var err error
	if classroomRole == models.Student {
		// Get the team ID
		studentTeam, err := client.GetTeamByName(ctx, classroom.OrgName, *classroom.StudentTeamName)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}

		// Invite the user to the organization
		classroomUser, err = s.inviteMemberToOrganization(ctx, client, *studentTeam.ID, classroom.ID, user)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
	} else {
		// Invite the user to the organization
		classroomUser, err = s.inviteAdminToOrganization(ctx, client, classroom.OrgName, classroom.ID, user)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
	}

	return classroomUser, nil
}

// Helper function to invite a student to the organization (adds them to the student team as well on acceptance)
func (s *ClassroomService) inviteMemberToOrganization(context context.Context, client github.GitHubBaseClient, teamID int64, classroomID int64, invitee models.User) (models.ClassroomUser, error) {
	err := client.AddTeamMember(context, teamID, invitee.GithubUsername, nil)
	if err != nil {
		return models.ClassroomUser{}, errs.InternalServerError()
	}
	classroomUser, err := s.store.ModifyUserStatus(context, classroomID, models.UserStatusOrgInvited, *invitee.ID)
	if err != nil {
		return models.ClassroomUser{}, errs.InternalServerError()
	}

	return classroomUser, nil
}

// Helper function to invite an admin to the organization
func (s *ClassroomService) inviteAdminToOrganization(context context.Context, client github.GitHubBaseClient, orgName string, classroomID int64, invitee models.User) (models.ClassroomUser, error) {
	err := client.SetUserMembershipInOrg(context, orgName, invitee.GithubUsername, "admin")
	if err != nil {
		return models.ClassroomUser{}, errs.InternalServerError()
	}
	classroomUser, err := s.store.ModifyUserStatus(context, classroomID, models.UserStatusOrgInvited, *invitee.ID)
	if err != nil {
		return models.ClassroomUser{}, errs.InternalServerError()
	}
	return classroomUser, nil
}

// Helper function to accept a pending invitation to an organization (Assumes there is a pending invitation)
func (s *ClassroomService) acceptOrgInvitation(context context.Context, userClient github.GitHubUserClient, orgName string, classroomID int64, invitee models.User) error {
	// user has a pending invitation, accept it
	err := userClient.AcceptOrgInvitation(context, orgName)
	if err != nil {
		return errs.InternalServerError()
	}
	_, err = s.store.ModifyUserStatus(context, classroomID, models.UserStatusActive, *invitee.ID)
	if err != nil {
		return errs.InternalServerError()
	}

	return nil
}

var semesterNameMap = map[time.Month]string{
	time.January:   "Spring",
	time.February:  "Spring",
	time.March:     "Spring",
	time.April:     "Spring",
	time.May:       "Summer 1",
	time.June:      "Summer 1",
	time.July:      "Summer 2",
	time.August:    "Summer 2",
	time.September: "Fall",
	time.October:   "Fall",
	time.November:  "Fall",
	time.December:  "Fall",
}

func (s *ClassroomService) getClassroomNames() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get the current date
		now := time.Now()
		semesterNames := []string{}
		currentMonthIndex := int(now.Month())
		for i := currentMonthIndex; i < currentMonthIndex+12; i++ {
			// generate a full year of semester names (i.e. Fall 2024, Spring 2025, etc.)
			month := time.Month((i-1)%12 + 1)
			year := now.Year()
			if i > 12 { // if the month is past December, we need to roll over to the next year
				year = year + 1
			}
			semester, ok := semesterNameMap[month]
			if !ok {
				continue // Skip if month not found in map
			}
			semesterName := fmt.Sprintf("%s %d", semester, year)
			if !stringInList(semesterNames, semesterName) {
				semesterNames = append(semesterNames, semesterName)
			}
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{"semester_names": semesterNames})
	}
}

func stringInList(list []string, str string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}
