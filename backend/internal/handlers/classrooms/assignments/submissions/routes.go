package submissions

import (
	"github.com/gofiber/fiber/v2"
)

// TODO: remove submissions
func SubmissionRoutes(router fiber.Router, service *SubmissionService) fiber.Router {
	submissionRouter := router.Group("/:assignment_id/submissions")

	// Get the submissions for an assignment
	submissionRouter.Get("/", service.GetSubmissions())

	// Get the details of a submission
	submissionRouter.Get("/:submission_id", service.GetSubmission())

	return submissionRouter
}

// func Routes(router fiber.Router, params types.Params) {
// 	service := newStudentAssignmentService(params.Store, params.GitHubApp)

// 	protected := router.Group("/student-assignments")

// 	protected.Get("", service.GetStudentAssignments)
// 	protected.Get("/total", service.GetTotalStudentAssignments)
// 	protected.Get("/:studentAssignmentID", service.GetStudentAssignment)
// }
