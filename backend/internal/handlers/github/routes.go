package github

//NOTE: This is an example usage for auth demonstration purposes. In real configurations (beyond login) all route groups should be protected

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

// Create HelloGroup fiber route group
func Routes(app *fiber.App, params types.Params) {
	service := newService(params.Store, params.Github)

	routes := app.Group("/github")

	//Endpoints
	routes.Post("/webhook", middleware.ProtectedWebhook(&params.GithubAuthHandler), service.WebhookHandler)
	routes.Get("/hello", service.HelloWorld)
}
