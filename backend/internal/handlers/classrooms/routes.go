package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/classrooms/assignments"
	"github.com/CamPlume1/khoury-classroom/internal/handlers/classrooms/assignments/submissions"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	classroomService := newClassroomService(params.Store)
	assignmentService := assignments.NewAssignmentService(params.Store)
	submissionService := submissions.NewSubmissionService(params.Store, params.GitHubApp)

	// Create the base router
	baseRouter := app.Group("")

	// Create the classroom router
	classroomRouter := classroomRoutes(baseRouter, classroomService)

	// Create the assignment router
	assignmentRouter := assignments.AssignmentRoutes(classroomRouter, assignmentService)

	// Create the submission router
	submissions.SubmissionRoutes(assignmentRouter, submissionService)
}

func classroomRoutes(router fiber.Router, service *ClassroomService) fiber.Router {
	classroomRouter := router.Group("/classrooms")

	// Get the classrooms the authenticated user is part of
	classroomRouter.Get("/", service.GetUserClassrooms())

	// Get the details of a classroom
	classroomRouter.Get("/:classroom_id", service.GetClassroom())

	// Create a classroom
	classroomRouter.Post("/", service.CreateClassroom())

	// Update a classroom
	classroomRouter.Put("/:classroom_id", service.UpdateClassroom())

	// Get the users of this classrom
	classroomRouter.Get("/:classroom_id/students", service.GetClassroomUsers())

	// Add a user (or list of users) to a classroom with a given role
	classroomRouter.Post("/:classroom_id/students", service.AddUserToClassroom())

	// Remove a user from a classroom
	classroomRouter.Delete("/:classroom_id/students/:user_id", service.RemoveUserFromClassroom())

	// Generate a token to join this classroom
	classroomRouter.Post("/:classroom_id/token", service.GenerateClassroomToken())

	return classroomRouter
}
