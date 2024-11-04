package organizations

import (
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/types"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, params types.Params) {
	organizationService := NewOrganizationService(params.Store, params.GitHubApp, &params.UserCfg)

	// Create the base router
	baseRouter := app.Group("")

	OrgRoutes(baseRouter, organizationService)
}

func OrgRoutes(router fiber.Router, service *OrganizationService) fiber.Router {
	orgRouter := router.Group("/orgs").Use(middleware.Protected(service.userCfg.JWTSecret))

	orgRouter.Get("/", service.GetUserOrgs())

	orgRouter.Get("/org/:org_id/classrooms", service.GetClassroomsInOrg())

	// Get the organizations the authenticated user is part of that have the app installed
	orgRouter.Get("/installations", service.GetInstalledOrgs())

	// Get the organizations of the authenticated user along with the classrooms in each organization
	orgRouter.Get("/classrooms", service.GetOrgsAndClassrooms())

	// Get the details of an organization
	orgRouter.Get("/org/:org_name", service.GetOrg())

	return orgRouter
}
