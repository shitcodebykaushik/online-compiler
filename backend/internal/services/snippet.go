package services

import (
	"github.com/google/uuid"
	"github.com/online-compiler/backend/internal/database"
	"github.com/online-compiler/backend/internal/models"
)

// CreateSnippet creates a new code snippet
func CreateSnippet(req *models.SnippetRequest) (*models.Snippet, error) {
	snippet := &models.Snippet{
		ID:       uuid.New().String(),
		Language: req.Language,
		Code:     req.Code,
		Title:    req.Title,
		Views:    0,
	}

	if err := database.DB.Create(snippet).Error; err != nil {
		return nil, err
	}

	return snippet, nil
}

// GetSnippet retrieves a snippet by ID
func GetSnippet(id string) (*models.Snippet, error) {
	var snippet models.Snippet

	if err := database.DB.First(&snippet, "id = ?", id).Error; err != nil {
		return nil, err
	}

	// Increment view count
	database.DB.Model(&snippet).Update("views", snippet.Views+1)

	return &snippet, nil
}
