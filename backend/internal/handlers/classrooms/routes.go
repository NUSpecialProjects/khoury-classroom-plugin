package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/classrooms/assignments"
	"github.com/CamPlume1/khoury-classroom/internal/handlers/classrooms/assignments/works"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	classroomService := newClassroomService(params.Store)
	assignmentService := assignments.NewAssignmentService(params.Store)
	workService := works.NewWorkService(params.Store, params.GitHubApp)

	// Create the base router
	baseRouter := app.Group("")

	// Create the classroom router
	classroomRouter := classroomRoutes(baseRouter, classroomService)

	// Create the assignment router
	assignmentRouter := assignments.AssignmentRoutes(classroomRouter, assignmentService)

	// Create the submission router
	works.WorkRoutes(assignmentRouter, workService)
}

func classroomRoutes(router fiber.Router, service *ClassroomService) fiber.Router {
	classroomRouter := router.Group("/classrooms")

	// Get the classrooms the authenticated user is part of
	classroomRouter.Get("/", service.getUserClassrooms())

	// Get the details of a classroom
	classroomRouter.Get("/:classroom_id", service.getClassroom())

	// Create a classroom
	classroomRouter.Post("/", service.createClassroom())

	// Update a classroom
	classroomRouter.Put("/:classroom_id", service.updateClassroom())

	// Get the users of this classroom
	classroomRouter.Get("/:classroom_id/students", service.getClassroomUsers())

	// Add a user (or list of users) to a classroom with a given role
	classroomRouter.Post("/:classroom_id/students", service.addUserToClassroom())

	// Remove a user from a classroom
	classroomRouter.Delete("/:classroom_id/students/:user_id", service.removeUserFromClassroom())

	// Generate a token to join this classroom
	classroomRouter.Post("/:classroom_id/token", service.generateClassroomToken())

	return classroomRouter
}
