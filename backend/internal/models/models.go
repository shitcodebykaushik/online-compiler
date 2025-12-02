package models

import "time"

// ExecuteRequest represents a code execution request
type ExecuteRequest struct {
	LanguageID int    `json:"language_id" binding:"required"`
	Code       string `json:"code" binding:"required"`
	Stdin      string `json:"stdin"`
}

// ExecuteResponse represents a code execution response
type ExecuteResponse struct {
	Success       bool    `json:"success"`
	Output        string  `json:"output,omitempty"`
	Error         string  `json:"error,omitempty"`
	ExecutionTime float64 `json:"execution_time,omitempty"`
	MemoryKB      int     `json:"memory_kb,omitempty"`
	Status        string  `json:"status,omitempty"`
}

// Judge0Submission represents Judge0 submission request
type Judge0Submission struct {
	SourceCode string `json:"source_code"`
	LanguageID int    `json:"language_id"`
	Stdin      string `json:"stdin,omitempty"`
}

// Judge0Response represents Judge0 submission response
type Judge0Response struct {
	Token string `json:"token"`
}

// Judge0Result represents Judge0 result
type Judge0Result struct {
	Stdout        *string `json:"stdout"`
	Stderr        *string `json:"stderr"`
	CompileOutput *string `json:"compile_output"`
	Message       *string `json:"message"`
	Time          *string `json:"time"`
	Memory        *int    `json:"memory"`
	Status        Status  `json:"status"`
}

// Status represents execution status
type Status struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Snippet represents a code snippet
type Snippet struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Language  string    `gorm:"not null" json:"language"`
	Code      string    `gorm:"type:text;not null" json:"code"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Views     int       `gorm:"default:0" json:"views"`
}

// SnippetRequest represents a snippet creation request
type SnippetRequest struct {
	Language string `json:"language" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Title    string `json:"title"`
}

// SnippetResponse represents a snippet response
type SnippetResponse struct {
	Success   bool   `json:"success"`
	SnippetID string `json:"snippet_id,omitempty"`
	ShareURL  string `json:"share_url,omitempty"`
	Error     string `json:"error,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    string `json:"code,omitempty"`
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status   string `json:"status"`
	Redis    string `json:"redis"`
	Database string `json:"database"`
	Judge0   string `json:"judge0"`
}
