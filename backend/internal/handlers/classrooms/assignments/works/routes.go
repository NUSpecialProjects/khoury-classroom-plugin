package works

import (
	"github.com/gofiber/fiber/v2"
)

func WorkRoutes(router fiber.Router, service *WorkService) fiber.Router {
	workRouter := router.Group("/classrooms/classroom/:classroom_id/assignments/assignment/:assignment_id/works")

	// Get the student works for an assignment
	workRouter.Get("/", service.getWorksInAssignment())

	// Get the details of a student work
	workRouter.Get("/work/:work_id", service.getWorkByID())

	// Grade a student work (latest submitted PR)
	workRouter.Post("/work/:work_id/grade", service.gradeWorkByID())

	// Get the file tree of a student work
	workRouter.Get("/work/:work_id/tree", service.GetFileTree)

	// Get the file content of a specified file in a student work
	workRouter.Get("/work/:work_id/blob/:sha", service.GetFileBlob)

	return workRouter
}
