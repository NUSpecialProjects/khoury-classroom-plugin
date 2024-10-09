package assignments

import (
	"net/http"
  "fmt"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)



func (s *AssignmentsService) GetAllAssignmentTemplates(c *fiber.Ctx) error {
	assignments, err := s.store.GetAllAssignmentTemplates(c.Context())
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(assignments)
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

func (s *AssignmentsService) CreateAssignmentTemplate(c *fiber.Ctx) error {
  var assignmentTemplateData models.AssignmentTemplate
  err := c.BodyParser(&assignmentTemplateData)
  if err != nil {
    return errs.InvalidJSON()
  }

  error := s.store.CreateAssignmentTemplate(c.Context(), assignmentTemplateData)
  if error != nil {
    return error
  }

  return c.Status(http.StatusOK).JSON(fiber.Map{
    "message": "Received assignment template data",
    "assignment_template":  assignmentTemplateData,
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
    fmt.Println(error.Error())
    return error
  }

  return c.Status(http.StatusOK).JSON(fiber.Map{
      "message": "Received assignment data",
      "assignment":  assignmentData,
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

  c.Status(http.StatusOK).JSON(fiber.Map{
    "message": "Received due date data",
    "due_date":  dueDateData,
  })
  return nil
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
    "regrade":  regradeData,
  })
}
