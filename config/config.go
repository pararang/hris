package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DB struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	SSLMode  string
}

type JWT struct {
	Secret        string
	ExpiryMinutes int
}

type Config struct {
	Env  string
	Port int
	DB   DB
	JWT  JWT
}

func New() Config {
	// load .env file from root directory
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Env:  getEnvString("ENV", "development"),
		Port: getEnvInt("PORT", 8080),
		DB: DB{
			Name:     getEnvString("DB_NAME", "postgres"),
			Host:     getEnvString("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnvString("DB_USER", "postgres"),
			Password: getEnvString("DB_PASSWORD", "password"),
			SSLMode:  getEnvString("DB_SSL_MODE", "disable"),
		},
		JWT: JWT{
			Secret:        getEnvString("JWT_SECRET", "secret"),
			ExpiryMinutes: getEnvInt("JWT_EXPIRY_MINUTES", 60),
		},
	}
}

func getEnvString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
