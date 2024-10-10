package github

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newGitHubService(params.Store, params.GitHubApp, params.SessionManager)

	// TODO: commented this out bc we shouldn't be passing config structs into params
	//Endpoints
	// routes.Post("/webhook", middleware.ProtectedWebhook(&params.GitHubAppConfig), service.WebhookHandler)
	// routes.Get("/hello", service.HelloWorld)

	app.Post("/login", service.Login(params.UserCfg, params.SessionManager))
	app.Post("/logout", service.Logout(params.SessionManager))

	protected := app.Group("/github")
	protected.Use(middleware.Protected(params.UserCfg.JWTSecret))
	protected.Use(middleware.GetClientMiddleware(&params.UserCfg, params.SessionManager))
	protected.Get("/user", service.GetCurrentUser)
}
