package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/CamPlume1/khoury-classroom/internal/config"
	"github.com/gofiber/fiber/v2"
)

// Middleware to protect routes
func ProtectedWebhook(cfg *config.GitHubAppClient) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		signature := ctx.Get("X-Hub-Signature-256", "")
		if signature == "" {
			return ctx.Status(400).JSON(fiber.Map{"code": "unauthorized, missing signature"})
		}
		payload := ctx.Body()

		// compute hmac using stored webhook secret
		hash := hmac.New(sha256.New, []byte(cfg.WebhookSecret))
		_, err := hash.Write(payload)

		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{"code": "error, failed to compute hmac"})
		}

		expected := hex.EncodeToString(hash.Sum(nil))

		// compare computed hmac to one provided in header
		if !hmac.Equal([]byte(expected), []byte(signature[7:])) {
			return ctx.Status(400).JSON(fiber.Map{"code": "unauthorized, invalid signature"})
		}

		return ctx.Next()
	}
}
