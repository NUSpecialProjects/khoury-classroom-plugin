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
			log.Default().Println("Error: bad body:", err)
			return errs.InvalidRequestBody(body)
		}

		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			log.Default().Println("Error: bad classroom_id:", err)
			return errs.BadRequest(err)
		}

		token, err := utils.GenerateToken(16)
		if err != nil {
			log.Default().Println("Error: generating token:", err)
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
			log.Default().Println("Error: creating classroom token:", err)
			return errs.NewDBError(err)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"token": classroomToken.Token})
	}
}

func (s *ClassroomService) useClassroomToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			log.Default().Println("Error: getting client:", err)
			return errs.AuthenticationError()
		}

		currentGitHubUser, err := client.GetCurrentUser(c.Context())
		if err != nil {
			log.Default().Println("Error: getting current user:", err)
			return errs.AuthenticationError()
		}

		token := c.Params("token")
		if token == "" {
			log.Default().Println("Error: missing token")
			return errs.MissingAPIParamError("token")
		}

		// Go get the token from the DB
		classroomToken, err := s.store.GetClassroomToken(c.Context(), token)
		if err != nil {
			log.Default().Println("Error: getting classroom token:", err)
			return errs.AuthenticationError()
		}

		// Check if the token is valid
		if classroomToken.ExpiresAt != nil && classroomToken.ExpiresAt.Before(time.Now()) {
			log.Default().Println("Error: classroom token expired")
			return errs.ExpiredTokenError()
		}

		// Add the user to the database if they don't exist already
		// Otherwise, get the user from the database
		user, err := s.store.GetUserByGitHubId(c.Context(), currentGitHubUser.ToUser().GithubUserID)
		if err != nil {
			log.Default().Println("Error: user doesn't already exist, creating user")
			user, err = s.store.CreateUser(c.Context(), currentGitHubUser.ToUser())
			if err != nil {
				log.Default().Println("Error: creating user:", err)
				return errs.NewDBError(err)
			}
		}

		// Check if the userWithRole is already in the classroom
		userWithRole, err := s.store.GetUserInClassroom(c.Context(), classroomToken.ClassroomID, *user.ID)
		log.Default().Printf("GOT USER IN CLASSROOM: %+v", userWithRole)
		if err == nil { // user is already in the classroom. If their role can be upgraded, do so. Don't downgrade.
			log.Default().Printf("User %s is already in classroom %d", userWithRole.GithubUsername, classroomToken.ClassroomID)
			roleComparison := models.ClassroomRole(userWithRole.Role).Compare(classroomToken.ClassroomRole)
			if roleComparison < 0 {
				// Upgrade the user's role in the classroom
				err := s.store.ModifyUserRole(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), *userWithRole.ID)
				if err != nil {
					log.Default().Println("Error: adding user to classroom:", err)
					return errs.NewAPIError(http.StatusInternalServerError, err)
				}
				log.Default().Printf("User %s role upgraded to %s", userWithRole.GithubUsername, classroomToken.ClassroomRole)
			} else if roleComparison >= 0 {
				// User's current role is higher than token role
				log.Default().Printf("User %s role is already %s or higher", userWithRole.GithubUsername, classroomToken.ClassroomRole)
				return errs.InvalidRoleOperation()
			}
		} else {
			log.Default().Printf("User %s not in classroom %d, adding them", user.GithubUsername, classroomToken.ClassroomID)
			// Add the user to the classroom
			_, err := s.store.AddUserToClassroom(c.Context(), classroomToken.ClassroomID, string(classroomToken.ClassroomRole), *user.ID)
			if err != nil {
				log.Default().Println("Error: adding user to classroom:", err)
				return errs.NewDBError(err)
			}
			log.Default().Printf("User %s added to classroom %d with role %s", user.GithubUsername, classroomToken.ClassroomID, classroomToken.ClassroomRole)
		}

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
