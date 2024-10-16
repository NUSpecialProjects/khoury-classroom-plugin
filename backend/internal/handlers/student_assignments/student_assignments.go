package student_assignments

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *StudentAssignmentService) GetRepoContents(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON("test")
	/*assignment, err := s.githubappclient.GetStudentAssignment(c.Context(), c.Params("owner"), c.Params("repo"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(assignment)*/
}
