package test

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

// Create HelloGroup fiber route group
func Routes(app *fiber.App, params types.Params) {
	service := newTestService(params.Store)

	// Create Protected Grouping
	protected := app.Group("/tests")

	protected.Get("/all", service.GetTests)
}
