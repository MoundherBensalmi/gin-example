package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP      AppConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

type AppConfig struct {
	Mode string
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type JWTConfig struct {
	AccessKey     string
	RefreshKey    string
	AccessExpire  string
	RefreshExpire string
}

var Cfg Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	Cfg = Config{
		APP: AppConfig{
			Mode: getEnv("GIN_MODE", "debug"),
			Port: getEnv("APP_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			Username: getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "gin"),
		},
		JWT: JWTConfig{
			AccessKey:     getEnv("ACCESS_SECRET_KEY", "access"),
			RefreshKey:    getEnv("REFRESH_SECRET_KEY", "refresh"),
			AccessExpire:  getEnv("ACCESS_TOKEN_EXPIRE", "600"),
			RefreshExpire: getEnv("REFRESH_TOKEN_EXPIRE", "10080"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
