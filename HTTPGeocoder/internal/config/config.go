package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenCageKey string
	Port        string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Note: .env file not found, reading from environment variables")
	}

	cfg := &Config{
		OpenCageKey: os.Getenv("OPENCAGE_API_KEY"),
		Port:        os.Getenv("PORT"),
	}

	if cfg.OpenCageKey == "" {
		log.Fatal("Error: OPENCAGE_API_KEY not found in environment")
	}

	// Встановлюємо порт за замовчуванням, якщо він не вказаний
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	return cfg
}
