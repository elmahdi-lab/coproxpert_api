package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

const (
	Production  = "PRODUCTION"
	Staging     = "STAGING"
	Development = "DEVELOPMENT"
	Testing     = "TESTING"
)

type ServerConfig struct {
	Host     string
	Port     string
	Timezone *time.Location
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type AuthConfig struct {
	JWTSecret          string
	RefreshTokenExpiry time.Duration
	AccessTokenExpiry  time.Duration
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		// It's okay if .env doesn't exist in production
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	timezone, err := time.LoadLocation(getEnv("TIMEZONE", "UTC"))
	if err != nil {
		return nil, err
	}

	return &Config{
		Server: ServerConfig{
			Host:     getEnv("HOST", "localhost"),
			Port:     getEnv("PORT", "8080"),
			Timezone: timezone,
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "coproxpert"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Auth: AuthConfig{
			JWTSecret:          getEnv("JWT_SECRET", "your-secret-key"),
			RefreshTokenExpiry: 30 * 24 * time.Hour, // 30 days
			AccessTokenExpiry:  24 * time.Hour,      // 1 day
		},
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
