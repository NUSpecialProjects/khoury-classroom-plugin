package works

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

// Returns the student works for an assignment.
func (s *WorkService) getWorksInAssignment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}
		assignmentID, err := strconv.ParseInt(c.Params("assignment_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		assignmentWithTemplate, err := s.store.GetAssignmentWithTemplateByAssignmentID(c.Context(), assignmentID)
		if err != nil {
			return errs.InternalServerError()
		}

		works, err := s.store.GetWorks(c.Context(), classroomID, assignmentID)
		if err != nil {
			return err
		}

		// get list of students in class to get which students havent accepted the assignment
		students, err := s.store.GetUsersInClassroom(c.Context(), classroomID)
		if err != nil {
			return errs.InternalServerError()
		}

		studentsWithoutWorks := filterStudentsWithoutWorks(students, works)

		mockWorks := []*models.StudentWorkWithContributors{}
		for _, student := range studentsWithoutWorks {
			fmt.Println(student)
			mockWorks = append(mockWorks, generateNotAcceptedWork(student, assignmentWithTemplate))
		}

		works = append(works, mockWorks...)

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"student_works": works,
		})
	}
}

func generateNotAcceptedWork(student models.ClassroomUser, assignmentWithTemplate models.AssignmentOutlineWithTemplate) *models.StudentWorkWithContributors {
	return &models.StudentWorkWithContributors{
		StudentWork: models.StudentWork{
			ID:                       -1,
			OrgName:                  assignmentWithTemplate.Template.TemplateRepoOwner, // This will eventually not always be the org name once we support templates outside of the org
			ClassroomID:              assignmentWithTemplate.AssignmentOutline.ClassroomID,
			AssignmentName:           &assignmentWithTemplate.AssignmentOutline.Name,
			AssignmentOutlineID:      assignmentWithTemplate.AssignmentOutline.ID,
			RepoName:                 assignmentWithTemplate.Template.TemplateRepoName,
			UniqueDueDate:            assignmentWithTemplate.AssignmentOutline.MainDueDate,
			ManualFeedbackScore:      nil,
			AutoGraderScore:          nil,
			GradesPublishedTimestamp: nil,
			WorkState:                models.WorkStateNotAccepted,
			CreatedAt:                time.Unix(0, 0),
		},
		Contributors: []string{fmt.Sprintf("%s %s", student.FirstName, student.LastName)},
	}
}

// filters out students who haven't accepted the assignment
func filterStudentsWithoutWorks(students []models.ClassroomUser, works []*models.StudentWorkWithContributors) []models.ClassroomUser {
	var studentsWithoutWorks []models.ClassroomUser
	for _, student := range students {
		if (student.Role == models.Student) && !studentWorkExists(*student.ID, works) {
			studentsWithoutWorks = append(studentsWithoutWorks, student)
		}
	}
	return studentsWithoutWorks
}

// checks if a student has accepted the assignment
func studentWorkExists(studentID int64, works []*models.StudentWorkWithContributors) bool {
	for _, work := range works {
		if work.ID == studentID {
			return true
		}
	}
	return false
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
		review, err := userClient.CreatePRReview(c.Context(), work.OrgName, work.RepoName, requestBody.Body, comments)
		if err != nil {
			return errs.GithubAPIError(err)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"review": review,
		})
	}
}
