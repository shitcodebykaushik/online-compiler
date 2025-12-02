package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/online-compiler/backend/internal/database"
	"github.com/online-compiler/backend/internal/models"
	"github.com/online-compiler/backend/internal/services"
)

// HealthCheck handles health check requests
func HealthCheck(c *gin.Context) {
	response := models.HealthResponse{
		Status:   "healthy",
		Redis:    "disconnected",
		Database: "disconnected",
		Judge0:   "unknown",
	}

	// Check Redis
	if services.RedisClient != nil {
		if _, err := services.RedisClient.Ping(c).Result(); err == nil {
			response.Redis = "connected"
		} else if err == redis.Nil {
			response.Redis = "connected"
		}
	}

	// Check Database
	if database.DB != nil {
		if sqlDB, err := database.DB.DB(); err == nil {
			if err := sqlDB.Ping(); err == nil {
				response.Database = "connected"
			}
		}
	}

	// Check Judge0
	j := services.NewJudge0Service()
	if _, err := j.SubmitCode(71, "print('health')", ""); err == nil {
		response.Judge0 = "available"
	} else {
		response.Judge0 = "unavailable"
	}

	c.JSON(http.StatusOK, response)
}
