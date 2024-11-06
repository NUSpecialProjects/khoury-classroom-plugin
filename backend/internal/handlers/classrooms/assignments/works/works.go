package works

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Returns the student works for an assignment.
func (s *WorkService) getWorks() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.Atoi(c.Params("classroom_id"))
		if err != nil {
			return errs.BadRequest(err)
		}
		assignmentID, err := strconv.Atoi(c.Params("assignment_id"))
		if err != nil {
			return errs.BadRequest(err)
		}

		works, err := s.store.GetWorks(c.Context(), classroomID, assignmentID)
		if err != nil {
			return c.SendStatus(404)
		}
		return c.Status(200).JSON(fiber.Map{
			"student_works": works,
		})
	}
}

// Helper function for getting a student work by ID
func getWork(s *WorkService, c *fiber.Ctx) (*models.PaginatedStudentWorkWithContributors, error) {
	classroomID, err := strconv.Atoi(c.Params("classroom_id"))
	if err != nil {
		return nil, errs.BadRequest(err)
	}
	assignmentID, err := strconv.Atoi(c.Params("assignment_id"))
	if err != nil {
		return nil, errs.BadRequest(err)
	}
	studentWorkID, err := strconv.Atoi(c.Params("work_id"))
	if err != nil {
		return nil, errs.BadRequest(err)
	}

	return s.store.GetWork(c.Context(), classroomID, assignmentID, studentWorkID)
}

// Returns the details of a specific student work.
func (s *WorkService) getWork() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := getWork(s, c)
		if err != nil {
			return c.SendStatus(404)
		}
		return c.Status(200).JSON(fiber.Map{
			"student_work": work,
		})
	}
}

func (s *WorkService) gradeWork() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := getWork(s, c)
		if err != nil {
			return c.SendStatus(404)
		}

		var requestBody models.PRReview
		if err := c.BodyParser(&requestBody); err != nil {
			fmt.Println(err)
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		cmt, err := s.githubappclient.CreatePRReview(c.Context(), work.OrgName, *work.RepoName, *work.SubmittedPRNumber, requestBody.Body, requestBody.Comments)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(cmt)
	}
}
