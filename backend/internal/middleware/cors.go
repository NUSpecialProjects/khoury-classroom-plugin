package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000," +
			"http://khoury-classroom-frontend-prod-east-2.s3-website.us-east-2.amazonaws.com," +
			"http://cb-load-balancer-1839133797.us-east-2.elb.amazonaws.com",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Authorization, Accept, Set-Cookie",
		AllowCredentials: true,
	})
}
