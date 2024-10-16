package student_assignments

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *StudentAssignmentService) GetStudentAssignmentFiles(c *fiber.Ctx) error {
	assignment, err := s.githubappclient.GetStudentAssignmentFiles(c.Params("owner"), c.Params("repo"), c.Params("*"))
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(assignment)
}
