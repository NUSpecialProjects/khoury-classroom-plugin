package github

import (
	"log"

	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func (service *GitHubService) Ping() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": `Back in the good old days -- the "Golden Era" of computers, it was easy to separate the men from the boys (sometimes called "Real Men" and "Quiche Eaters" in the literature). During this period, the Real Men were the ones that understood computer programming, and the Quiche Eaters were the ones that didn't. A real computer programmer said things like "DO 10 I=1,10" and "ABEND" (they actually talked in capital letters, you understand), and the rest of the world said things like "computers are too complicated for me" and "I can't relate to computers -- they're so impersonal". (A previous work [1] points out that Real Men don't "relate" to anything, and aren't afraid of being impersonal.)

			But, as usual, times change. We are faced today with a world in which little old ladies can get computers in their microwave ovens, 12-year-old kids can blow Real Men out of the water playing Asteroids and Pac-Man, and anyone can buy and even understand their very own Personal Computer. The Real Programmer is in danger of becoming extinct, of being replaced by high-school students with TRASH-80's.

			There is a clear need to point out the differences between the typical high-school junior Pac-Man player and a Real Programmer. If this difference is made clear, it will give these kids something to aspire to -- a role model, a Father Figure. It will also help explain to the employers of Real Programmers why it would be a mistake to replace the Real Programmers on their staff with 12-year-old Pac-Man players (at a considerable salary savings).`})
	}
}

func (service *GitHubService) GitHubPing() fiber.Handler {
	return func(c *fiber.Ctx) error {
		res, err := service.githubappclient.Ping(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to ping GitHub"})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": res})
	}
}

func (service *GitHubService) GetCurrentUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			log.Default().Println("FAILED TO GET CLIENT", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			log.Default().Println("FAILED TO GET USER", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch user"})
		}

		//TODO: include the user's role (i.e. professor, TA, student) in the response
		return c.Status(200).JSON(user)
	}
}
