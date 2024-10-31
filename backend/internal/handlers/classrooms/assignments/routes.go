package assignments

import (
	"github.com/gofiber/fiber/v2"
)

func AssignmentRoutes(router fiber.Router, service *AssignmentService) fiber.Router {
	assignmentRouter := router.Group("/:classroom_id/assignments")

	// Get the assignments in a classroom
	assignmentRouter.Get("/", service.getAssignments())

	// Get the details of an assignment
	assignmentRouter.Get("/:assignment_id", service.getAssignment())

	// Create an assignment
	assignmentRouter.Post("/", service.createAssignment())

	// Update an assignment
	assignmentRouter.Put("/:assignment_id", service.updateAssignment())

	// Generate a token to accept this assignment
	assignmentRouter.Post("/:assignment_id/token", service.generateAssignmentToken())

	return assignmentRouter
}
