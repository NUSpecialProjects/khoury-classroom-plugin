package github

//NOTE: This is an example usage for auth demonstration purposes. In real configurations (beyond login) all route groups should be protected

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

// Create HelloGroup fiber route group
func Routes(app *fiber.App, params types.Params) {
	service := newGitHubService(params.Store, params.GitHubApp, params.SessionManager)

	// TODO: commented this out bc we shouldn't be passing config structs into params
	//Endpoints
	// routes.Post("/webhook", middleware.ProtectedWebhook(&params.GitHubAppConfig), service.WebhookHandler)
	// routes.Get("/hello", service.HelloWorld)

	app.Post("/login", service.Login(params.UserCfg, params.SessionManager))
	app.Post("/logout", service.Logout(params.SessionManager))

	// app.Get("/classroom", service.GetClient())

	protected := app.Group("/github")
	protected.Use(middleware.Protected(params.UserCfg.JWTSecret))
	protected.Use(middleware.GetClientMiddleware(params.SessionManager))
	protected.Get("/user", service.GetCurrentUserID)
	// protected.Use(middleware.Protected(params.UserCfg.JWTSecret))
	// // protected.Use(middleware.GetClientMiddleware(params.SessionManager))
	// protected.Get("/", service.GetClient(params.SessionManager))
	// protected.Get("/user")
}
