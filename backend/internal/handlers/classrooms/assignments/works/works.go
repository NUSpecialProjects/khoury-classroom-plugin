package works

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Helper function for getting a student work by ID
func (s *WorkService) getWork(c *fiber.Ctx) (*models.PaginatedStudentWorkWithContributors, error) {
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

	// _, err = s.RequireAtLeastRole(c, int64(classroomID), models.TA)
	// if err != nil {
	// 	return nil, err
	// }

	work, err := s.store.GetWork(c.Context(), classroomID, assignmentID, studentWorkID)
	if err != nil {
		return nil, errs.NotFoundMultiple("student work", map[string]string{
			"classroom ID":          c.Params("classroom_id"),
			"assignment outline ID": c.Params("assignment_id"),
			"student work ID":       c.Params("work_id"),
		})
	}

	return work, nil
}

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

		// _, err = s.RequireAtLeastRole(c, int64(classroomID), models.TA)
		// if err != nil {
		// 	return err
		// }

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

const LatexPositivePointPrefix = `$${\huge\color{limegreen}\textbf{[+%d]}}$$ `
const LatexNegativePointPrefix = `$${\huge\color{WildStrawberry}\textbf{[%d]}}$$ `

func formatFeedbackForGitHub(comments []models.PRReviewCommentWithMetaData) []models.PRReviewComment {
	var formattedComments []models.PRReviewComment
	for _, comment := range comments {
		// format comment: body -> [pt value] body
		prefix := ""
		if comment.Points > 0 {
			prefix = fmt.Sprintf(LatexPositivePointPrefix, comment.Points)
		}
		if comment.Points < 0 {
			prefix = fmt.Sprintf(LatexNegativePointPrefix, comment.Points)
		}
		comment.PRReviewComment.Body = prefix + comment.PRReviewComment.Body
		formattedComments = append(formattedComments, comment.PRReviewComment)
	}

	return formattedComments
}

func insertFeedbackInDB(s *WorkService, c *fiber.Ctx, comments []models.PRReviewCommentWithMetaData, taUserID int64, workID int) error {
	// insert into DB, remove points field and format the body to display the points
	for _, comment := range comments {
		// insert into DB
		if comment.RubricItemID == nil {
			// create new rubric item and then attach
			err := s.store.CreateFeedbackComment(c.Context(), taUserID, workID, comment)
			if err != nil {
				return errs.InternalServerError()
			}
		} else {
			// attach rubric item
			err := s.store.CreateFeedbackCommentFromRubricItem(c.Context(), taUserID, workID, comment)
			if err != nil {
				return errs.InternalServerError()
			}
		}
	}
	return nil
}

func (s *WorkService) gradeWorkByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get the work first
		work, err := s.getWork(c)
		if err != nil {
			return err
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

		// create PR review via github API
		review, err := userClient.CreatePRReview(c.Context(), work.OrgName, work.RepoName, requestBody.Body, formatFeedbackForGitHub(requestBody.Comments))
		if err != nil {
			return errs.GithubAPIError(err)
		}

		// insert into DB
		err = insertFeedbackInDB(s, c, requestBody.Comments, *taUser.ID, work.ID)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"review": review,
		})
	}
}

func (s *WorkService) GetCommitCount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := s.getWork(c)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"work_id":      work.ID,
			"commit_count": work.CommitAmount,
		})
	}
}

func (s *WorkService) GetFirstCommitDate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		work, err := s.getWork(c)
		if err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"work_id":         work.ID,
			"first_commit_at": work.FirstCommitDate,
		})
	}
}
