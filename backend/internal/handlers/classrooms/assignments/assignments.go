package assignments

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/CamPlume1/khoury-classroom/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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

		// Error if assignment already exists
		existingAssignment, err := s.store.GetAssignmentByNameAndClassroomID(c.Context(), assignmentData.Name, assignmentData.ClassroomID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return err
		}
		if existingAssignment != nil {
			return errs.BadRequest(errors.New("assignment with that name already exists"))
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

		// Create base repository and store locally
		baseRepoName := generateForkName(classroom.OrgName, assignmentData.Name)
		baseRepo, err := s.appClient.CreateRepoFromTemplate(c.Context(), classroom.OrgName, template.TemplateRepoName, baseRepoName)
		if err != nil {
			return err
		}
		err = s.store.CreateBaseRepo(c.Context(), *baseRepo)
		if err != nil {
			return err
		}

		// Store assignment locally
		assignmentData.BaseRepoID = baseRepo.BaseID
		createdAssignment, err := s.store.CreateAssignment(c.Context(), assignmentData)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"created_assignment": createdAssignment,
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
			return err
		}

		// Get assignment base repository
		baseRepo, err := s.store.GetBaseRepoByID(c.Context(), assignment.BaseRepoID)
		if err != nil {
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
			// Ensure student team is removed
			err = client.RemoveRepoFromTeam(c.Context(), classroom.OrgName, *classroom.StudentTeamName, classroom.OrgName, forkName)
			if err != nil {
				return errs.GithubAPIError(err)
			}

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

		// Wait to perform actions on the fork until it is finished initializing
		initialDelay := 1 * time.Second
		maxDelay := 30 * time.Second
		for {
			studentWorkRepo, _ = client.GetRepository(c.Context(), classroom.OrgName, forkName)
			if studentWorkRepo != nil {
				break
			}

			if initialDelay > maxDelay {
				return errs.GithubAPIError(errors.New("fork unsuccessful, please try again later"))
			}

			time.Sleep(initialDelay)
			initialDelay *= 2
		}

		// Remove student team's access to forked repo
		err = client.RemoveRepoFromTeam(c.Context(), classroom.OrgName, *classroom.StudentTeamName, classroom.OrgName, *studentWorkRepo.Name)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// Insert into DB
		_, err = s.store.CreateStudentWork(c.Context(), assignment.ID, user.ID, forkName, models.WorkStateAccepted, *assignment.MainDueDate)
		if err != nil {
			return err
		}

		// Instead of getting the repository immediately, construct the expected URL
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message":  "Assignment Accepted!",
			"repo_url": studentWorkRepo.HTMLURL,
		})
	}
}

func (s *AssignmentService) checkAssignmentName() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Fetch assignment name and classrooID from request
		assignmentName := c.Params("assignment_name")
		if assignmentName == "" {
			return errs.BadRequest(errors.New("assignment name is required"))
		}
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		// Check if assignment with name exists
		assignment, err := s.store.GetAssignmentByNameAndClassroomID(c.Context(), assignmentName, classroomID)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"exists": assignment != nil,
		})
	}
}

// KHO-209
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
