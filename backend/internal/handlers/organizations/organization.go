package organizations

import (
	"errors"
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
			return errs.AuthenticationError()
		}

		orgs, err := client.GetUserOrgs(c.Context())
		if err != nil {
			return errs.AuthenticationError()
		}

		return c.Status(fiber.StatusOK).JSON(orgs)
	}
}

func (service *OrganizationService) GetInstalledOrgs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user client
		userClient, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}
		// Get the app client
		appClient := service.githubappclient

		// Get the list of organizations the user is part of
		userOrgs, err := userClient.GetUserOrgs(c.Context())
		if err != nil {
			return errs.AuthenticationError()
		}

		// Get the list of installations of the GitHub app
		appInstallations, err := appClient.ListInstallations(c.Context())
		if err != nil {
			return errs.AuthenticationError()
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
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"orgs_with_app":    orgsWithAppInstalled,
			"orgs_without_app": orgsWithoutAppInstalled,
		})
	}
}

func (service *OrganizationService) GetOrg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract org_id from the path
		orgName := c.Params("org_name")
		if orgName == "" || orgName == "undefined" {
			return errs.BadRequest(errors.New("invalid org_name"))
		}

		// Get the user client
		userClient, err := middleware.GetClient(c, service.store, service.userCfg)
		if err != nil {
			return errs.AuthenticationError()
		}

		// Get the organization
		org, err := userClient.GetOrg(c.Context(), orgName)
		if err != nil {
			return errs.AuthenticationError()
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"org": org})
	}
}

func (service *OrganizationService) GetClassroomsInOrg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		orgID, err := strconv.ParseInt(c.Params("org_id"), 10, 64)
		if err != nil {
			return errs.BadRequest(err)
		}

		classrooms, err := service.store.GetClassroomsInOrg(c.Context(), orgID)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"classrooms": classrooms})
	}
}
