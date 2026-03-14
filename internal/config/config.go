package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenCageKey string
	Port        int
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, reading from environment variables")
	}

	cfg := &Config{
		OpenCageKey: os.Getenv("OPENCAGE_API_KEY"),
	}

	if cfg.OpenCageKey == "" {
		log.Fatal("Error: OPENCAGE_API_KEY not found in environment")
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		cfg.Port = 8080
	} else {
		portInt, err := strconv.Atoi(portStr)
		if err != nil {
			log.Fatal("Error: Invalid PORT environment variable", err)
		}
		cfg.Port = portInt
	}

	return cfg
}
