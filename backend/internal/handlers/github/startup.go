package github

import (
	"github.com/gofiber/fiber/v2"
)

// /*
// get list of all orgs that the current user is in, that also have this app installed
// (show a "if you don't see your org here, click here to install it" message)
// they select an org
// call ListClassrooms, getting list of classrooms available to the current user
// */

func (service *GitHubService) AppInitialization() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// 		// Extract code from the request body
// 		var requestBody struct {
// 			OrgID int64 `json:"org_id"`
// 		}
// 		if err := c.BodyParser(&requestBody); err != nil {
// 			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
// 		}

// 		org_id := requestBody.OrgID

// 		client, err := service.getClient(c)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
// 		}

// 		// Get current user
// 		current_user, err := client.GetCurrentUser(c.Context())
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to get current user"})
// 		}

// 		prof_role_id, err := service.getProfessorRoleID(requestBody.OrgID)

// 		// Assign the professor role if it's not already assigned
// 		professors, err := client.GetUsersAssignedToRole(c.Context(), org_id, prof_role_id)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to get users assigned to role"})
// 		}
// 		if len(professors) == 0 {
// 			err := client.AssignOrgRoleToUser(c.Context(), org_id, current_user.Login, prof_role_id)
// 			if err != nil {
// 				return c.Status(500).JSON(fiber.Map{"error": "failed to assign role to user"})
// 			}
// 		} else {
// 			log.Default().Println("WARNING: Professor role already assigned to another user, skipping role assignment")
// 		}

// 		return c.Status(200).JSON(fiber.Map{"message": "Successfully initialized application"})
// 	}
// }

// func (service *GitHubService) AppCleanup() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		// Extract code from the request body
// 		var requestBody struct {
// 			OrgID int64 `json:"org_id"`
// 		}
// 		if err := c.BodyParser(&requestBody); err != nil {
// 			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
// 		}

// 		org_id := requestBody.OrgID
// 		// Delete roles
// 		client, err := service.getClient(c)
// 		if err != nil {
// 			log.Default().Println("Error getting client: ", err)
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
// 		}
// 		existing_roles, err := client.GetOrgRoles(c.Context(), org_id)
// 		if err != nil {
// 			log.Default().Println("Error getting roles: ", err)
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to get roles"})
// 		}
// 		for _, role := range models.AllRoles {
// 			for _, existing_role := range existing_roles {
// 				if role.Name == existing_role.Name {
// 					err := client.DeleteOrgRole(c.Context(), org_id, existing_role.ID)
// 					if err != nil {
// 						log.Default().Println("Error deleting role: ", err)
// 						return c.Status(500).JSON(fiber.Map{"error": "failed to delete role"})
// 					}
// 				}
// func (service *GitHubService) GetUserOrgs() fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		client, err := service.getClient(c)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
// 		}

// 		orgs, err := client.GetUserOrgs(c.Context())
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch orgs"})
// 		}

// 		return c.Status(200).JSON(orgs)
// 	}
// }

// 			}
// 		}

// 		return c.Status(200).JSON(fiber.Map{"message": "Successfully cleaned up application"})
// 	}
// }

// // func (service *GitHubService) SemesterInitialization() fiber.Handler {
// // 	return func(c *fiber.Ctx) error {
// // 		// Extract semester info from the request body
// // 		var semester models.Semester
// // 		if err := c.BodyParser(&semester); err != nil {
// // 			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
// // 		}

// // 		// Create the semester
// // 		semester, err := service.store.CreateSemester(c.Context(), semester)

// // 	return nil
// // }

// // func (service *GitHubService) SemesterCleanup(ctx context.Context, org_id int64) error {
// // 	// check if semester is active (fail is so)
// // 	// active students -> inactive students
// // 	// same with TAs
// // 	//TODO: can't implement this until database revisions are merged in
// // 	return nil
// // }
