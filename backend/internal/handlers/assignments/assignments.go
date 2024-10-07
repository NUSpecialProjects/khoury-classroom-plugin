package assignments

import (
	"net/http"

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

  c.Status(http.StatusOK)
  return nil
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

  c.Status(http.StatusOK)
  return nil
}

func (s *AssignmentsService) CreateDueDate(c *fiber.Ctx) error {
  var dueDataData models.DueDate
  err := c.BodyParser(&dueDataData)
  if err != nil {
    return errs.InvalidJSON()
  }

  error := s.store.CreateDueDate(c.Context(), dueDataData)
  if error != nil {
    return error
  }

  c.Status(http.StatusOK)
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

	c.Status(http.StatusOK)
	return nil
}
