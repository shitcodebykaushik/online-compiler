package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/online-compiler/backend/internal/models"
)

// ExecuteCodeMock handles code execution with mock results (for demo when Judge0 is unavailable)
func ExecuteCodeMock(c *gin.Context) {
	var req models.ExecuteRequest

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

	// Mock execution based on language
	var output string

	switch req.LanguageID {
	case 71: // Python
		if strings.Contains(req.Code, "print") {
			// Extract what's being printed (simple mock)
			output = "Hello, World!\n\n[DEMO MODE - Judge0 Unavailable]\nThis is a simulated output.\n"
		} else {
			output = "[DEMO MODE] Python code executed successfully\n"
		}
	case 50, 54: // C, C++
		output = "Hello World!\n[DEMO MODE] C/C++ code compiled and executed\n"
	case 63: // JavaScript
		output = "Hello\n[DEMO MODE] JavaScript code executed with Node.js\n"
	case 62: // Java
		output = "Hello, World!\n[DEMO MODE] Java code compiled and executed\n"
	case 60: // Go
		output = "Hello, World!\n[DEMO MODE] Go code compiled and executed\n"
	case 73: // Rust
		output = "Hello, World!\n[DEMO MODE] Rust code compiled and executed\n"
	case 68: // PHP
		output = "Hello, World!\n[DEMO MODE] PHP code executed\n"
	default:
		output = "[DEMO MODE] Code executed successfully\n"
	}

	// Simulate execution response
	response := &models.ExecuteResponse{
		Success:       true,
		Output:        output,
		Error:         "",
		ExecutionTime: 42.5,
		MemoryKB:      256,
		Status:        "Accepted (Demo Mode)",
	}

	c.JSON(http.StatusOK, response)
}
