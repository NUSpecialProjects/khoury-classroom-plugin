package assignments

import (
	"github.com/gofiber/fiber/v2"
)

// Returns the assignments in a classroom.
func (s *AssignmentService) getAssignments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Returns the details of an assignment.
func (s *AssignmentService) getAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

func (s *AssignmentService) createAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Updates an existing assignment.
func (s *AssignmentService) updateAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Generates a token to accept an assignment.
func (s *AssignmentService) generateAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Uses a token to accept an assignment.
func (s *AssignmentService) useAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
