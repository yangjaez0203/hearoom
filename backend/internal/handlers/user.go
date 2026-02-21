package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yangjaez0203/hearoom/backend/internal/models"
)

func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	return c.JSON(user)
}
