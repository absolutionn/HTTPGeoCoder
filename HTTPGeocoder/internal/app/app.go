package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"GeocoderTgBot/internal/geocoder"
)

type App struct {
	router *gin.Engine
	geo    *geocoder.Client
	port   int
}

// New ініціалізує додаток та налаштовує маршрутизацію (Routing)
func New(geo *geocoder.Client, port int) *App {
	router := gin.Default() // Стандартний роутер Gin з логером та recovery

	a := &App{
		router: router,
		geo:    geo,
		port:   port,
	}

	a.setupRoutes()
	return a
}

func (a *App) setupRoutes() {
	// Створюємо групу маршрутів для API
	api := a.router.Group("/api")
	{
		// GET /api/geocode?q=Zhmerynka або GET /api/geocode?q=49.02,28.06
		api.GET("/geocode", a.handleGeocode)
	}
}

func (a *App) handleGeocode(c *gin.Context) {
	query := c.Query("q")

	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter 'q'"})
		return
	}

	// Запит до геокодера
	result, err := a.geo.Geocode(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error connecting to geocoding service"})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "NO RESULTS FOUND"})
		return
	}

	formattedAddress := *result.Formatted
	lat := *result.Geometry.Lat
	lng := *result.Geometry.Lng

	// Генерація правильного посилання на Google Maps
	googleMapsURL := fmt.Sprintf("https://maps.google.com/maps?q=%f,%f", lat, lng)

	// Повертаємо структуровану JSON-відповідь
	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"address": formattedAddress,
		"coordinates": gin.H{
			"lat": lat,
			"lng": lng,
		},
		"google_maps_url": googleMapsURL,
	})
}

// Run запускає HTTP сервер
func (a *App) Run() error {
	return a.router.Run(fmt.Sprintf(":%d", a.port))
}
