package hello

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Service) HelloWorld(c *fiber.Ctx) error {
	s.github.Ping()
	return c.SendString("Hello, World!")
}
