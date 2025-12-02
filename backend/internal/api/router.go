package api

import (
	"github.com/gin-gonic/gin"
	"github.com/online-compiler/backend/internal/api/handlers"
	"github.com/online-compiler/backend/internal/api/middleware"
)

// SetupRouter configures all routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Apply middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", handlers.HealthCheck)

		// Code execution (with rate limiting) - Using Piston
		v1.POST("/execute", middleware.RateLimitMiddleware(), handlers.ExecuteCodePiston)

		// Snippet management
		v1.POST("/snippets", handlers.CreateSnippet)
		v1.GET("/snippets/:id", handlers.GetSnippet)
	}

	return router
}
