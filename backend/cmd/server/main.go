package main

import (
	"log"

	"github.com/online-compiler/backend/configs"
	"github.com/online-compiler/backend/internal/api"
	"github.com/online-compiler/backend/internal/database"
	"github.com/online-compiler/backend/internal/services"
)

func main() {
	// Load configuration
	configs.LoadConfig()
	log.Printf("Judge0 URL: %s", configs.AppConfig.Judge0URL)
	log.Println("Configuration loaded")

	// Initialize database
	if err := database.InitDatabase(configs.AppConfig.DatabasePath); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	log.Println("Database initialized")

	// Initialize Redis (optional)
	if err := services.InitRedis(); err != nil {
		log.Printf("Warning: Failed to initialize Redis: %v (continuing without Redis)", err)
	} else {
		log.Println("Redis initialized")
	}

	// Setup router
	router := api.SetupRouter()

	// Start server
	addr := ":" + configs.AppConfig.Port
	log.Printf("Starting server on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
