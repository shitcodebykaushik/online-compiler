package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/online-compiler/backend/internal/models"
	"github.com/online-compiler/backend/internal/services"
)

// ExecuteCodePiston handles code execution using Piston
func ExecuteCodePiston(c *gin.Context) {
	var req models.ExecuteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Piston service
	pistonService := services.NewPistonService()

	// Execute code
	result, err := pistonService.ExecuteCode(req.LanguageID, req.Code, req.Stdin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
