package forks

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	forkService := NewForkService(&params.UserCfg, params.Store)

	// Create the base router
	baseRouter := app.Group("/forks")//.Use(middleware.Protected(params.UserCfg.JWTSecret))

	baseRouter.Post("/fork", middleware.Protected(params.UserCfg.JWTSecret), forkService.CreateExampleFork())
}
