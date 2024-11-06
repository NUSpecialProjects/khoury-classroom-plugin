package classrooms

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
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
			return err
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

		createdClassroom, err := s.store.CreateClassroom(c.Context(), classroomData)
		if err != nil {
			return err
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
			return err
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
			return err
		}
		existingClassroom.Name = classroomData.Name

		updatedClassroom, err := s.store.UpdateClassroom(c.Context(), existingClassroom)
		if err != nil {
			return err
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

		usersInClassroom, err := s.store.GetUsersInClassroom(c.Context(), classroomID)
		if err != nil {
			return err
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{"users": usersInClassroom})
	}
}

// Removes a user from a classroom.
func (s *ClassroomService) removeUserFromClassroom() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Generates a token to join a classroom.
func (s *ClassroomService) generateClassroomToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := struct {
			ClassroomRole string `json:"classroom_role"`
			Duration      *int   `json:"duration,omitempty"` // Duration is optional
		}{}

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

		tokenData := models.ClassroomToken{
			ClassroomID:   classroomID,
			ClassroomRole: models.ClassroomRole(body.ClassroomRole),
			BaseToken: models.BaseToken{
				Token:     token,
				CreatedAt: time.Now(),
			},
		}

		// Set ExpiresAt only if Duration is provided
		if body.Duration != nil {
			expiresAt := time.Now().Add(time.Duration(*body.Duration) * time.Minute)
			tokenData.ExpiresAt = &expiresAt
		}

		classroomToken, err := s.store.CreateClassroomToken(c.Context(), tokenData)
		if err != nil {
			return errs.NewDBError(err)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"token": classroomToken.Token})
	}
}

// Uses a classroom token to join a classroom.
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

		// Add the user to the database if they don't exist already
		// Otherwise, get the user from the database
		user, err := s.store.GetUserByGitHubID(c.Context(), currentGitHubUser.ToUser().GithubUserID)
		if err != nil {
			user, err = s.store.CreateUser(c.Context(), currentGitHubUser.ToUser())
			if err != nil {
				return errs.NewDBError(err)
			}
		}

		// Check if the userWithRole is already in the classroom
		userWithRole, err := s.store.GetUserInClassroom(c.Context(), classroomToken.ClassroomID, *user.ID)
		if err == nil { // user is already in the classroom. If their role can be upgraded, do so. Don't downgrade.
			roleComparison := models.ClassroomRole(userWithRole.Role).Compare(classroomToken.ClassroomRole)
			if roleComparison < 0 {
				// Upgrade the user's role in the classroom
				err := s.store.ModifyUserRole(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), *userWithRole.ID)
				if err != nil {
					return errs.NewAPIError(http.StatusInternalServerError, err)
				}
			} else if roleComparison >= 0 {
				// User's current role is higher than token role, therefore do nothing and return an error
				return errs.InvalidRoleOperation()
			}
		} else { // user is not in the classroom, add them with the token's role
			_, err := s.store.AddUserToClassroom(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), *user.ID)
			if err != nil {
				return errs.NewDBError(err)
			}
		}

		// Return the classroom the user was added to
		classroom, err := s.store.GetClassroomByID(c.Context(), classroomToken.ClassroomID)
		if err != nil {
			return errs.NewDBError(err)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":   "Token applied successfully",
			"classroom": classroom,
		})
	}
}

func (service *ClassroomService) GetCurrentClassroomUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			log.Default().Println("Error getting client:", err)
			return errs.AuthenticationError()
		}

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			log.Default().Println("Error parsing classroom ID:", err)
			return errs.BadRequest(err)
		}

		githubUser, err := client.GetCurrentUser(c.Context())
		if err != nil {
			log.Default().Println("Error getting GitHub user:", err)
			return errs.AuthenticationError()
		}

		user, err := service.store.GetUserByGitHubID(c.Context(), githubUser.ToUser().GithubUserID)
		if err != nil {
			log.Default().Println("Error getting user by GitHub ID:", err)
			return errs.AuthenticationError()
		}

		userWithRole, err := service.store.GetUserInClassroom(c.Context(), classroomID, *user.ID)
		if err != nil {
			log.Default().Println("Error getting user in classroom:", err)
			return errs.AuthenticationError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"user": userWithRole})
	}
}
