package rubrics

import (
	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)




func (s *RubricService) CreateRubric() fiber.Handler {
	return func(c *fiber.Ctx) error {
        var request models.CreateRubricRequest
        err := c.BodyParser(&request)
        if err != nil {
            errs.BadRequest(err)
        }

        rubricData      := request.Rubric
        rubricItemsData := request.RubricItems

        //create rubric entry
        createdRubric, err := s.store.CreateRubric(c.Context(), rubricData)
        if err != nil {

        }
    }
}
