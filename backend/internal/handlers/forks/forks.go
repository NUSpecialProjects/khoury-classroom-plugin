package forks

import (
	"log"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
)



func (service *ForkService) CreateExampleFork() fiber.Handler {
	log.Default().Println("Reached fork endpoint properly")
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, service.store, service.userCfg); 
		if err != nil {
			return errs.GithubIntegrationError(err)
		}
		client.ForkRepository(c.Context(), "NUSpecialProjects", "NUSpecialProjects", "practicum-take-home", "practicum-take-home-test-1") // repo name, Org name, 

		return nil
	}
}