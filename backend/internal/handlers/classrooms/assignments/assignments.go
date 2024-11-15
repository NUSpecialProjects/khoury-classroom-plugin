package assignments

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/CamPlume1/khoury-classroom/internal/utils"
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

// Generates a token to accept an assignment.
func (s *AssignmentService) generateAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Generating assignment token")
		body := models.AssignmentTokenRequestBody{}

		if err := c.BodyParser(&body); err != nil {
			fmt.Println("Error parsing body", err)
			return errs.InvalidRequestBody(body)
		}

		assignmentID, err := strconv.ParseInt(c.Params("assignment_id"), 10, 64)
		if err != nil {
			fmt.Println("Error parsing assignment ID", err)
			return errs.BadRequest(err)
		}

		token, err := utils.GenerateToken(16)
		if err != nil {
			fmt.Println("Error generating token", err)
			return errs.InternalServerError()
		}

		tokenData := models.AssignmentToken{
			AssignmentID: assignmentID,
			BaseToken: models.BaseToken{
				Token:     token,
				CreatedAt: time.Now(),
			},
		}

		// Set ExpiresAt only if Duration is provided
		if body.Duration != nil {
			expiresAt := time.Now().Add(time.Duration(*body.Duration) * time.Minute)
			tokenData.ExpiresAt = &expiresAt
		}

		assignmentToken, err := s.store.CreateAssignmentToken(c.Context(), tokenData)
		if err != nil {
			fmt.Println("Error creating assignment token", err)
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"token": assignmentToken.Token})
	}
}

// Uses an assignment token to accept an assignment.
func (s *AssignmentService) useAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("Using assignment token")
		token := c.Params("token")
		if token == "" {
			fmt.Println("Token is required")
			return errs.BadRequest(errors.New("token is required"))
		}

		assignment, err := s.store.GetAssignmentByToken(c.Context(), token)
		if err != nil {
			fmt.Println("Error getting assignment by token", err)
			return errs.InternalServerError()
		}

		assignmentWithTemplate, err := s.store.GetAssignmentWithTemplateByAssignmentID(c.Context(), int64(assignment.ID))
		if err != nil {
			fmt.Println("Error getting assignment with template by assignment ID", err)
			return errs.InternalServerError()
		}

		//Retrieve user client
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			fmt.Println("Error getting client", err)
			return errs.AuthenticationError()
		}

		// Retrieve current session
		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			fmt.Println("Error getting current user", err)
			return errs.GithubAPIError(err)
		}

		classroom, err := s.store.GetClassroomByID(c.Context(), assignment.ClassroomID)
		if err != nil {
			fmt.Println("Error getting classroom by ID", err)
			return errs.InternalServerError()
		}

		templateRepoName := assignmentWithTemplate.Template.TemplateRepoName

		//Insert into DB
		forkName := generateForkName(templateRepoName, user.Login)
		studentwork := createMockStudentWork(forkName, assignment.Name, int(assignment.ID))
		err = s.store.CreateStudentWork(c.Context(), &studentwork, user.ID)
		if err != nil {
			fmt.Println("Error creating student work", err)
			return errs.InternalServerError()
		}

		// Generate Fork via GH User
		err = client.ForkRepository(c.Context(), classroom.OrgName, classroom.OrgName, templateRepoName, forkName)
		if err != nil {
			fmt.Println("Error forking repository", err)
			return errs.InternalServerError()
		}

		fmt.Println("Assignment accepted")

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":    "Assignment accepted",
			"assignment": assignment,
		})
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

// Updates an existing assignment.
func (s *AssignmentService) updateAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}
