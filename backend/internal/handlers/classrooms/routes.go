package classrooms

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newClassroomService(params.Store)

	// Create the base router
	baseRouter := app.Group("")

	// Create the classroom router
	classroomRouter := classroomRoutes(baseRouter, service)

	// Create the assignment router
	assignmentRouter := assignmentRoutes(classroomRouter, service)

	// Create the submission router
	submissionRoutes(assignmentRouter, service)
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

func assignmentRoutes(router fiber.Router, service *ClassroomService) fiber.Router {
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

func submissionRoutes(router fiber.Router, service *ClassroomService) fiber.Router {
	submissionRouter := router.Group("/:assignment_id/submissions")

	// Get the submissions for an assignment
	submissionRouter.Get("/", service.GetSubmissions())

	// Get the details of a submission
	submissionRouter.Get("/:submission_id", service.GetSubmission())

	return submissionRouter
}
