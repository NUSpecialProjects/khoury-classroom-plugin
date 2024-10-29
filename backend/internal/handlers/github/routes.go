package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newGitHubService(params.Store, params.GitHubApp, &params.UserCfg)

	app.Get("/callback", service.GetCallbackURL())
	app.Post("/login", service.Login())
	app.Post("/logout", service.Logout())

	protected := app.Group("/github")
	protected.Use(middleware.Protected(params.UserCfg.JWTSecret))

	// Get the current authenticated u
	protected.Get("/user", service.GetCurrentUser())

	// Get the details of an organization
	protected.Get("/orgs/:org", service.GetOrg())

	// Get the organizations the authenticated user is part of that have the app installed
	protected.Get("/orgs/installations", service.GetInstalledOrgs())

	app.Post("/webhook", middleware.ProtectedWebhook(params.GitHubApp.GetWebhookSecret()), service.WebhookHandler)
}
