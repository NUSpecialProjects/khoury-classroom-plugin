package classrooms

import (
	"github.com/gofiber/fiber/v2"
)

// GetSubmissions returns the submissions for an assignment.
func (s *ClassroomService) GetSubmissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// GetSubmission returns the details of a submission.
func (s *ClassroomService) GetSubmission() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
