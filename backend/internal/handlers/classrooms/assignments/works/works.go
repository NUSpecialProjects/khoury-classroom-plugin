package works

import (
	"github.com/gofiber/fiber/v2"
)

// Returns the student works for an assignment.
func (s *WorkService) getWorks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Returns the details of a specific student work.
func (s *WorkService) getWork() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
