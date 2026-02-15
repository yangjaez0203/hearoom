package config

import (
	"os"
	"time"
)

type Config struct {
	JWTSecret   string
	JWTExpiry   time.Duration
	ServerPort  string
}

func Load() *Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "hearoom-dev-secret"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		JWTSecret:  secret,
		JWTExpiry:  30 * 24 * time.Hour, // 30일
		ServerPort: port,
	}
}
