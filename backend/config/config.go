package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	Port        string
	DbFile      string
	DbURL       string
	Schema      string
	JwtSecret   string
	JwtExpires  string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using default environment variables.")
	}

	return Config{
		Environment: os.Getenv("NODE_ENV"),
		Port:        os.Getenv("PORT"),
		DbFile:      os.Getenv("DB_FILE"),
		DbURL:       os.Getenv("DATABASE_URL"),
		Schema:      os.Getenv("SCHEMA"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
		JwtExpires:  os.Getenv("JWT_EXPIRES_IN"),
	}
}
