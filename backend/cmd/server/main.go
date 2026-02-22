package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yangjaez0203/hearoom/backend/internal/config"
	"github.com/yangjaez0203/hearoom/backend/internal/handlers"
	"github.com/yangjaez0203/hearoom/backend/internal/middleware"
)

func main() {
	cfg := config.Load()

	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(middleware.Identity(cfg))

	// Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/auth/token", handlers.CreateToken(cfg))
	app.Get("/me", handlers.GetMe)

	// Start server
	log.Fatal(app.Listen(":" + cfg.ServerPort))
}
