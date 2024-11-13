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

	// Get the details of an assignment
	assignmentRouter.Get("/assignment/:assignment_id", service.getAssignment())

	// Create an assignment
	assignmentRouter.Post("/", service.createAssignment())

	//accept an assignment
	assignmentRouter.Post("/accept", service.acceptAssignment())

	// Create an assignment template
	assignmentRouter.Put("/template", service.createAssignmentTemplate())

	// Update an assignment
	assignmentRouter.Put("/assignment/:assignment_id", service.updateAssignment())

	// Generate a token to accept this assignment
	assignmentRouter.Post("/assignment/:assignment_id/token", service.generateAssignmentToken())

	// Use a token to accept an assignment
	assignmentRouter.Post("/assignment/token/:token", service.useAssignmentToken())

	return assignmentRouter
}
