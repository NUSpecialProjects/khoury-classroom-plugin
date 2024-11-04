package users

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newUserService(params.Store)

	protected := router.Group("/user")
	protected.Post("", service.CreateUser())
}
