package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
	User string
	DBName string
	Password string
	SSLMode string
	JWT_SECRET string
}

var ENVS = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config {
		Host: getEnv("DB_HOST", "localhost"),
		Port: getEnv("DB_PORT", "5432"),
		User: getEnv("DB_USER", "postgres"),
		DBName: getEnv("DB_Name", "goecom"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		SSLMode: getEnv("DB_SSLMODE", "disable"),
		JWT_SECRET: getEnv("JWT_SECRET", "jwt_test"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

