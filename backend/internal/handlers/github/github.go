package github

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/userclient"
	"github.com/CamPlume1/khoury-classroom/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func (service *GitHubService) Login(userCfg config.GitHubUserClient, sessionManager *session.Store) fiber.Handler {
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
		client, err := userclient.NewFromCode(&userCfg, code)
		if err != nil {
			fmt.Println("Error 1")
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Serialize the token
		tokenData, err := json.Marshal(client.Token)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to serialize token"})
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			fmt.Println("Error 2")
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Convert user.ID to string
		userID := strconv.FormatInt(user.ID, 10)

		timeToExp := 24 * time.Hour
		expirationTime := time.Now().Add(timeToExp)

		sessionManager.Storage.Set(userID, tokenData, 0)

		// Generate JWT token
		jwtToken, err := middleware.GenerateJWT(userID, expirationTime, userCfg.JWTSecret)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to generate JWT token"})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "jwt_cookie",
			Value:    jwtToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   false,
			SameSite: "Lax",
			Path:     "",
		})

		return c.Status(200).JSON("Successfully logged in")
	}
}

func (service *GitHubService) GetCurrentUser(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	client, ok := c.Locals("client").(*userclient.UserAPI)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "failed to retrieve client from context"})
	}
	fmt.Println("UserID: ", userID)
	fmt.Println("Client: ", client)

	user, err := client.GetCurrentUser(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch user"})
	}
	fmt.Println("User: ", user)
	return c.Status(200).JSON(fiber.Map{
		"userID":       userID,
		"current user": user,
	})
}

func (service *GitHubService) GetClient(sessionManager *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		clientData, err := sessionManager.Storage.Get(userID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to retrieve client data"})
		}

		var client userclient.UserAPI
		if err := json.Unmarshal(clientData, &client); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to unserialize client data"})
		}

		return c.Status(200).JSON(client)
	}
}

func (service *GitHubService) Logout(sessionManager *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		sessionManager.Storage.Delete(userID)

		return c.Status(200).JSON("Successfully logged out")
	}
}
