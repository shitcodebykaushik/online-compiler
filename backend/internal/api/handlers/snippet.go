package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/online-compiler/backend/internal/models"
	"github.com/online-compiler/backend/internal/services"
)

// CreateSnippet handles snippet creation
func CreateSnippet(c *gin.Context) {
	var req models.SnippetRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Invalid request format",
			Code:    "INVALID_INPUT",
		})
		return
	}

	// Validate code size
	if len(req.Code) > 65536 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Success: false,
			Error:   "Code exceeds maximum size of 64KB",
			Code:    "INVALID_INPUT",
		})
		return
	}

	snippet, err := services.CreateSnippet(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Success: false,
			Error:   "Failed to create snippet",
			Code:    "INTERNAL_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, models.SnippetResponse{
		Success:   true,
		SnippetID: snippet.ID,
		ShareURL:  "/snippets/" + snippet.ID,
	})
}

// GetSnippet handles snippet retrieval
func GetSnippet(c *gin.Context) {
	id := c.Param("id")

	snippet, err := services.GetSnippet(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Success: false,
			Error:   "Snippet not found",
			Code:    "NOT_FOUND",
		})
		return
	}

	c.JSON(http.StatusOK, snippet)
}
