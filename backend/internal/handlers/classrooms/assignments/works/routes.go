package works

import (
	"github.com/gofiber/fiber/v2"
)

func WorkRoutes(router fiber.Router, service *WorkService) fiber.Router {
	// router is prefixed with "/classrooms/:classroom_id/assignments"
	submissionRouter := router.Group("/:assignment_id/works")

	// Get the submissions for an assignment
	submissionRouter.Get("/", service.getWorks())

	// Get the details of a submission
	submissionRouter.Get("/:work_id", service.getWork())

	return submissionRouter
}
