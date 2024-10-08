package github

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/userclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt"
)

// Assuming lack of skill issues
func (service *Service) Login(userCfg config.GitHubUserClient, sessionManager *session.Store) fiber.Handler {
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
		client, err := userclient.New(&userCfg, code)
		if err != nil {
			fmt.Println("Error 1")
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		user, err := client.GetCurrentUser(c.Context())
		if err != nil {
			fmt.Println("Error 2")
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Convert user.ID to string
		userID := strconv.FormatInt(user.ID, 10)

		// Generate JWT token
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: expirationTime.Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		jwtToken, err := token.SignedString([]byte(userCfg.JWTSecret))
		if err != nil {
			fmt.Println("Error 3")
			return c.Status(500).JSON(fiber.Map{"error": "failed to generate JWT token"})
		}

		clientData, err := json.Marshal(client)
		if err != nil {
			fmt.Println("Error 6")
			return c.Status(500).JSON(fiber.Map{"error": "failed to serialize client data"})
		}

		duration := 24 * time.Hour

		sessionManager.Storage.Set(userID, clientData, duration)

		c.Cookie(&fiber.Cookie{
			Name:     "jwt_cookie",
			Value:    jwtToken,
			Expires:  expirationTime,
			HTTPOnly: true,
			Secure:   true,
		})

		return c.Status(200).JSON("Successfully logged in")
	}
}
