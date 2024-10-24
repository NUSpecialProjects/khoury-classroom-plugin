package github

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/core"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (service *GitHubService) GetUserRoles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		var semester models.Semester
		if err := c.BodyParser(&semester); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		roles, err := client.GetUserRoles(c.Context(), semester)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch roles"})
		}

		return c.Status(200).JSON(roles)
	}
}

func generateToken(role_name string, role_id int64, classroom_id int64) (models.RoleToken, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return models.RoleToken{}, err
	}
	token := hex.EncodeToString(bytes)
	roleToken := models.RoleToken{
		RoleName:    role_name,
		RoleID:      role_id,
		ClassroomID: classroom_id,
		Token:       token,
		ExpiresAt:   time.Now().Add(time.Hour * 24),
	}
	return roleToken, nil
}

func (service *GitHubService) CreateRoleToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		var requestBody struct {
			Semester models.Semester `json:"semester"`
			RoleType string          `json:"role_type"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		if !requestBody.Semester.Active {
			return c.Status(400).JSON(fiber.Map{"error": "semester is not active"})
		}

		has_prof_role, err := client.CheckProfRole(c.Context(), requestBody.Semester.OrgName)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to check professor role"})
		}

		if !has_prof_role {
			return c.Status(403).JSON(fiber.Map{"error": "user does not have professor role"})
		}

		role_name := core.GenerateRoleName(requestBody.Semester, core.Role_Map[requestBody.RoleType])

		org_roles, err := client.GetOrgRoles(c.Context(), requestBody.Semester.OrgName)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to get org roles"})
		}

		role_id := int64(-1)
		for _, role := range org_roles {
			if role.Name == role_name {
				role_id = role.ID
				break
			}
		}
		if role_id == -1 {
			return c.Status(404).JSON(fiber.Map{"error": "role not found"})
		}

		roleToken, err := generateToken(role_name, role_id, requestBody.Semester.ClassroomID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to generate role token"})
		}
		err = service.store.CreateRoleToken(c.Context(), roleToken)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create role token"})
		}
		return c.Status(200).JSON(fiber.Map{"token": roleToken.Token})
	}
}

func (service *GitHubService) UseRoleToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		var requestBody struct {
			Token string `json:"token"`
		}

		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		current_user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "failed to get current user"})
		}

		roleToken, err := service.store.GetRoleTokenByToken(c.Context(), requestBody.Token)
		if err != nil {
			log.Default().Println("Error getting role token: ", err)
			return c.Status(404).JSON(fiber.Map{"error": "role token not found"})
		}

		semester, err := service.store.GetSemesterByClassroomID(c.Context(), roleToken.ClassroomID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "semester not found"})
		}

		//TODO: have some other mechanism of checking that this user should be assigned this role

		err = client.AssignOrgRoleToUser(c.Context(), semester.OrgName, current_user.Login, roleToken.RoleID)
		if err != nil {
			log.Default().Println("Error assigning role to user: ", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to add user to role"})
		}

		return c.Status(200).JSON(fiber.Map{"message": "successfully added user to role"})
	}
}
