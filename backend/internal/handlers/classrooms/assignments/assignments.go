package assignments

import (
	"fmt"
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// GetAssignments returns the assignments in a classroom.
func (s *AssignmentService) GetAssignments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// GetAssignment returns the details of an assignment.
func (s *AssignmentService) GetAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// UpdateAssignment updates an existing assignment.
func (s *AssignmentService) UpdateAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// GenerateAssignmentToken generates a token to accept an assignment.
func (s *AssignmentService) GenerateAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// GetAssignmentSubmissions returns the submissions for an assignment.
func (s *AssignmentService) GetAssignmentSubmissions() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

func (s *AssignmentService) CreateRubric(c *fiber.Ctx) error {
	var rubricData models.Rubric

	error := c.BodyParser(&rubricData)
	if error != nil {
		return errs.InvalidJSON()
	}

	err := s.store.CreateRubric(c.Context(), rubricData)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Received rubric data",
		"rubric":  rubricData,
	})
}

// Create a new assignment in the classroom
func (s *AssignmentService) CreateAssignment(c *fiber.Ctx) error {
	var assignmentData models.Assignment
	err := c.BodyParser(&assignmentData)
	if err != nil {
		return errs.InvalidJSON()
	}

	error := s.store.CreateAssignment(c.Context(), assignmentData)
	if error != nil {
		return error
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":             "Received assignment data",
		"assignment_template": assignmentData,
	})
}

func (s *AssignmentService) CreateStudentAssignment(c *fiber.Ctx) error {
	var studentAssignmentData models.StudentAssignment
	err := c.BodyParser(&studentAssignmentData)
	if err != nil {
		return errs.InvalidJSON()
	}

	error := s.store.CreateStudentAssignment(c.Context(), studentAssignmentData)
	if error != nil {
		fmt.Println(error.Error())
		return error
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":    "Received student assignment data",
		"assignment": studentAssignmentData,
	})
}

func (s *AssignmentService) CreateDueDate(c *fiber.Ctx) error {
	var dueDateData models.DueDate
	err := c.BodyParser(&dueDateData)
	if err != nil {
		return errs.InvalidJSON()
	}

	error := s.store.CreateDueDate(c.Context(), dueDateData)
	if error != nil {
		return error
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message":  "Received due date data",
		"due_date": dueDateData,
	})
}

func (s *AssignmentService) CreateRegrade(c *fiber.Ctx) error {

	var regradeData models.Regrade
	err := c.BodyParser(&regradeData)
	if err != nil {
		return errs.InvalidJSON()
	}

	error := s.store.CreateRegrade(c.Context(), regradeData)
	if error != nil {
		return error
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Received regrade data",
		"regrade": regradeData,
	})
}
