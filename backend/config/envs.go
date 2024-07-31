package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost    string
	Port          string
	DBUser        string
	DBPassword    string
	DBAddress     string
	DBName        string
	Email         string
	EmailPassword string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:    getEnv("PUBLIC_HOST", "http://localhost"),
		Port:          getEnv("PORT", "8080"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "password"),
		DBAddress:     fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:        getEnv("DB_NAME", "mydatabase"),
		Email:         getEnv("EMAIL", "name@mail.com"),
		EmailPassword: getEnv("EMAIL_PASSWORD", "password"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
