package handlers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/online-compiler/backend/internal/models"
	"github.com/online-compiler/backend/internal/services"
)

// ExecuteCode handles code execution requests
func ExecuteCode(c *gin.Context) {
	var req models.ExecuteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request format",
			Code:    "INVALID_INPUT",
		})
		return
	}

	// Validate code size (max 64KB)
	if len(req.Code) > 65536 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Code exceeds maximum size of 64KB",
			Code:    "INVALID_INPUT",
		})
		return
	}

	// Validate language ID (1-100 for Judge0)
	if req.LanguageID < 1 || req.LanguageID > 100 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid language ID",
			Code:    "INVALID_INPUT",
		})
		return
	}

	// Check cache
	codeHash := generateHash(fmt.Sprintf("%d:%s:%s", req.LanguageID, req.Code, req.Stdin))
	if cached, err := services.GetCachedResult(codeHash); err == nil {
		var response models.ExecuteResponse
		if json.Unmarshal(cached, &response) == nil {
			c.JSON(http.StatusOK, response)
			return
		}
	}

	// Create Judge0 service instance
	judge0Service := services.NewJudge0Service()

	// Execute code
	result, err := judge0Service.ExecuteCode(req.LanguageID, req.Code, req.Stdin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Code execution failed",
			Code:    "EXECUTION_ERROR",
		})
		return
	}

	// Cache result
	services.CacheResult(codeHash, result)

	c.JSON(http.StatusOK, result)
}

// generateHash creates a SHA256 hash for caching
func generateHash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", hash)
}
