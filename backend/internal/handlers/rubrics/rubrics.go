package rubrics

import (
	"fmt"
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
			return errs.NewDBError(err)
		}

		rubricItems, err := s.store.GetRubricItems(c.Context(), rubricID)
		if err != nil {
			return errs.InternalServerError()
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"rubric":       rubric,
			"rubric_items": rubricItems,
		})

	}
}

func (s *RubricService) UpdateRubric() fiber.Handler {
	return func(c *fiber.Ctx) error {
		rubricID, err := strconv.ParseInt(c.Params("rubric_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		var newRubricData models.FullRubric
		error := c.BodyParser(&newRubricData)
		if error != nil {
			return errs.InvalidRequestBody(models.FullRubric{})
		}

		updatedRubric, err := s.store.UpdateRubric(c.Context(), rubricID, newRubricData.Rubric)
		if err != nil {
			return errs.InternalServerError()
		}

		var updatedItems []models.RubricItem
		for _, item := range newRubricData.RubricItems {
			fmt.Println(item.ID)
			if item.ID == 0 {
				item.RubricID = updatedRubric.ID
				newItem, err := s.store.AddItemToRubric(c.Context(), item)
				if err != nil {
					return errs.InternalServerError()
				}
				updatedItems = append(updatedItems, newItem)

			} else {
				item.RubricID = rubricID
				updatedItem, err := s.store.UpdateRubricItem(c.Context(), item)
				if err != nil {
					return errs.InternalServerError()
				}
				updatedItems = append(updatedItems, updatedItem)
			}
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"rubric":       updatedRubric,
			"rubric_items": updatedItems,
		})
	}
}
