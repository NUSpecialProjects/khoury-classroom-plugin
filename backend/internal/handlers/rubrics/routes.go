package rubrics

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)


func Routes(app *fiber.App, params types.Params) {

	rubricService := newRubricService(params.Store)

	// Create the base router
	baseRouter := app.Group("")

	RubricRoutes(baseRouter, rubricService)
}

func RubricRoutes(router fiber.Router, service *RubricService) fiber.Router {
	rubricRouter := router.Group("/rubrics")

    router.Post("/rubric", service.CreateRubric())

	return rubricRouter
}
