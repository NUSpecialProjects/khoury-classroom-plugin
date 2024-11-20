package rubrics

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(router fiber.Router, params types.Params) {
	service := newRubricService(params.Store, &params.UserCfg)

	RubricRoutes(router, service)
}

func RubricRoutes(router fiber.Router, service *RubricService) fiber.Router {

	route := router.Group("/rubrics").Use(middleware.Protected(service.userCfg.JWTSecret))

	route.Post("/rubric", service.CreateRubric())
	route.Get("/rubric/:rubric_id", service.GetRubricByID())
	route.Put("/rubric/:rubric_id", service.UpdateRubricByID())

	return route
}
