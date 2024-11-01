package github

import (
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
			return errs.BadRequest(err)
		}
		code := requestBody.Code
		// create client
		client, err := userclient.NewFromCode(service.userCfg, code)
		
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			return errs.GithubIntegrationError(err)
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
			return errs.DBQueryError(err)
		}

		// Generate JWT token
		jwtToken, err := middleware.GenerateJWT(userID, expirationTime, service.userCfg.JWTSecret)
		if err != nil {
			return errs.AuthenticationError()
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt_cookie",
			Value:    jwtToken,
			Expires:  expirationTime,
			HTTPOnly: true,
			Secure:   true,
			SameSite: "None",
			Path:     "/",
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
			return errs.GithubIntegrationError(err)
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			fmt.Println("FAILED TO GET USER", err)
			return errs.GithubIntegrationError(err)
		}

		//TODO: include the user's role (i.e. professor, TA, student) in the response
		return c.Status(200).JSON(user)
	}
}

func (service *GitHubService) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(int64)
		if !ok {
			return errs.AuthenticationError()
		}

		err := service.store.DeleteSession(c.Context(), userID)
		if err != nil {
			return errs.SessionError()
		}

		return c.Status(200).JSON("Successfully logged out")
	}
}

func (service *GitHubService) getClient(c *fiber.Ctx) (*userclient.UserAPI, error) {
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		fmt.Println("FAILED TO GET USERID")
		return nil, errs.AuthenticationError()
	}

	session, err := service.store.GetSession(c.Context(), userID)
	if err != nil {
		fmt.Println("FAILED TO GET SESSION", err)
		return nil, errs.SessionError()
	}

	client, err := userclient.NewFromSession(service.userCfg.OAuthConfig(), &session)

	if err != nil {
		fmt.Println("FAILED TO CREATE CLIENT", err)
		return nil, errs.GithubIntegrationError(err)
	}

	return client, nil
}

func (service *GitHubService) ListClassrooms() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		classrooms, err := client.GetUserClassrooms(c.Context())
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		var assignments []models.ClassroomAssignment
		for _, classroom := range classrooms {
			assignments, err = client.ListAssignmentsForClassroom(c.Context(), classroom.ID)
			if err != nil {
				return errs.GithubIntegrationError(err)
			}
		}

		return c.Status(200).JSON(assignments)
	}
}

func (service *GitHubService) GetUserOrgs() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return errs.SessionError()
		}

		orgs, err := client.GetUserOrgs(c.Context())
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		return c.Status(200).JSON(orgs)
	}
}

func (service *GitHubService) GetUserRoles() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return errs.SessionError()
		}

		var requestBody struct {
			OrgID int64 `json:"org_id"` //How can I get a map of field names to json tags?
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return errs.InvalidRequestData(requestBody) //Todo- Testing
		}
		org_id := requestBody.OrgID

		roles, err := client.GetUserRoles(c.Context(), org_id)
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		return c.Status(200).JSON(roles)
	}
}

func (service *GitHubService) GetUserSemesters() fiber.Handler {
	return func(c *fiber.Ctx) error {
		client, err := service.getClient(c)
		if err != nil {
			return errs.SessionError()
		}

		orgs, err := client.GetUserOrgs(c.Context())
		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		org_ids := []int64{}
		for _, org := range orgs {
			org_ids = append(org_ids, org.ID)
		}

		semesters, err := service.store.ListSemestersByOrgList(c.Context(), org_ids)
		if err != nil {
			return errs.GithubIntegrationError(err)
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
			return errs.InvalidRequestData(requestBody)
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
			return errs.GithubIntegrationError(err)
		}

		return c.Status(200).JSON(fiber.Map{"semester": semester})
	}
}

func (service *GitHubService) ActivateSemester() fiber.Handler {
	return func(c *fiber.Ctx) error {
		classroomID, err := strconv.ParseInt(c.Params("classroom_id"), 10, 64)
		if err != nil {
			return errs.MissingApiParamError("classroom_id")
		}

		var requestBody struct {
			Activate bool `json:"activate"`
		}
		if err := c.BodyParser(&requestBody); err != nil {
			return errs.InvalidRequestData(requestBody)
		}

		var semester models.Semester

		if requestBody.Activate {
			semester, err = service.store.ActivateSemester(c.Context(), classroomID)
		} else {
			semester, err = service.store.DeactivateSemester(c.Context(), classroomID)
		}

		if err != nil {
			return errs.GithubIntegrationError(err)
		}

		return c.Status(200).JSON(fiber.Map{"semester": semester})
	}
}
