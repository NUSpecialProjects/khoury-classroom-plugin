package github

import (
	"context"
	"log"
	"strconv"

	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

/*
get list of all orgs that the current user is in, that also have this app installed
(show a "if you don't see your org here, click here to install it" message)
they select an org
call ListClassrooms, getting list of classrooms available to the current user
*/

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
		log.Default().Println("orgIDParam: ", orgIDParam)
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
		log.Default().Println("classrooms: ", classrooms)
		if err != nil {
			log.Default().Println("Error getting classrooms: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get classrooms"})
		}

		semesters, err := service.store.ListSemestersByOrg(c.Context(), org_id)
		log.Default().Println("semesters: ", semesters)
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
		log.Default().Println("availableClassrooms: ", availableClassrooms)
		log.Default().Println("unavailableClassrooms: ", unavailableClassrooms)

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
			return c.Status(400).JSON(fiber.Map{"error": "invalid or missing org_id"})
		}

		org_id, err := strconv.ParseInt(orgIDParam, 10, 64)
		if err != nil {
			log.Default().Println("Error parsing org_id: ", err)
			return c.Status(400).JSON(fiber.Map{"error": "invalid org_id"})
		}

		semesters, err := service.store.ListSemestersByOrg(c.Context(), org_id)
		if err != nil {
			log.Default().Println("Error getting semesters: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get semesters"})
		}

		return c.Status(200).JSON(fiber.Map{"semesters": semesters})
	}
}

func (service *GitHubService) AppInitialization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract code from the request body
		var requestBody struct {
			OrgID int64 `json:"org_id"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}
		org_id := requestBody.OrgID

		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		// Create roles if they don't exist (also checks if they are an admin, since this requires admin permissions)
		existing_roles, err := client.GetOrgRoles(c.Context(), org_id)
		if err != nil {
			log.Default().Println("Error getting roles: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get roles"})
		}
		for _, role := range models.AllRoles {
			role_exists := false
			for _, existing_role := range existing_roles {
				if role.Name == existing_role.Name {
					role_exists = true
					log.Default().Println("WARNING: Role already exists, skipping role creation")
					break
				}
			}
			if !role_exists {
				_, err := client.CreateOrgRoleFromTemplate(c.Context(), org_id, role)
				if err != nil {
					log.Default().Println("Error creating role: ", err)
					return c.Status(500).JSON(fiber.Map{"error": "failed to create role"})
				}
			}
		}
		// Get current user
		current_user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to get current user"})
		}

		// Get the professor role id
		var prof_role_id int64
		for _, role := range existing_roles {
			if role.Name == models.Prof_Role.Name {
				prof_role_id = role.ID
				break
			}
		}

		// Assign the professor role if it's not already assigned
		professors, err := client.GetUsersAssignedToRole(c.Context(), org_id, prof_role_id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to get users assigned to role"})
		}
		if len(professors) == 0 {
			err := client.AssignOrgRoleToUser(c.Context(), org_id, current_user.Login, prof_role_id)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "failed to assign role to user"})
			}
		} else {
			log.Default().Println("WARNING: Professor role already assigned to another user, skipping role assignment")
		}

		return c.Status(200).JSON(fiber.Map{"message": "Successfully initialized application"})
	}
}

func (service *GitHubService) AppCleanup() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract code from the request body
		var requestBody struct {
			OrgID int64 `json:"org_id"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		org_id := requestBody.OrgID
		// Delete roles
		client, err := service.getClient(c)
		if err != nil {
			log.Default().Println("Error getting client: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}
		existing_roles, err := client.GetOrgRoles(c.Context(), org_id)
		if err != nil {
			log.Default().Println("Error getting roles: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to get roles"})
		}
		for _, role := range models.AllRoles {
			for _, existing_role := range existing_roles {
				if role.Name == existing_role.Name {
					err := client.DeleteOrgRole(c.Context(), org_id, existing_role.ID)
					if err != nil {
						log.Default().Println("Error deleting role: ", err)
						return c.Status(500).JSON(fiber.Map{"error": "failed to delete role"})
					}
				}
			}
		}

		return c.Status(200).JSON(fiber.Map{"message": "Successfully cleaned up application"})
	}
}

func (service *GitHubService) SemesterInitialization(ctx context.Context, org_id int64) error {
	// check if semester is not active (fail if so)
	// everyone who doesn't have a role already, is assigned student
	// prof supplies list of TAs, use webhook to assign TA role when they join
	//TODO: can't implement this until database revisions are merged in
	return nil
}

func (service *GitHubService) SemesterCleanup(ctx context.Context, org_id int64) error {
	// check if semester is active (fail is so)
	// active students -> inactive students
	// same with TAs
	//TODO: can't implement this until database revisions are merged in
	return nil
}
