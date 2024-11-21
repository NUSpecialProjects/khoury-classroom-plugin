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

		// Get classroom and assignment template
		classroom, err := s.store.GetClassroomByID(c.Context(), assignmentData.ClassroomID)
		if err != nil {
			return err
		}
		template, err := s.store.GetAssignmentTemplateByID(c.Context(), assignmentData.TemplateID)
		if err != nil {
			return err
		}

		// Create base repository using assignment template
		baseRepoName := generateForkName(classroom.OrgName, assignmentData.Name)
		baseRepo, err := s.appClient.CreateRepoFromTemplate(c.Context(), classroom.OrgName, template.TemplateRepoName, baseRepoName)
		if err != nil {
			fmt.Println(err)
			return errs.InternalServerError()
		}
		err = s.store.CreateBaseRepo(c.Context(), *baseRepo)
		if err != nil {
			fmt.Println("Error creating base repo")
			return err
		}

		// Store assignment in DB
		createdAssignmentID, err := s.store.CreateAssignment(c.Context(), assignmentData)
		if err != nil {
			return err
		}
		err = s.store.AttachBaseRepoToAssignment(c.Context(), createdAssignmentID, baseRepo.BaseID)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"created_assignment_id": createdAssignmentID,
		})
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
			fmt.Println("Error getting assignment by token")
			return err
		}

		// Get assignment base repository
		baseRepo, err := s.store.GetBaseRepoByID(c.Context(), assignment.BaseRepoID)
		if err != nil {
			fmt.Println("Error getting base repo by ID")
			return err
		}

		// Retrieve user client and session
		client, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}
		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// Get classroom
		classroom, err := s.store.GetClassroomByID(c.Context(), assignment.ClassroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		// Check if fork already exists
		forkName := generateForkName(assignment.Name, user.Login)
		studentWorkRepo, _ := client.GetRepository(c.Context(), classroom.OrgName, forkName)
		if studentWorkRepo != nil {
			return c.Status(http.StatusOK).JSON(fiber.Map{
				"message":  "Assignment already accepted",
				"repo_url": studentWorkRepo.HTMLURL,
			})
		}

		// Otherwise generate fork
		err = client.ForkRepository(c.Context(),
			baseRepo.BaseRepoOwner,
			classroom.OrgName,
			baseRepo.BaseRepoName,
			forkName)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// Remove student team's access to forked repo
		// TODO: dynamically find student team name (KHO-177)
		studentTeamName := "student_team_test"
		err = client.RemoveRepoFromTeam(c.Context(), classroom.OrgName, studentTeamName, classroom.OrgName, forkName)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// Insert into DB
		_, err = s.store.CreateStudentWork(c.Context(), assignment.ID, user.ID, forkName, "ACCEPTED", *assignment.MainDueDate)
		if err != nil {
			fmt.Println("Error creating student work")
			return err
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
// TODO: ensure duplicates are impossible, just append an incrementing -x to name in that case
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
