package assignments

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (s *AssignmentsService) GetAssignmentsInSemester(c *fiber.Ctx) error {
	classroomID, err := strconv.ParseInt(c.Params("classroomID"), 10, 64)
	if err != nil {
		return err
	}

	assignments, err := s.store.GetAssignmentsInSemester(c.Context(), classroomID)
	if err != nil {
		fmt.Println("error in service func", err)
		return err
	}

	return c.Status(http.StatusOK).JSON(assignments)
}

func (s *AssignmentsService) GetAssignment(c *fiber.Ctx) error {
	classroomID, _ := strconv.ParseInt(c.Params("classroomID"), 10, 64)
	assignmentID, _ := strconv.ParseInt(c.Params("assignmentID"), 10, 64)

	assignment, err := s.store.GetAssignment(c.Context(), classroomID, assignmentID)
	if err != nil {
		fmt.Println("error in service func", err)
		return err
	}

	return c.Status(http.StatusOK).JSON(assignment)
}

func (s *AssignmentsService) CreateRubric(c *fiber.Ctx) error {
	var rubricData models.Rubric

	error := c.BodyParser(&rubricData)
	if error != nil {
		return errs.InvalidJSON()
	}

	err := s.store.CreateRubric(c.Context(), rubricData)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Received rubric data",
		"rubric":  rubricData,
	})
}

func (s *AssignmentsService) CreateAssignment(c *fiber.Ctx) error {
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

func (s *AssignmentsService) CreateStudentAssignment(c *fiber.Ctx) error {
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

func (s *AssignmentsService) CreateDueDate(c *fiber.Ctx) error {
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

func (s *AssignmentsService) CreateRegrade(c *fiber.Ctx) error {

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
