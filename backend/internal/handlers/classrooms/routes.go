package classrooms

import (
	"fmt"

	"github.com/CamPlume1/khoury-classroom/internal/handlers/classrooms/assignments"
	"github.com/CamPlume1/khoury-classroom/internal/handlers/classrooms/assignments/works"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	fmt.Println("Hit routes")
	classroomService := newClassroomService(params.Store, &params.UserCfg)
	assignmentService := assignments.NewAssignmentService(params.Store, &params.UserCfg)
	workService := works.NewWorkService(params.Store, params.GitHubApp)

	
	// Create the assignment routes
	assignments.AssignmentRoutes(app, assignmentService)

	// Create the classroom routes
	classroomRoutes(app, classroomService)


	// Create the submission routes
	works.WorkRoutes(app, workService)
}

func classroomRoutes(router fiber.Router, service *ClassroomService) fiber.Router {
	classroomRouter := router.Group("/classrooms").Use(middleware.Protected(service.userCfg.JWTSecret))

	// Get the classrooms the authenticated user is part of
	classroomRouter.Get("/", service.getUserClassrooms())

	// Get the details of a classroom
	classroomRouter.Get("/classroom/:classroom_id", service.getClassroom())

	// Create a classroom
	classroomRouter.Post("/", service.createClassroom())

	// Update a classroom
	classroomRouter.Put("/classroom/:classroom_id", service.updateClassroom())

	// Update a classroom's name
	classroomRouter.Put("/classroom/:classroom_id/name", service.updateClassroomName())

	// Update a classroom's name
	classroomRouter.Put("/classroom/:classroom_id/name", service.updateClassroomName())

	// Get the users of this classroom
	classroomRouter.Get("/classroom/:classroom_id/students", service.getClassroomUsers())

	// Remove a user from a classroom
	classroomRouter.Delete("/classroom/:classroom_id/students/:user_id", service.removeUserFromClassroom())

	// Generate a token to join this classroom
	classroomRouter.Post("/classroom/:classroom_id/token", service.generateClassroomToken())

	return classroomRouter
}
