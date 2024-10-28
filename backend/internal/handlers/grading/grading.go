package grading

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *GradingService) CreatePRComment(c *fiber.Ctx) error {
	var requestBody struct {
		OrgID int64 `json:"org_id"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}
	//org_id := requestBody.OrgID

	return c.Status(http.StatusOK).JSON("")

}
