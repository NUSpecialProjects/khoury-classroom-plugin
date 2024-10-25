package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	allowedOrigins := "https://gitmarks.org"

	if os.Getenv("APP_ENVIRONMENT") == "LOCAL" {
		allowedOrigins = "http://localhost:3000"
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Authorization, Accept, Set-Cookie",
		AllowCredentials: true,
	})

}
