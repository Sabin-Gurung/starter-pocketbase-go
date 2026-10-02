package config

import (
	"os"
	"strings"
)

type Config struct {
	Port          string
	DevMode       bool
	AdminEmail    string
	AdminPassword string
}

func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	devMode := strings.ToLower(os.Getenv("DEV_MODE")) == "true" || os.Getenv("DEV_MODE") == "1"

	adminEmail := os.Getenv("ADMIN_EMAIL")
	if adminEmail == "" {
		adminEmail = "admin@example.com"
	}

	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		adminPassword = "supersecret"
	}

	return Config{
		Port:          port,
		DevMode:       devMode,
		AdminEmail:    adminEmail,
		AdminPassword: adminPassword,
	}
}
