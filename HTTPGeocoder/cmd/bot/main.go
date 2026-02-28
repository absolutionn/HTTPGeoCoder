package main

import (
	"log"
	"strconv"

	"GeocoderTgBot/internal/app"
	"GeocoderTgBot/internal/config"
	"GeocoderTgBot/internal/geocoder"
)

func main() {
	// 1. Завантаження конфігурації
	cfg := config.Load()

	// Конвертація порту в int
	portInt, err := strconv.Atoi(cfg.Port)
	if err != nil {
		log.Fatal("Invalid port:", err)
	}

	// 2. Ініціалізація сервісів
	geoClient := geocoder.NewClient(cfg.OpenCageKey)

	// 3. Створення додатку (Dependency Injection)
	application := app.New(geoClient, portInt)

	log.Printf("HTTP Server is running on port %d...", portInt)

	// 4. Запуск сервера
	if err := application.Run(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
