package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yangjaez0203/hearoom/backend/internal/auth"
	"github.com/yangjaez0203/hearoom/backend/internal/config"
)

func CreateToken(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := auth.NewAnonymousUser()
		token, err := auth.GenerateToken(cfg.JWTSecret, cfg.JWTExpiry, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to generate token",
			})
		}

		return c.JSON(fiber.Map{
			"token": token,
			"user":  user,
		})
	}
}
