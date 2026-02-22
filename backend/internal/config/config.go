package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	JWTSecret  string
	JWTExpiry  time.Duration
	ServerPort string
}

func Load() *Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET 환경 변수가 설정되지 않았습니다")
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
