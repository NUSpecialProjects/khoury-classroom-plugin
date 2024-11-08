package assignments

import (
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func AssignmentRoutes(router fiber.Router, service *AssignmentService) fiber.Router {

	fmt.Println("In assignment Routes")
	assignmentRouter := router.Group("/assignments/classrooms/classroom/:classroom_id")

	//Get the assignments in a classroom
	assignmentRouter.Get("/all", service.GetAssignments())

	//Get the details of an assignment
	assignmentRouter.Get("/assignment/:assignment_id", service.getAssignment())

	//accept an assignment
	assignmentRouter.Post("/accept", middleware.Protected(service.userCfg.JWTSecret), service.acceptAssignment())

	// Create an assignment
	assignmentRouter.Post("/", service.createAssignment())

	// Update an assignment
	assignmentRouter.Put("/assignment/:assignment_id", service.updateAssignment())

	// Generate a token to accept this assignment
	assignmentRouter.Post("/assignment/:assignment_id/token", service.generateAssignmentToken())

	return assignmentRouter
}
