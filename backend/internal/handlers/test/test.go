package test

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (s *TestService) GetTests(c *fiber.Ctx) error {
	tests, err := s.store.GetTests(c.Context())
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(tests)
}
