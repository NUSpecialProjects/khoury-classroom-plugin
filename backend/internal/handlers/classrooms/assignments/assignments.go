package assignments

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Returns the assignments in a classroom.
func (s *AssignmentService) getAssignments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		assignments, err := s.store.GetAssignmentsInClassroom(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"assignment_outlines": assignments,
		})
	}
}

// Returns the details of an assignment.
func (s *AssignmentService) getAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		assignmentID, err := strconv.ParseInt(c.Params("assignment_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		assignment, err := s.store.GetAssignmentByID(c.Context(), assignmentID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"assignment_outline": assignment,
		})
	}
}

func (s *AssignmentService) createAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse request body
		var assignmentData models.AssignmentOutline
		error := c.BodyParser(&assignmentData)
		if error != nil {
			return errs.InvalidRequestBody(assignmentData)
		}

		// Store assignment in DB
		createdAssignment, err := s.store.CreateAssignment(c.Context(), assignmentData)
		if err != nil {
			return err
		}

		// Get classroom and assignment template
		classroom, err := s.store.GetClassroomByID(c.Context(), createdAssignment.ClassroomID)
		if err != nil {
			return err
		}
		template, err := s.store.GetAssignmentTemplateByID(c.Context(), createdAssignment.TemplateID)
		if err != nil {
			return err
		}

		// Create base repository using assignment template
		baseRepoName := generateForkName(classroom.OrgName, assignmentData.Name)
		err = s.appClient.CreateBaseAssignmentRepo(c.Context(), classroom.OrgName, template.TemplateRepoName, baseRepoName)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"created_assignment": createdAssignment,
		})
	}
}

// TODO: Choose naming pattern once we have a full assignment flow. Stub for now
func generateForkName(sourceName, userName string) string {
	return sourceName + "-" + strings.ReplaceAll(userName, " ", "")
}

// Updates an existing assignment.
func (s *AssignmentService) updateAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Generates a token to accept an assignment.
func (s *AssignmentService) generateAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

// Uses a token to accept an assignment.
func (s *AssignmentService) useAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
