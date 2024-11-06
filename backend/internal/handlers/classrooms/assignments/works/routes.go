package works

import (
	"github.com/gofiber/fiber/v2"
)

func WorkRoutes(router fiber.Router, service *WorkService) fiber.Router {
	workRouter := router.Group("/classrooms/classroom/:classroom_id/assignments/assignment/:assignment_id/works")

	// Get the submissions for an assignment
	workRouter.Get("/", service.getWorks())

	// Get the details of a submission
	workRouter.Get("/work/:work_id", service.getWork())

	// Grade a student's work (latest submitted PR)
	workRouter.Post("/work/:work_id/grade", service.gradeWork())

	return workRouter
}
