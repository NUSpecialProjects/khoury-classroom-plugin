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
