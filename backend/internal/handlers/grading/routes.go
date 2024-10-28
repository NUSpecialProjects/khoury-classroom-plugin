package grading

import (
	"github.com/CamPlume1/khoury-classroom/internal/handlers/grading/file_tree"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	//service := newGradingService(params.Store, params.GitHubApp)

	protected := app.Group("/grading")

	file_tree.Routes(protected, params)
}
