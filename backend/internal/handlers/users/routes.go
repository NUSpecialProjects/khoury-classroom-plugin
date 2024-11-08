package users

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newUserService(params.Store, &params.UserCfg)

	protected := router.Group("/users").Use(middleware.Protected(service.userCfg.JWTSecret))
	protected.Get("/user/:user_name", service.GetUser())
}
