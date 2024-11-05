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

		return c.Status(http.StatusOK).JSON(fiber.Map{"token": classroomToken})
	}
}

func (s *ClassroomService) useClassroomToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return err
		}

		currentUser, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return err
		}

		token := c.Params("token")
		if token == "" {
			return errs.MissingAPIParamError("token")
		}

		// Go get the token from the DB
		classroomToken, err := s.store.GetClassroomToken(c.Context(), token)
		if err != nil {
			return err
		}

		// Check if the token is valid
		if classroomToken.ExpiresAt != nil && classroomToken.ExpiresAt.Before(time.Now()) {
			return errs.AuthenticationError()
		}

		// Add the user to the database if they don't exist already
		_, err = s.store.GetUserByGitHubId(c.Context(), currentUser.ToUser().GithubUserID)
		if err == nil {
			s.store.CreateUser(c.Context(), currentUser.ToUser())
		}

		// Check if the userWithRole is already in the classroom
		userWithRole, err := s.store.GetUserInClassroom(c.Context(), classroomToken.ClassroomID, currentUser.ID)
		if err == nil { // user is already in the classroom. If their role can be upgraded, do so. Don't downgrade.
			if models.ClassroomRole(userWithRole.Role).Compare(classroomToken.ClassroomRole) < 0 {
				// Upgrade the user's role in the classroom
				userID, err := s.store.AddUserToClassroom(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), currentUser.ID)
				if err != nil {
					return errs.NewDBError(err)
				}
				// Otherwise do nothing to their current role
				log.Default().Println("User ", userID, " role upgraded to ", classroomToken.ClassroomRole)
			}
		} else {
			// Add the user to the classroom
			_, err := s.store.AddUserToClassroom(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), currentUser.ID)
			if err != nil {
				return errs.NewDBError(err)
			}
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Token applied successfully"})
	}
}
