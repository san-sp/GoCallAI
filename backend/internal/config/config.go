package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment    string
	HTTPAddr       string
	LogLevel       string
	DatabaseURL    string
	RedisAddr      string
	RedisPassword  string
	APIAccessToken string
}

func Load() (Config, error) {
	// Support running from either the repository root or backend directory.
	_ = godotenv.Load(".env", "../.env")

	cfg := Config{
		Environment:    env("APP_ENV", "development"),
		HTTPAddr:       env("HTTP_ADDR", ":8080"),
		LogLevel:       env("LOG_LEVEL", "info"),
		DatabaseURL:    env("DATABASE_URL", ""),
		RedisAddr:      env("REDIS_ADDR", "localhost:6381"),
		RedisPassword:  env("REDIS_PASSWORD", ""),
		APIAccessToken: env("API_ACCESS_TOKEN", ""),
	}

	switch strings.ToLower(cfg.Environment) {
	case "development", "test", "production":
	default:
		return Config{}, fmt.Errorf("APP_ENV must be development, test, or production")
	}
	if strings.TrimSpace(cfg.HTTPAddr) == "" {
		return Config{}, fmt.Errorf("HTTP_ADDR must not be empty")
	}
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}
	if cfg.APIAccessToken == "" || len(cfg.APIAccessToken) < 32 || strings.HasPrefix(cfg.APIAccessToken, "replace_with_") {
		return Config{}, fmt.Errorf("API_ACCESS_TOKEN must contain at least 32 characters")
	}
	return cfg, nil
}

func env(key, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}
	return value
}
