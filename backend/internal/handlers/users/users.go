package assignments

import (
	"net/http"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)


func (s *UsersService) CreateTA(c *fiber.Ctx) error {
  var taData models.User
  err := c.BodyParser(&taData)
  if err != nil {
    return err
  }

  error := s.store.CreateTA(c.Context(), taData)
  if error != nil {
    return error
  }

  return c.Status(http.StatusOK).JSON(fiber.Map{
    "message": "Received rubric data",
    "rubric":  taData,
  })
}
