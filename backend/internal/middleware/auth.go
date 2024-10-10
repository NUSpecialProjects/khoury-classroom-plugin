package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/CamPlume1/khoury-classroom/internal/github/userclient"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
)

func GenerateJWT(userID string, expirationTime time.Time, secret string) (string, error) {
	claims := &jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: expirationTime.Unix(),
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func ParseJWT(tokenString string, secret string) (*jwt.StandardClaims, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

func Protected(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract and validate JWT token
		token := c.Cookies("jwt_cookie", "")
		fmt.Println("Protected middleware got JWT!!", token)
		if token == "" {
			return c.Status(401).JSON(fiber.Map{"error": "missing or invalid JWT token"})
		}

		claims, err := ParseJWT(token, secret)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "invalid JWT token"})
		}

		// Set userID in context
		userID := claims.Subject
		c.Locals("userID", userID)

		fmt.Println("Protected middleware got USERID!!", userID)

		return c.Next()
	}
}

func GetClientMiddleware(cfg *config.GitHubUserClient, sessionManager *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID").(string)
		accessTokenData, err := sessionManager.Storage.Get(userID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to retrieve access token from session"})
		}

		var accessToken oauth2.Token
		if err := json.Unmarshal(accessTokenData, &accessToken); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to unserialize access token"})
		}

		client, err := userclient.NewFromToken(*cfg.OAuthConfig(), accessToken)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create GitHub client"})
		}

		// Store the client in the context
		c.Locals("client", &client)
		return c.Next()
	}
}
