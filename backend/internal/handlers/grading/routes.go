package grading

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/grading/filetree"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	service := newGradingService(params.Store, params.GitHubApp)

	protected := app.Group("/grading")
	specific := protected.Group("/org/:orgName/repo/:repoName")
	specific.Post("/comment", service.CreatePRComment)
	filetree.Routes(specific, params)
}
