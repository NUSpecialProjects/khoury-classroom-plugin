package users

import (
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Returns the classrooms the authenticated user is part of.
func (s *UserService) CreateUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userToCreate models.User
        err := c.BodyParser(&userToCreate)
        if err != nil {
            return err
        }

        createdUser, err := s.store.CreateUser(c.Context(), userToCreate)
        if err != nil {
            return err
        }

        // Implement logic here
        return c.Status(http.StatusOK).JSON(fiber.Map{
            "message": "Created User",
            "created_user" : createdUser,
        })
	}
}


