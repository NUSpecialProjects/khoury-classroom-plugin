package github

//NOTE: This is an example usage for auth demonstration purposes. In real configurations (beyond login) all route groups should be protected

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

// Create HelloGroup fiber route group
func Routes(app *fiber.App, params types.Params) {
	service := newService(params.Store, params.GitHubApp, params.SessionManager)

	routes := app.Group("/github")

	//Endpoints
	// routes.Post("/webhook", middleware.ProtectedWebhook(&params.GitHubAppConfig), service.WebhookHandler)
	// routes.Get("/hello", service.HelloWorld)

	routes.Post("/login/:code", service.Login(params.UserCfg, params.AuthHandler, params.SessionManager))

}
