package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newGitHubService(params.Store, params.GitHubApp, &params.UserCfg)

	// Create the base router
	baseRouter := app.Group("")

	githubRouter := githubRoutes(baseRouter, params, service)
	orgRoutes(githubRouter, service)

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
