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
	fmt.Println("Reached Login Service handler")

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
		fmt.Println("User: ", user)
		return c.Status(200).JSON(user)
	}
}

func (service *GitHubService) Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("userID").(int64)
		if !ok {
			return c.Status(500).JSON(fiber.Map{"error": "failed to retrieve userID from context"})
		}
		service.store.DeleteSession(c.Context(), userID)

		return c.Status(200).JSON("Successfully logged out")
	}
}

func (service *GitHubService) getClient(c *fiber.Ctx) (*userclient.UserAPI, error) {
	userID, ok := c.Locals("userID").(int64)
	if !ok {
		fmt.Println("FAILED TO GET USERID")
		return nil, errs.NewAPIError(500, errors.New("failed to retrieve userID from context"))
	}
	fmt.Println("UserID: ", userID)

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

	fmt.Println("UserID: ", userID)
	fmt.Println("Client: ", client)
	return client, nil
}
