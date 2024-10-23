package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (service *GitHubService) GetUserRoles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		var semester models.Semester
		if err := c.BodyParser(&semester); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		roles, err := client.GetUserRoles(c.Context(), semester)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch roles"})
		}

		return c.Status(200).JSON(roles)
	}
}
