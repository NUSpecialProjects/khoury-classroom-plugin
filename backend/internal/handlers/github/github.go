package github

import "github.com/gofiber/fiber/v2"

func (service *GitHubService) Ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"message": "pong"})
	}
}

func (service *GitHubService) GitHubPing() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := service.githubappclient.Ping(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to ping GitHub"})
		}

		return c.Status(200).JSON(fiber.Map{"message": res})
	}
}
