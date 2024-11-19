package classrooms

import (
	"context"
	"net/http"
	"strconv"
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
		var classroomData models.Classroom
		err := c.BodyParser(&classroomData)
		if err != nil {
			return errs.InvalidRequestBody(models.Classroom{})
		}

		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		githubUser, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return errs.AuthenticationError()
		}

		user, err := s.store.GetUserByGitHubID(c.Context(), githubUser.ID)
		if err != nil {
			return errs.AuthenticationError()
		}

		membership, err := client.GetUserOrgMembership(c.Context(), classroomData.OrgName, githubUser.Login)
		if err != nil {
			return errs.InconsistentOrgMembershipError()
		}

		if *membership.Role != "admin" {
			return errs.InconsistentOrgMembershipError()
		}

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

		updatedClassroom, err := s.store.UpdateClassroom(c.Context(), classroomData)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"classroom": updatedClassroom})
	}
}

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
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
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
			if err != nil && newClassroomUser.Status != models.UserStatusNotInOrg { // failed to update their status, so just keep the old one
				updatedUsersInClassroom = append(updatedUsersInClassroom, classroomUser)
			} else if newClassroomUser.Status != models.UserStatusNotInOrg { // add the updated user to the list
				updatedUsersInClassroom = append(updatedUsersInClassroom, newClassroomUser)
			}

		}
		return c.Status(http.StatusOK).JSON(fiber.Map{"users": updatedUsersInClassroom})
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

		token, err := utils.GenerateToken(16)
		if err != nil {
			return errs.InternalServerError()
		}

		classroomRole, err := models.NewClassroomRole(body.ClassroomRole)
		if err != nil {
			return errs.BadRequest(err)
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
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		currentGitHubUser, err := client.GetCurrentUser(c.Context())
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

		user, err := s.store.GetUserByGitHubID(c.Context(), currentGitHubUser.ToUser().GithubUserID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroomUser, err := s.store.GetUserInClassroom(c.Context(), classroomToken.ClassroomID, *user.ID)

		// If the user is already in the classroom, we need to check if their role can be upgraded
		if err == nil { // user is already in the classroom. If their role can be upgraded, do so. Don't downgrade them.
			roleComparison := classroomUser.Role.Compare(classroomToken.ClassroomRole)
			if roleComparison < 0 {
				// Upgrade the user's role in the classroom
				classroomUser, err = s.store.ModifyUserRole(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), *classroomUser.ID)
				if err != nil {
					return errs.InternalServerError()
				}
			} else if roleComparison >= 0 {
				// User's current role is higher than token role, therefore do nothing and return an error
				return errs.InvalidRoleOperation()
			}
		} else if classroomUser.Status == models.UserStatusNotInOrg { // user previously denied their invite, but has clicked their link, so modify their role
			classroomUser, err = s.store.ModifyUserRole(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), *classroomUser.ID)
			if err != nil {
				return errs.InternalServerError()
			}
		} else { // user is not in the classroom, add them with the token's role
			classroomUser, err = s.store.AddUserToClassroom(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), models.UserStatusRequested, *user.ID)
			if err != nil {
				return errs.InternalServerError()
			}
		}

		// Invite the user to the organization
		classroomUser, err = s.inviteUserToOrganization(c.Context(), s.githubappclient, classroom.OrgName, classroomToken.ClassroomID, classroomToken.ClassroomRole, user)
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

func (s *ClassroomService) getCurrentClassroomUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		githubUser, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return errs.AuthenticationError()
		}

		user, err := s.store.GetUserByGitHubID(c.Context(), githubUser.ID)
		if err != nil {
			return errs.AuthenticationError()
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

		if classroomUser.Status == models.UserStatusNotInOrg {
			return errs.InconsistentOrgMembershipError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"user": classroomUser})
	}
}

// Updates the user's status in our DB to reflect their org membership, as of this moment
func (s *ClassroomService) updateUserStatus(ctx context.Context, client github.GitHubUserClient, user models.User, classroom models.Classroom) (models.ClassroomUser, error) {
	//TODO: check if user has declined the invite
	classroomUser, err := s.store.GetUserInClassroom(ctx, classroom.ID, *user.ID)
	if err != nil {
		return models.ClassroomUser{}, errs.UserNotFoundInClassroomError()
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
			modifiedClassroomUser, err := s.inviteUserToOrganization(c.Context(), s.githubappclient, classroom.OrgName, classroomID, classroomRole, classroomUser.User)
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

		userID, err := strconv.ParseInt(c.Params("user_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		invitee, err := s.store.GetUserInClassroom(c.Context(), classroomID, userID)
		if err != nil {
			return errs.InternalServerError()
		}

		// use the current user's client to invite the user to the organization
		invitee, err = s.inviteUserToOrganization(c.Context(), s.githubappclient, classroom.OrgName, classroomID, classroomRole, invitee.User)
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

		err = s.store.RemoveUserFromClassroom(c.Context(), classroomID, userID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroomUser, err := s.store.GetUserInClassroom(c.Context(), classroomID, userID)
		if err != nil {
			return errs.InternalServerError()
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		err = s.githubappclient.CancelOrgInvitation(c.Context(), classroom.OrgName, classroomUser.GithubUsername)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.SendStatus(http.StatusOK)
	}
}

// Invites a user to the organization
func (s *ClassroomService) inviteUserToOrganization(context context.Context, client github.GitHubBaseClient, orgName string, classroomID int64, classroomRole models.ClassroomRole, invitee models.User) (models.ClassroomUser, error) {
	var err error
	var classroomUser models.ClassroomUser
	if classroomRole == "STUDENT" {
		err = client.SetUserMembershipInOrg(context, orgName, invitee.GithubUsername, "member")
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
		classroomUser, err = s.store.ModifyUserStatus(context, classroomID, models.UserStatusOrgInvited, *invitee.ID)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
	} else {
		err = client.SetUserMembershipInOrg(context, orgName, invitee.GithubUsername, "admin")
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
		classroomUser, err = s.store.ModifyUserStatus(context, classroomID, models.UserStatusOrgInvited, *invitee.ID)
		if err != nil {
			return models.ClassroomUser{}, errs.InternalServerError()
		}
	}
	return classroomUser, nil
}

// Accepts a pending invitation to an organization (Assumes there is a pending invitation)
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
