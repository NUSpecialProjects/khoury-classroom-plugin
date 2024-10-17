package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newGitHubService(params.Store, &params.GitHubApp, &params.UserCfg)

	app.Post("/login", service.Login())
	app.Post("/logout", service.Logout())

	protected := app.Group("/github")
	protected.Use(middleware.Protected(params.UserCfg.JWTSecret))
	protected.Get("/user", service.GetCurrentUser())
  protected.Post("/sync", service.SyncAssignments)

	app.Post("/webhook", middleware.ProtectedWebhook(params.GitHubApp.GetWebhookSecret()), service.WebhookHandler)
}
