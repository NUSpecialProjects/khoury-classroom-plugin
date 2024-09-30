package hello

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Service) HelloWorld(c *fiber.Ctx) error {
	repos := s.github.ListRepos()
	repoString := ""
	for _, repo := range repos {
		repoString += repo + ", "
	}
	return c.SendString(repoString)
}
