package student_assignments

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *StudentAssignmentService) GetStudentAssignmentsByLocalID(c *fiber.Ctx) error {
	classroomID, _ := strconv.ParseInt(c.Params("classroomID"), 10, 64)
	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)

	studentAssignment, err := s.store.GetStudentAssignmentsByLocalID(c.Context(), classroomID, assignmentID)
	if err != nil {
		fmt.Println("error in service func")
		return err
	}

	return c.Status(http.StatusOK).JSON(studentAssignment)
}

func (s *StudentAssignmentService) GetStudentAssignmentByLocalID(c *fiber.Ctx) error {
	classroomID, _ := strconv.ParseInt(c.Params("classroomID"), 10, 64)
	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)
	studentAssignmentID, _ := strconv.ParseInt(c.Params("studentAssignmentID"), 10, 64)

	studentAssignment, err := s.store.GetStudentAssignmentByLocalID(c.Context(), classroomID, assignmentID, studentAssignmentID)
	if err != nil {
		fmt.Println("error in service func")
		return err
	}

	return c.Status(http.StatusOK).JSON(studentAssignment)
}


func (s *StudentAssignmentService) GetStudentAssignmentsByAssignmentID(c *fiber.Ctx) error {
	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)

    studentAssignments, err := s.store.GetStudentAssignmentsByAssignmentID(c.Context(), assignmentID)
    if err != nil {
        return err
    }

    fmt.Println(studentAssignments[0].StudentGHUsernames)

    return c.Status(http.StatusOK).JSON(studentAssignments)
}
