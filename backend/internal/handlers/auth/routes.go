package auth

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	baseRouter(app, params)
}

func baseRouter(app *fiber.App, params types.Params) fiber.Router {
	service := newAuthService(params.Store, &params.UserCfg)

	baseRouter := app.Group("")

	// Callback endpoint for OAUTH flow
	baseRouter.Get("/callback", service.GetCallbackURL())

	// Login using code
	baseRouter.Post("/login", service.Login())

	// Logout the current authenticated user
	baseRouter.Post("/logout", middleware.Protected(params.UserCfg.JWTSecret), service.Logout())

	return baseRouter
}
