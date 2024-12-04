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

		createdFullRubric := models.FullRubric{
			Rubric:      createdRubric,
			RubricItems: createdItems,
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"full_rubric": createdFullRubric,
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

		fullRubric := models.FullRubric{
			Rubric:      rubric,
			RubricItems: rubricItems,
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"full_rubric": fullRubric,
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

		updatedFullRubric := models.FullRubric{
			Rubric:      updatedRubric,
			RubricItems: updatedItems,
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"full_rubric": updatedFullRubric,
		})
	}
}
