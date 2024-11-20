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
			return errs.AuthenticationError()
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
		body := models.AssignmentTokenRequestBody{}

		if err := c.BodyParser(&body); err != nil {
			return errs.InvalidRequestBody(body)
		}

		assignmentID, err := strconv.ParseInt(c.Params("assignment_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		token, err := utils.GenerateToken(16)
		if err != nil {
			return errs.InternalServerError()
		}

		tokenData := models.AssignmentToken{
			AssignmentID: assignmentID,
			BaseToken: models.BaseToken{
				Token: token,
			},
		}

		// Set ExpiresAt only if Duration is provided
		if body.Duration != nil {
			expiresAt := time.Now().Add(time.Duration(*body.Duration) * time.Minute)
			tokenData.ExpiresAt = &expiresAt
		}

		assignmentToken, err := s.store.CreateAssignmentToken(c.Context(), tokenData)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"token": assignmentToken.Token})
	}
}

// Uses an assignment token to accept an assignment.
func (s *AssignmentService) useAssignmentToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Params("token")
		if token == "" {
			return errs.BadRequest(errors.New("token is required"))
		}

		// Get assignment using the token
		assignment, err := s.store.GetAssignmentByToken(c.Context(), token)
		if err != nil {
			return errs.InternalServerError()
		}

		// Get assignment with additional template information
		assignmentWithTemplate, err := s.store.GetAssignmentWithTemplateByAssignmentID(c.Context(), int64(assignment.ID))
		if err != nil {
			return errs.InternalServerError()
		}

		//Retrieve user client
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		// Retrieve current session
		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// Get classroom
		classroom, err := s.store.GetClassroomByID(c.Context(), assignment.ClassroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		templateRepoName := assignmentWithTemplate.Template.TemplateRepoName
		templateRepoOwner := assignmentWithTemplate.Template.TemplateRepoOwner

		// Generate fork name
		forkName := generateForkName(templateRepoName, user.Login)

		// Check if fork already exists
		studentWorkRepo, _ := client.GetRepository(c.Context(), classroom.OrgName, forkName)
		if studentWorkRepo != nil { // Fork already exists, early return
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message":  "Assignment already accepted",
				"repo_url": studentWorkRepo.HTMLURL,
			})
		}

		// Insert into DB
		studentwork := createMockStudentWork(forkName, assignment.Name, int(assignment.ID))
		err = s.store.CreateStudentWork(c.Context(), &studentwork, user.ID)
		if err != nil {
			return errs.InternalServerError()
		}

		// Generate Fork
		err = client.ForkRepository(c.Context(), templateRepoOwner, classroom.OrgName, templateRepoName, forkName)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// Instead of getting the repository immediately, construct the expected URL
		expectedRepoURL := fmt.Sprintf("https://github.com/%s/%s", classroom.OrgName, forkName)

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":  "Assignment accepted - it may take a few minutes to create the repository",
			"repo_url": expectedRepoURL,
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
