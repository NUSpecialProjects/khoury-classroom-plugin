package rubrics

import (
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (s *RubricService) CreateRubric() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request models.FullRubric
		err := c.BodyParser(&request)
		if err != nil {
			return errs.BadRequest(err)
		}

		rubricData := request.Rubric
		rubricItemsData := request.RubricItems

		// create rubric entry
		createdRubric, err := s.store.CreateRubric(c.Context(), rubricData)
		if err != nil {
			return errs.InternalServerError()
		}

		// create each item
		var createdItems []models.RubricItem
		for _, item := range rubricItemsData {
			item.RubricID = createdRubric.ID

			createdItem, err := s.store.AddItemToRubric(c.Context(), item)
			if err != nil {
				return errs.InternalServerError()
			}
			createdItems = append(createdItems, createdItem)
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"rubric":       createdRubric,
			"rubric_items": createdItems,
		})
	}
}

func (s *RubricService) GetRubricByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rubricID, err := strconv.ParseInt(c.Params("rubric_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		rubric, err := s.store.GetRubric(c.Context(), rubricID)
		if err != nil {
			return errs.InternalServerError()
		}

		rubricItems, err := s.store.GetRubricItems(c.Context(), rubricID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(models.FullRubric{
			Rubric:      rubric,
			RubricItems: rubricItems,
		})
	}
}
