package assignments

import (
	"fmt"
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
func (s *AssignmentService) GetAssignments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		assignments, err := s.store.GetAssignmentsInClassroom(c.Context(), classroomID)
		if err != nil {
			fmt.Println("This one: " + err.Error())
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
		var assignmentData models.AssignmentOutline
		err := c.BodyParser(&assignmentData)
		
		if err != nil  {
			return errs.InvalidRequestBody(models.AssignmentOutline{})
		}

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
		fmt.Printf("Reached handler")
		var assignment models.AssignmentAcceptRequest
		err := c.BodyParser(&assignment)
		if err != nil {
			return errs.InvalidRequestBody(models.AssignmentOutline{})
		}

		client, err := middleware.GetClient(c, s.store, s.userCfg)

		if err != nil {
			return errs.AuthenticationError()
		}

		username, err := client.GetCurrentUser(c.Context())
		if err != nil {
			//TODO: Rebase on sebyBranch for correct errors
			return err
		}

		forkName := assignment.SourceRepoName + "-" + strings.ReplaceAll((username.Login), " ", "")

		studentwork := createMockStudentWork(forkName, assignment.AssignmentName, int(assignment.AssignmentID))


		err = s.store.CreateStudentWork(c.Context(), &studentwork, username.ID)
		if err != nil {
			fmt.Printf(err.Error())
			return errs.InternalServerError()
		}

		err = client.ForkRepository(c.Context(), assignment.OrgName, assignment.OrgName, assignment.SourceRepoName, forkName)
		if err != nil {
			errs.InternalServerError()
		}

		c.Status(http.StatusOK)
		return nil
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





func createMockStudentWork(repo string, assName string, assID int) models.StudentWork {
	assignmentName := assName
	repoName := repo
	assignmentID := assID
	fmt.Println(assignmentID)
	submittedPRNumber := 42
	manualFeedbackScore := 85
	autoGraderScore := 90
	uniqueDueDate := time.Now().AddDate(0, 0, 7)            // Due in 7 days MOCK
	submissionTimestamp := time.Now().AddDate(0, 0, -1)     // Submitted yesterday MOCK
	gradesPublishedTimestamp := time.Now().AddDate(0, 0, -1) // Grades published yesterday MOCK
	return models.StudentWork{
		ID:                       1,
		ClassroomID:              101,
		AssignmentName:           &assignmentName,
		AssignmentOutlineID:      1,
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