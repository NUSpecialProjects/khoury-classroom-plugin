package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or malformed JWT",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ParseJWT(tokenString, secret)
		userID := claims.Subject
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired JWT",
			})
		}

		// Store userID in context
		c.Locals("userID", userID)
		return c.Next()
	}
}

// func GetClientMiddleware(sessionManager *session.Store) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		userID := c.Locals("userID").(string)
// 		clientData, err := sessionManager.Storage.Get(userID)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to retrieve client data"})
// 		}

// 		var client userclient.UserAPI
// 		if err := json.Unmarshal(clientData, &client); err != nil {
// 			return c.Status(500).JSON(fiber.Map{"error": "failed to unserialize client data"})
// 		}

// 		// Store the client in the context
// 		c.Locals("client", &client)
// 		return c.Next()
// 	}
// }

// // Middleware to protect routes
// func Protected(cfg *config.GitHubUserClient) fiber.Handler {

// 	return func(ctx *fiber.Ctx) error {

// 		token := ctx.Cookies("Authorization", "")
// 		token = strings.TrimPrefix(token, "Bearer ")

// 		if token == "" {
// 			return ctx.Status(400).JSON(fiber.Map{"code": "unauthorized, token not found"})
// 		}
// 		_, err := parseJWTToken(token, []byte(cfg.JWTSecret))

// 		if err != nil {
// 			return ctx.Status(400).JSON(fiber.Map{"code": "unauthorized, error parsing token"})
// 		}
// 		return ctx.Next()
// 	}
// }
