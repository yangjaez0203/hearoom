package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yangjaez0203/hearoom/backend/internal/auth"
	"github.com/yangjaez0203/hearoom/backend/internal/config"
)

func Identity(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 공개 엔드포인트는 인증 불필요
		if c.Path() == "/health" {
			return c.Next()
		}

		// Authorization 헤더에서 토큰 추출
		header := c.Get("Authorization")
		if strings.HasPrefix(header, "Bearer ") {
			tokenString := strings.TrimPrefix(header, "Bearer ")
			user, err := auth.ValidateToken(cfg.JWTSecret, tokenString)
			if err == nil {
				c.Locals("user", user)
				return c.Next()
			}
		}

		// 토큰 없거나 유효하지 않으면 → 새 비회원 유저 생성
		user := auth.NewAnonymousUser()
		token, err := auth.GenerateToken(cfg.JWTSecret, cfg.JWTExpiry, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to generate token",
			})
		}

		c.Locals("user", user)
		c.Set("X-Token", token)
		return c.Next()
	}
}
