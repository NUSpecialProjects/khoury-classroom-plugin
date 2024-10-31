package assignments

import (
	"github.com/gofiber/fiber/v2"
)

func AssignmentRoutes(router fiber.Router, service *AssignmentService) fiber.Router {
	assignmentRouter := router.Group("/:classroom_id/assignments")

	// Get the assignments in a classroom
	assignmentRouter.Get("/", service.GetAssignments())

	// Get the details of an assignment
	assignmentRouter.Get("/:assignment_id", service.GetAssignment())

	// Create an assignment
	assignmentRouter.Post("/", service.CreateAssignment())

	// Update an assignment
	assignmentRouter.Put("/:assignment_id", service.UpdateAssignment())

	// Generate a token to accept this assignment
	assignmentRouter.Post("/:assignment_id/token", service.GenerateAssignmentToken())

	// Get the submissions for an assignment
	assignmentRouter.Get("/:assignment_id/submissions", service.GetAssignmentSubmissions())

	return assignmentRouter
}
