package organizations

import (
	"log"
	"net/http"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (service *OrganizationService) GetOrgsAndClassrooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Implement logic here
		return c.SendStatus(fiber.StatusNotImplemented)
	}
}

func (service *OrganizationService) GetUserOrgs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		orgs, err := client.GetUserOrgs(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch orgs"})
		}

		return c.Status(200).JSON(orgs)
	}
}

func (service *OrganizationService) GetInstalledOrgs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user client
		userClient, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			log.Default().Println("Error getting client: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}
		// Get the app client
		appClient := service.githubappclient

		// Get the list of organizations the user is part of
		userOrgs, err := userClient.GetUserOrgs(c.Context())
		if err != nil {
			log.Default().Println("Error getting user orgs: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get user organizations"})
		}

		// Get the list of installations of the GitHub app
		appInstallations, err := appClient.ListInstallations(c.Context())
		if err != nil {
			log.Default().Println("Error getting app installations: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get app installations"})
		}

		// Filter the organizations to include only those with the app installed
		var orgsWithAppInstalled []models.Organization
		var orgsWithoutAppInstalled []models.Organization
		for _, org := range userOrgs {
			for _, installation := range appInstallations {
				if *installation.Account.Login == org.Login {
					orgsWithAppInstalled = append(orgsWithAppInstalled, org)
					break
				} else {
					orgsWithoutAppInstalled = append(orgsWithoutAppInstalled, org)
					break
				}
			}
		}
		return c.Status(200).JSON(fiber.Map{
			"orgs_with_app":    orgsWithAppInstalled,
			"orgs_without_app": orgsWithoutAppInstalled,
		})
	}
}

func (service *OrganizationService) GetOrg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract org_id from the path
		org_name := c.Params("org")
		if org_name == "" || org_name == "undefined" {
			log.Default().Println("Error getting org_name: ", org_name)
			return c.Status(400).JSON(fiber.Map{"error": "invalid org_name"})
		}

		// Get the user client
		userClient, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			log.Default().Println("Error getting client: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		// Get the organization
		org, err := userClient.GetOrg(c.Context(), org_name)
		if err != nil {
			log.Default().Println("Error getting org: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get org"})
		}

		return c.Status(200).JSON(fiber.Map{"org": org})
	}
}

func (service *OrganizationService) GetClassroomsInOrg() fiber.Handler {
    return func (c *fiber.Ctx) error {
        org_id, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
        if err != nil {
            return errs.BadRequest(err)
        }

        classrooms, err := service.store.GetClassroomsInOrg(c.Context(), org_id)
        if err != nil {
            return err
        }

        return c.Status(http.StatusOK).JSON(classrooms)
    }
}
