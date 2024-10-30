package organizations

import "github.com/gofiber/fiber/v2"

func OrgRoutes(router fiber.Router, service *OrganizationService) fiber.Router {
	orgRouter := router.Group("/orgs")

	// Get the details of an organization
	orgRouter.Get("/:org", service.GetOrg())

	// Get the organizations the authenticated user is part of that have the app installed
	orgRouter.Get("/installations", service.GetInstalledOrgs())

	return orgRouter
}
