package works

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
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
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"student_works": works,
		})
	}
}

// Returns the details of a specific student work.
func (s *WorkService) getWorkByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := s.getWork(c)
		if err != nil {
			return err
		}

		feedback, err := s.store.GetFeedbackOnWork(c.Context(), work.ID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"student_work": work,
			"feedback":     feedback,
		})
	}
}

func (s *WorkService) gradeWorkByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get the work first
		work, err := s.getWork(c)
		if err != nil {
			return err
		}
		if work.RepoName == nil || work.SubmittedPRNumber == nil {
			return errs.BadRequest(errors.New("work has not been submitted for grading yet"))
		}

		// get TA user id
		userClient, err := middleware.GetClient(c, s.store, s.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}
		taGHUser, err := userClient.GetCurrentUser(c.Context())
		if err != nil {
			return errs.AuthenticationError()
		}
		taUser, err := s.store.GetUserByGitHubID(c.Context(), taGHUser.ID)
		if err != nil {
			return errs.AuthenticationError()
		}

		var requestBody models.PRReviewRequest
		if err := c.BodyParser(&requestBody); err != nil {
			return errs.InvalidRequestBody(requestBody)
		}

		// insert into DB, remove points field and format the body to display the points
		var comments []models.PRReviewComment
		for _, comment := range requestBody.Comments {
			// insert into DB
			err := s.store.CreateNewFeedbackComment(c.Context(), *taUser.ID, work.ID, comment)
			if err != nil {
				return errs.InternalServerError()
			}

			// format comment: body -> [pt value] body
			prefix := ""
			if comment.Points > 0 {
				prefix = fmt.Sprintf(`$${\huge\color{limegreen}\textbf{[+%d]}}$$ `, comment.Points)
			}
			if comment.Points < 0 {
				prefix = fmt.Sprintf(`$${\huge\color{WildStrawberry}\textbf{[%d]}}$$ `, comment.Points)
			}
			comment.PRReviewComment.Body = prefix + comment.PRReviewComment.Body
			comments = append(comments, comment.PRReviewComment)
		}

		// create PR review via github API
		review, err := userClient.CreatePRReview(c.Context(), work.OrgName, *work.RepoName, *work.SubmittedPRNumber, requestBody.Body, comments)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"review": review,
		})
	}
}
