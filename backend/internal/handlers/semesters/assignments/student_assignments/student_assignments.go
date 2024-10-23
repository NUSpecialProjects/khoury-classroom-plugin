package student_assignments

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *StudentAssignmentService) GetStudentAssignment(c *fiber.Ctx) error {
	classroomID, _ := strconv.ParseInt(c.Params("classroomID"), 10, 64)
	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)
	studentAssignmentID, _ := strconv.ParseInt(c.Params("studentAssignmentID"), 10, 64)

	studentAssignment, err := s.store.GetStudentAssignment(c.Context(), classroomID, assignmentID, studentAssignmentID)
	if err != nil {
		fmt.Println("error in service func")
		return err
	}

	return c.Status(http.StatusOK).JSON(studentAssignment)
}
