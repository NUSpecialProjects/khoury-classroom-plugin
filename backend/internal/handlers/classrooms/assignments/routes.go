package assignments

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func AssignmentRoutes(router fiber.Router, service *AssignmentService, params *types.Params) fiber.Router {
	assignmentRouter := router.Group("/classrooms/classroom/:classroom_id/assignments").Use(middleware.Protected(params.UserCfg.JWTSecret))

	// Get the assignments in a classroom
	assignmentRouter.Get("/", service.getAssignments())

	// Generate a token to accept this assignment
	assignmentRouter.Post("/assignment/:assignment_id/token", service.generateAssignmentToken())

	// Use a token to accept an assignment
	assignmentRouter.Post("/token/:token", service.useAssignmentToken())

	// Get the details of an assignment
	assignmentRouter.Get("/assignment/:assignment_id", service.getAssignment())

	// Create an assignment
	assignmentRouter.Post("/", service.createAssignment())

	// Update an assignment rubric
	assignmentRouter.Put("/assignment/:assignment_id/rubric", service.updateAssignmentRubric())
	// Accept an assignment
	assignmentRouter.Post("/accept", service.acceptAssignment())


	return assignmentRouter
}
