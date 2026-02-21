package main

import (
	"log"

	"GeocoderTgBot/internal/app"
	"GeocoderTgBot/internal/config"
	"GeocoderTgBot/internal/geocoder"
)

func main() {
	// 1. Завантаження конфігурації
	cfg := config.Load()

	// 2. Ініціалізація сервісів
	geoClient := geocoder.NewClient(cfg.OpenCageKey)

	// 3. Створення додатку (Dependency Injection)
	application := app.New(geoClient, cfg.Port)

	log.Printf("HTTP Server is running on port %s...", cfg.Port)

	// 4. Запуск сервера
	if err := application.Run(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
