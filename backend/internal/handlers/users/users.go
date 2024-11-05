package users

import (
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Creates a user representing a TA, Student, or Professor.
func (s *UserService) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userToCreate models.User
		err := c.BodyParser(&userToCreate)
		if err != nil {
			return errs.InvalidRequestBody(models.User{})
		}

		createdUser, err := s.store.CreateUser(c.Context(), userToCreate)
		if err != nil {
			return err
		}

		// Implement logic here
		return c.Status(http.StatusOK).JSON(createdUser)
	}
}
