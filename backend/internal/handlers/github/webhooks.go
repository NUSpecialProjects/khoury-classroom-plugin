package github

import (
	"github.com/gofiber/fiber/v2"
)

func (s *GitHubService) WebhookHandler(c *fiber.Ctx) error {
	var dispatch = map[string]func(c *fiber.Ctx) error{
		"push":   s.Push,
		"create": s.Create,
	}
	event := c.Get("X-GitHub-Event", "")

	handler, exists := dispatch[event]
	if !exists {
		return c.SendStatus(400)
	}

	return handler(c)
}

func (s *GitHubService) Push(c *fiber.Ctx) error {
	println("push webhook event")
	return c.SendStatus(200)
}

func (s *GitHubService) Create(c *fiber.Ctx) error {
	println("create webhook event")
	return c.SendStatus(200)
}
