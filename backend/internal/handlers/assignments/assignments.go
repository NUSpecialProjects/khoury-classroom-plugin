package assignments

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)



func (s *AssignmentsService) GetAllAssignmentTemplates(c *fiber.Ctx) error {
	assignments, err := s.store.GetAllAssignmentTemplates(c.Context())
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(assignments)
}


