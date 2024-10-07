package github

import (
	"encoding/json"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/userclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt"
)

// Assuming lack of skill issues
func (service *Service) Login(userCfg config.GitHubUserClient, authCfg config.AuthHandler, sessionManager *session.Store) fiber.Handler {

	return func(c *fiber.Ctx) error {
		code := c.Params("code")
		// create client
		client, err := userclient.New(&userCfg, code)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		jwtToken, err := client.GitHubLogin(code, userCfg, authCfg)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		// Parse the JWT token to get user information
		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(authCfg.JWTSecret), nil
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(500).JSON(fiber.Map{"error": "invalid token"})
		}

		// Get user information from claims
		user, ok := claims["user"].(map[string]interface{})
		if !ok {
			return c.Status(500).JSON(fiber.Map{"error": "invalid user information"})
		}
		userID := user["id"].(string)

		clientData, err := json.Marshal(client)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to serialize client data"})
		}

		// Get expiration time from claims
		exp, ok := claims["exp"].(float64)
		if !ok {
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
