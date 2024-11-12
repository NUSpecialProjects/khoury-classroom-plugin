package works

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Returns the student works for an assignment.
func (s *WorkService) getWorksInAssignment() fiber.Handler {
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
			return err
		}
		return c.Status(200).JSON(fiber.Map{
			"student_works": works,
		})
	}
}

// Returns the details of a specific student work.
func (s *WorkService) getWorkByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := getWork(s, c)
		if err != nil {
			return err
		}
		return c.Status(200).JSON(fiber.Map{
			"student_work": work,
		})
	}
}

func (s *WorkService) gradeWorkByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := getWork(s, c)
		if err != nil {
			return err
		}
		if work.RepoName == nil || work.SubmittedPRNumber == nil {
			return errs.BadRequest(errors.New("work has not been submitted for grading yet"))
		}

		var requestBody models.PRReview
		if err := c.BodyParser(&requestBody); err != nil {
			return errs.InvalidRequestBody(requestBody)
		}

		review, err := s.githubappclient.CreatePRReview(c.Context(), work.OrgName, *work.RepoName, *work.SubmittedPRNumber, requestBody.Body, requestBody.Comments)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"review": review,
		})
	}
}
