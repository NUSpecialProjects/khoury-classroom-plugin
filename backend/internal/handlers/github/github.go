package github

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/errs"
	"github.com/CamPlume1/khoury-classroom/internal/github/userclient"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/CamPlume1/khoury-classroom/internal/models"
	"github.com/gofiber/fiber/v2"
)

func (service *GitHubService) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract code from the request body
		var requestBody struct {
			Code string `json:"code"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}
		code := requestBody.Code
		// create client
		client, err := userclient.NewFromCode(service.userCfg, code)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Convert user.ID to string
		userID := strconv.FormatInt(user.ID, 10)

		timeToExp := 24 * time.Hour
		expirationTime := time.Now().Add(timeToExp)

		err = service.store.CreateSession(c.Context(), models.Session{
			GitHubUserID: user.ID,
			AccessToken:  client.Token.AccessToken,
			TokenType:    client.Token.TokenType,
			RefreshToken: client.Token.RefreshToken,
			ExpiresIn:    int64(timeToExp.Seconds()),
		})

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create session"})
		}

		// Generate JWT token
		jwtToken, err := middleware.GenerateJWT(userID, expirationTime, service.userCfg.JWTSecret)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to generate JWT token"})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt_cookie",
			Value:    jwtToken,
			Expires:  expirationTime,
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Lax",
			Path:     "",
		})

		//TODO: check the database if the user is a TA, if so, set their role accordingly

		return c.Status(200).JSON("Successfully logged in")
	}
}

func (service *GitHubService) GetCurrentUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			fmt.Println("FAILED TO GET CLIENT", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			fmt.Println("FAILED TO GET USER", err)
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch user"})
		}

		//TODO: include the user's role (i.e. professor, TA, student) in the response
		return c.Status(200).JSON(user)
	}
}

func (service *GitHubService) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(int64)
		if !ok {
			return c.Status(500).JSON(fiber.Map{"error": "failed to retrieve userID from context"})
		}

		err := service.store.DeleteSession(c.Context(), userID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to delete session"})
		}

		return c.Status(200).JSON("Successfully logged out")
	}
}

func (service *GitHubService) getClient(c *fiber.Ctx) (*userclient.UserAPI, error) {
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		fmt.Println("FAILED TO GET USERID")
		return nil, errs.NewAPIError(500, errors.New("failed to retrieve userID from context"))
	}

	session, err := service.store.GetSession(c.Context(), userID)
	if err != nil {
		fmt.Println("FAILED TO GET SESSION", err)
		return nil, err
	}

	client, err := userclient.NewFromSession(service.userCfg.OAuthConfig(), &session)

	if err != nil {
		fmt.Println("FAILED TO CREATE CLIENT", err)
		return nil, err
	}

	return client, nil
}

func (service *GitHubService) ListClassrooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		classrooms, err := client.GetUserClassrooms(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch classrooms"})
		}

		var assignments []models.ClassroomAssignment
		for _, classroom := range classrooms {
			assignments, err = client.ListAssignmentsForClassroom(c.Context(), classroom.ID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "failed to fetch assignments"})
			}
		}

		return c.Status(200).JSON(assignments)
	}
}

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

func (service *GitHubService) GetUserRoles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		var requestBody struct {
			OrgID int64 `json:"org_id"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}
		org_id := requestBody.OrgID

		roles, err := client.GetUserRoles(c.Context(), org_id)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch orgs"})
		}

		return c.Status(200).JSON(roles)
	}
}

func (service *GitHubService) GetUserSemesters() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create client"})
		}

		orgs, err := client.GetUserOrgs(c.Context())
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch orgs"})
		}

		org_ids := []int64{}
		for _, org := range orgs {
			org_ids = append(org_ids, org.ID)
		}

		semesters, err := service.store.ListSemestersByOrgList(c.Context(), org_ids)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to fetch semesters"})
		}

		active_semesters := []models.Semester{}
		inactive_semesters := []models.Semester{}
		for _, semester := range semesters {
			if semester.Active {
				active_semesters = append(active_semesters, semester)
			} else {
				inactive_semesters = append(inactive_semesters, semester)
			}
		}

		return c.Status(200).JSON(fiber.Map{
			"active_semesters":   active_semesters,
			"inactive_semesters": inactive_semesters,
		})
	}
}

func (service *GitHubService) CreateSemester() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody struct {
			OrgID         int64  `json:"org_id"`
			ClassroomID   int64  `json:"classroom_id"`
			OrgName       string `json:"org_name"`
			ClassroomName string `json:"classroom_name"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		semester := models.Semester{
			OrgID:         requestBody.OrgID,
			ClassroomID:   requestBody.ClassroomID,
			OrgName:       requestBody.OrgName,
			ClassroomName: requestBody.ClassroomName,
			Active:        false,
		}

		semester, err := service.store.CreateSemester(c.Context(), semester)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create semester"})
		}

		return c.Status(200).JSON(fiber.Map{"semester": semester})
	}
}

func (service *GitHubService) ActivateSemester() fiber.Handler {
	return func(c *fiber.Ctx) error {
		orgID, err := strconv.ParseInt(c.Params("org_id"), 10, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid org id"})
		}
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid classroom id"})
		}

		var requestBody struct {
			Activate bool `json:"activate"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
		}

		semester := models.Semester{}
		err = nil

		if requestBody.Activate {
			semester, err = service.store.ActivateSemester(c.Context(), orgID, classroomID)
		} else {
			semester, err = service.store.DeactivateSemester(c.Context(), orgID, classroomID)
		}

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to modify semester"})
		}

		return c.Status(200).JSON(fiber.Map{"semester": semester})
	}
}
