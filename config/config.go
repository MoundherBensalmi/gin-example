package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	APP struct {
		Mode string
		Port string
	}
	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		Name     string
	}
	JWT struct {
		Secret string
	}
}

var cfg Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	cfg = Config{
		// App configuration
		APP: struct {
			Mode string
			Port string
		}{
			Mode: getEnv("GIN_MODE", "debug"),
			Port: getEnv("APP_PORT", "8080"),
		},

		// Database configuration
		Database: struct {
			Host     string
			Port     string
			Username string
			Password string
			Name     string
		}{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Username: os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},

		// JWT configuration
		JWT: struct {
			Secret string
		}{
			Secret: os.Getenv("JWT_SECRET"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func Get() Config {
	return cfg
}
