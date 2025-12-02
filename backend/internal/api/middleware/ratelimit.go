package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/online-compiler/backend/internal/models"
	"github.com/online-compiler/backend/internal/services"
)

// RateLimitMiddleware implements rate limiting
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		allowed, err := services.CheckRateLimit(ip)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Success: false,
				Error:   "Rate limit check failed",
				Code:    "INTERNAL_ERROR",
			})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusTooManyRequests, models.ErrorResponse{
				Success: false,
				Error:   "Rate limit exceeded. Please try again later.",
				Code:    "RATE_LIMIT_EXCEEDED",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
