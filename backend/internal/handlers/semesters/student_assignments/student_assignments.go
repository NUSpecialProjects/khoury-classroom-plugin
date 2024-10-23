package student_assignments

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *StudentAssignmentService) GetStudentAssignment(c *fiber.Ctx) error {
	// get the requested student assignment
	studentAssignment, err := s.store.GetStudentAssignment(
		c.Context(),
		c.Params("semesterID"),
		c.Params("assignmentID"),
		c.Params("studentAssignmentID"))
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(studentAssignment)
}
