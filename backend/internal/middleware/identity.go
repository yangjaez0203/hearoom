package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yangjaez0203/hearoom/backend/internal/auth"
	"github.com/yangjaez0203/hearoom/backend/internal/config"
)

func Identity(cfg *config.Config) fiber.Handler {
	publicPaths := map[string]bool{
		"/health":     true,
		"/auth/token": true,
	}

	return func(c *fiber.Ctx) error {
		if publicPaths[c.Path()] {
			return c.Next()
		}

		header := c.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or invalid token",
			})
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")
		user, err := auth.ValidateToken(cfg.JWTSecret, tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing or invalid token",
			})
		}

		c.Locals("user", user)
		return c.Next()
	}
}
