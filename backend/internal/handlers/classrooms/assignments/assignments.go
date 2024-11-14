package assignments

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
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
		var assignmentData models.AssignmentOutlineRequest
		error := c.BodyParser(&assignmentData)
		if error != nil {
			return errs.InvalidRequestBody(assignmentData)
		}

		// Store assignment in DB
		createdAssignment, err := s.store.CreateAssignment(c.Context(), assignmentData)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"created_assignment": createdAssignment,
		})
	}
}

func (s *AssignmentService) acceptAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check + parse FE request
		var assignment models.AssignmentAcceptRequest
		err := c.BodyParser(&assignment)
		if err != nil {
			return errs.InvalidRequestBody(models.AssignmentOutline{})
		}

		// Retrieve user client
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		// Retrieve current session
		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// Insert into DB
		forkName := generateForkName(assignment.SourceRepoName, user.Login)
		studentwork := createMockStudentWork(forkName, assignment.AssignmentName, int(assignment.AssignmentID))
		err = s.store.CreateStudentWork(c.Context(), &studentwork, user.ID)
		if err != nil {
			return err
		}

		// Generate Fork via GH User
		err = client.ForkRepository(c.Context(), assignment.OrgName, assignment.OrgName, assignment.SourceRepoName, forkName)
		if err != nil {
			return err
		}

		c.Status(http.StatusOK)
		return nil
	}
}

// TODO: Choose naming pattern once we have a full assignment flow. Stub for now
func generateForkName(sourceName, userName string) string {
	return sourceName + "-" + strings.ReplaceAll(userName, " ", "")
}

// TODO: Integrate with actual assignment information when infrastructure is available
func createMockStudentWork(repo string, assName string, assID int) models.StudentWork {
	assignmentName := assName
	repoName := repo
	submittedPRNumber := 42
	manualFeedbackScore := 85
	autoGraderScore := 90
	uniqueDueDate := time.Now().AddDate(0, 0, 7)             // Due in 7 days MOCK
	submissionTimestamp := time.Now().AddDate(0, 0, -1)      // Submitted yesterday MOCK
	gradesPublishedTimestamp := time.Now().AddDate(0, 0, -1) // Grades published yesterday MOCK
	return models.StudentWork{
		ID:                       1,
		ClassroomID:              101,
		AssignmentName:           &assignmentName,
		AssignmentOutlineID:      assID,
		RepoName:                 &repoName,
		UniqueDueDate:            &uniqueDueDate,
		SubmittedPRNumber:        &submittedPRNumber,
		ManualFeedbackScore:      &manualFeedbackScore,
		AutoGraderScore:          &autoGraderScore,
		SubmissionTimestamp:      &submissionTimestamp,
		GradesPublishedTimestamp: &gradesPublishedTimestamp,
		WorkState:                "SUBMITTED",
		CreatedAt:                time.Now(),
	}
}

func (s *AssignmentService) createAssignmentTemplate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var assignmentData models.AssignmentTemplate

		error := c.BodyParser(&assignmentData)
		if error != nil {
			return errs.InvalidRequestBody(assignmentData)
		}

		// Check if the template already exists
		exists, err := s.store.AssignmentTemplateExists(c.Context(), assignmentData.TemplateID)
		if err != nil {
			return errs.InternalServerError()
		}
		if exists {
			return c.Status(http.StatusOK).JSON("Template already exists")
		}

		// Create the template if it does not exist
		createdTemplate, err := s.store.CreateAssignmentTemplate(c.Context(), assignmentData)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"created_template": createdTemplate,
		})
	}
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
