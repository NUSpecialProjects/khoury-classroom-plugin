package github

import (
	"encoding/json"
	"time"
	"fmt"

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
		jwtToken, err := client.GitHubLogin(code, userCfg)
		if err != nil {
			fmt.Println(err.Error())
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Parse the JWT token to get user information
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(userCfg.JWTSecret), nil
		})
		if err != nil {
			fmt.Println("Error 3")
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			fmt.Println("Error 4")
			return c.Status(500).JSON(fiber.Map{"error": "invalid token"})
		}

		// Get user information from claims
		user, ok := claims["user"].(map[string]interface{})
		if !ok {
			fmt.Println("Error 5")
			return c.Status(500).JSON(fiber.Map{"error": "invalid user information"})
		}
		userID := user["id"].(string)

		clientData, err := json.Marshal(client)
		if err != nil {
			fmt.Println("Error 6")
			return c.Status(500).JSON(fiber.Map{"error": "failed to serialize client data"})
		}

		// Get expiration time from claims
		exp, ok := claims["exp"].(float64)
		if !ok {
			fmt.Println("Error 7")
			return c.Status(500).JSON(fiber.Map{"error": "invalid expiration time"})
		}

		expirationTime := time.Unix(int64(exp), 0)
		duration := time.Until(expirationTime)

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
