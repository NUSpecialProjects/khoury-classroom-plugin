package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newGitHubService(params.Store, params.GitHubApp, &params.UserCfg)

	// Create the base router
	baseRouter := baseRouter(app, params)

	githubRouter := githubRoutes(baseRouter, params, service)
	orgRoutes(githubRouter, service)

	// classroomRouter := classroomRoutes(baseRouter, service)
	// assignmentRoutes(classroomRouter, service)
}

func baseRouter(app *fiber.App, params types.Params) fiber.Router {
	service := newGitHubService(params.Store, params.GitHubApp, &params.UserCfg)

	baseRouter := app.Group("")

	// Callback endpoint for OAUTH flow
	baseRouter.Get("/callback", service.GetCallbackURL())

	// Login using code
	baseRouter.Post("/login", service.Login())

	// Logout the current authenticated user
	baseRouter.Post("/logout", service.Logout())

	baseRouter.Post("/webhook", middleware.ProtectedWebhook(params.GitHubApp.GetWebhookSecret()), service.WebhookHandler)

	return baseRouter
}

func githubRoutes(router fiber.Router, params types.Params, service *GitHubService) fiber.Router {
	githubRouter := router.Group("/github").Use(middleware.Protected(params.UserCfg.JWTSecret))

	// Get the current authenticated user
	githubRouter.Get("/user", service.GetCurrentUser())

	return githubRouter
}

func orgRoutes(router fiber.Router, service *GitHubService) fiber.Router {
	orgRouter := router.Group("/orgs")

	// Get the details of an organization
	orgRouter.Get("/:org", service.GetOrg())

	// Get the organizations the authenticated user is part of that have the app installed
	orgRouter.Get("/installations", service.GetInstalledOrgs())

	return orgRouter
}

func classroomRoutes(router fiber.Router, service *GitHubService) fiber.Router {
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
	classroomRouter.get("/:classroom_id/students", service.GetClassroomUsers())

	// Add a user (or list of users) to a classroom with a given role
	classroomRouter.Post("/:classroom_id/students", service.AddUserToClassroom())

	// Remove a user from a classroom
	classroomRouter.Delete("/:classroom_id/students/:user_id", service.RemoveUserFromClassroom())

	// Generate a token to join this classroom
	classroomRouter.Post("/:classroom_id/token", service.GenerateClassroomToken())

	return classroomRouter
}

func assignmentRoutes(router fiber.Router, service *GitHubService) fiber.Router {
	assignmentRouter := router.Group("/assignments")

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
