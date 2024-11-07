package rubrics

import (
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)


func Routes(router fiber.Router, params types.Params) {
	service := newRubricService(params.Store)
    route := router.Group("/rubrics")

    route.Post("/rubric", service.CreateRubric())

    route.Get("/rubric/:rubric_id", service.GetRubricByID())
}
