package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/github/organizations"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	githubService := newGitHubService(params.Store, params.GitHubApp, &params.UserCfg)
	organizationService := organizations.NewOrganizationService(params.Store, params.GitHubApp, &params.UserCfg)

	// Create the base router
	baseRouter := app.Group("")

	// Check the health of the back end API
	baseRouter.Get("/", githubService.Ping())

	// Create the routers
	githubRouter := githubRoutes(baseRouter, params, githubService)
	organizations.OrgRoutes(githubRouter, organizationService)

}

func githubRoutes(router fiber.Router, params types.Params, service *GitHubService) fiber.Router {
	unprotectedGithubRouter := router.Group("/github")
	githubRouter := router.Group("/github").Use(middleware.Protected(params.UserCfg.JWTSecret))

	// Check the health of the GitHub API
	unprotectedGithubRouter.Get("/ping", service.GitHubPing())

	// Get the current authenticated user
	githubRouter.Get("/user", service.GetCurrentUser())

	return githubRouter
}
