package github

import (
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (service *GitHubService) GetCurrentUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			fmt.Println("FAILED TO GET CLIENT", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			fmt.Println("FAILED TO GET USER", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch user"})
		}

		//TODO: include the user's role (i.e. professor, TA, student) in the response
		return c.Status(200).JSON(user)
	}
}
