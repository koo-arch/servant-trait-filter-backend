package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if GetEnv("GO_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}