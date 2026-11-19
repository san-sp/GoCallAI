package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Environment string
	HTTPAddr    string
	LogLevel    string
}

func Load() (Config, error) {
	cfg := Config{
		Environment: env("APP_ENV", "development"),
		HTTPAddr:    env("HTTP_ADDR", ":8080"),
		LogLevel:    env("LOG_LEVEL", "info"),
	}

	switch strings.ToLower(cfg.Environment) {
	case "development", "test", "production":
	default:
		return Config{}, fmt.Errorf("APP_ENV must be development, test, or production")
	}

	if strings.TrimSpace(cfg.HTTPAddr) == "" {
		return Config{}, fmt.Errorf("HTTP_ADDR must not be empty")
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
