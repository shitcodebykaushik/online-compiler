package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/online-compiler/backend/configs"
	"github.com/online-compiler/backend/internal/models"
)

// PistonService handles Piston API interactions
type PistonService struct {
	BaseURL string
	Client  *http.Client
}

// NewPistonService creates a new Piston service
func NewPistonService() *PistonService {
	return &PistonService{
		BaseURL: configs.AppConfig.Judge0URL,
		Client: &http.Client{
			Timeout: time.Duration(configs.AppConfig.Judge0Timeout) * time.Second,
		},
	}
}

// PistonRequest represents a Piston execution request
type PistonRequest struct {
	Language string   `json:"language"`
	Version  string   `json:"version"`
	Files    []File   `json:"files"`
	Stdin    string   `json:"stdin,omitempty"`
	Args     []string `json:"args,omitempty"`
}

// File represents a code file
type File struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content"`
}

// PistonResponse represents a Piston execution response
type PistonResponse struct {
	Language string `json:"language"`
	Version  string `json:"version"`
	Run      struct {
		Stdout string `json:"stdout"`
		Stderr string `json:"stderr"`
		Code   int    `json:"code"`
		Signal *int   `json:"signal"`
		Output string `json:"output"`
	} `json:"run"`
	Compile *struct {
		Stdout string `json:"stdout"`
		Stderr string `json:"stderr"`
		Code   int    `json:"code"`
		Output string `json:"output"`
	} `json:"compile,omitempty"`
}

// Language mapping from Judge0 IDs to Piston languages
var languageMap = map[int]struct {
	Language string
	Version  string
	FileName string
}{
	71: {"python", "3.10.0", "main.py"},      // Python 3
	63: {"javascript", "18.15.0", "main.js"}, // JavaScript (Node.js)
	62: {"java", "15.0.2", "Main.java"},      // Java
	54: {"c++", "10.2.0", "main.cpp"},        // C++17
	50: {"c", "10.2.0", "main.c"},            // C
	51: {"c#", "6.12.0", "main.cs"},          // C#
	60: {"go", "1.16.2", "main.go"},          // Go
	68: {"php", "8.2.3", "main.php"},         // PHP
	72: {"ruby", "3.0.1", "main.rb"},         // Ruby
	73: {"rust", "1.68.2", "main.rs"},        // Rust
	78: {"kotlin", "1.8.20", "main.kt"},      // Kotlin
	80: {"r", "4.1.1", "main.r"},             // R
	82: {"sql", "3.36.0", "main.sql"},        // SQL (SQLite)
	83: {"swift", "5.3.3", "main.swift"},     // Swift
	74: {"typescript", "5.0.3", "main.ts"},   // TypeScript
}

// ExecuteCode executes code using Piston
func (p *PistonService) ExecuteCode(languageID int, code, stdin string) (*models.ExecuteResponse, error) {
	// Get language info
	langInfo, exists := languageMap[languageID]
	if !exists {
		return &models.ExecuteResponse{
			Success: false,
			Error:   fmt.Sprintf("Language ID %d not supported", languageID),
		}, nil
	}

	// Create request
	pistonReq := PistonRequest{
		Language: langInfo.Language,
		Version:  langInfo.Version,
		Files: []File{
			{
				Name:    langInfo.FileName,
				Content: code,
			},
		},
		Stdin: stdin,
	}

	jsonData, err := json.Marshal(pistonReq)
	if err != nil {
		return &models.ExecuteResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Execute code
	url := fmt.Sprintf("%s/api/v2/execute", p.BaseURL)
	resp, err := p.Client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return &models.ExecuteResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to connect to Piston: %v", err),
		}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return &models.ExecuteResponse{
			Success: false,
			Error:   fmt.Sprintf("Piston error: %s", string(body)),
		}, nil
	}

	var pistonResp PistonResponse
	if err := json.NewDecoder(resp.Body).Decode(&pistonResp); err != nil {
		return &models.ExecuteResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Build response
	response := &models.ExecuteResponse{
		Success: true,
		Status:  "Completed",
	}

	// Check for compilation errors
	if pistonResp.Compile != nil && pistonResp.Compile.Code != 0 {
		response.Error = pistonResp.Compile.Stderr
		if response.Error == "" {
			response.Error = pistonResp.Compile.Output
		}
		return response, nil
	}

	// Set output
	if pistonResp.Run.Stdout != "" {
		response.Output = pistonResp.Run.Stdout
	}

	// Set errors
	if pistonResp.Run.Stderr != "" {
		response.Error = pistonResp.Run.Stderr
	}

	// If exit code is non-zero and no stderr, use output
	if pistonResp.Run.Code != 0 && response.Error == "" {
		response.Error = fmt.Sprintf("Process exited with code %d", pistonResp.Run.Code)
	}

	return response, nil
}
