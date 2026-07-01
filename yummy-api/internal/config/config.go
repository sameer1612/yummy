package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Env file not found!")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000"
	}

	if databaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	config := Config{
		DatabaseURL: databaseURL,
		Port:        port,
	}

	return config, nil
}
