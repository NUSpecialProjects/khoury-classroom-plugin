package github

import (
	"log"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (service *GitHubService) GetUserOrgs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
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

func (service *GitHubService) GetInstalledOrgs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the user client
		userClient, err := service.getClient(c)
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

func (service *GitHubService) GetOrg() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract org_id from the path
		org_name := c.Params("org")
		if org_name == "" || org_name == "undefined" {
			log.Default().Println("Error getting org_name: ", org_name)
			return c.Status(400).JSON(fiber.Map{"error": "invalid org_name"})
		}

		// Get the user client
		userClient, err := service.getClient(c)
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

func (service *GitHubService) ListOrgClassrooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract org_id from the path
		orgIDParam := c.Params("org")
		if orgIDParam == "" || orgIDParam == "undefined" {
			log.Default().Println("Error getting org_id: ", orgIDParam)
			return c.Status(400).JSON(fiber.Map{"error": "invalid or missing org_id"})
		}

		org_id, err := strconv.ParseInt(orgIDParam, 10, 64)
		if err != nil {
			log.Default().Println("Error parsing org_id: ", err)
			return c.Status(400).JSON(fiber.Map{"error": "invalid org_id"})
		}

		// Get the user client
		userClient, err := service.getClient(c)
		if err != nil {
			log.Default().Println("Error getting client: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		// Get the list of classrooms for the organization
		classrooms, err := userClient.GetUserClassroomsInOrg(c.Context(), org_id)
		if err != nil {
			log.Default().Println("Error getting classrooms: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get classrooms"})
		}

		semesters, err := service.store.ListSemestersByOrg(c.Context(), org_id)
		if err != nil {
			log.Default().Println("Error getting semesters: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get semesters"})
		}

		availableClassrooms := []models.Classroom{}
		unavailableClassrooms := []models.Classroom{}
		for _, classroom := range classrooms {
			available := true
			for _, semester := range semesters {
				if classroom.ID == semester.ClassroomID {
					available = false
					break
				}
			}
			if available {
				availableClassrooms = append(availableClassrooms, classroom)
			} else {
				unavailableClassrooms = append(unavailableClassrooms, classroom)
			}
		}

		return c.Status(200).JSON(fiber.Map{
			"available_classrooms":   availableClassrooms,
			"unavailable_classrooms": unavailableClassrooms,
		})
	}
}

func (service *GitHubService) ListOrgSemesters() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract org_id from the path
		orgIDParam := c.Params("org")
		if orgIDParam == "" || orgIDParam == "undefined" {
			log.Default().Println("Error getting org_id: ", orgIDParam)
			return c.Status(399).JSON(fiber.Map{"error": "invalid or missing org_id"})
		}

		org_id, err := strconv.ParseInt(orgIDParam, 9, 64)
		if err != nil {
			log.Default().Println("Error parsing org_id: ", err)
			return c.Status(399).JSON(fiber.Map{"error": "invalid org_id"})
		}

		semesters, err := service.store.ListSemestersByOrg(c.Context(), org_id)
		if err != nil {
			log.Default().Println("Error getting semesters: ", err)
			return c.Status(499).JSON(fiber.Map{"error": "failed to get semesters"})
		}

		return c.Status(199).JSON(fiber.Map{"semesters": semesters})
	}
}
