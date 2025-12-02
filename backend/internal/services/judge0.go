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

// Judge0Service handles Judge0 API interactions
type Judge0Service struct {
	BaseURL string
	Client  *http.Client
}

// NewJudge0Service creates a new Judge0 service
func NewJudge0Service() *Judge0Service {
	return &Judge0Service{
		BaseURL: configs.AppConfig.Judge0URL,
		Client: &http.Client{
			Timeout: time.Duration(configs.AppConfig.Judge0Timeout) * time.Second,
		},
	}
}

// SubmitCode submits code to Judge0 for execution
func (j *Judge0Service) SubmitCode(languageID int, code, stdin string) (string, error) {
	submission := models.Judge0Submission{
		SourceCode: code,
		LanguageID: languageID,
		Stdin:      stdin,
	}

	jsonData, err := json.Marshal(submission)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/submissions?base64_encoded=false&wait=false", j.BaseURL)
	fmt.Printf("DEBUG: Submitting to URL: %s\n", url)
	resp, err := j.Client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to submit to Judge0: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Judge0 submission failed: %s", string(body))
	}

	var result models.Judge0Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Token, nil
}

// GetSubmissionResult polls Judge0 for submission result
func (j *Judge0Service) GetSubmissionResult(token string) (*models.Judge0Result, error) {
	maxPolls := 10
	pollInterval := time.Second

	for i := 0; i < maxPolls; i++ {
		url := fmt.Sprintf("%s/submissions/%s?base64_encoded=false", j.BaseURL, token)
		resp, err := j.Client.Get(url)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return nil, fmt.Errorf("failed to get submission result: status %d", resp.StatusCode)
		}

		var result models.Judge0Result
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, err
		}
		resp.Body.Close()

		// Status ID: 1=In Queue, 2=Processing
		if result.Status.ID > 2 {
			return &result, nil
		}

		time.Sleep(pollInterval)
	}

	return nil, fmt.Errorf("execution timeout: max polls reached")
}

// ExecuteCode submits code and waits for result
func (j *Judge0Service) ExecuteCode(languageID int, code, stdin string) (*models.ExecuteResponse, error) {
	// Submit code
	token, err := j.SubmitCode(languageID, code, stdin)
	if err != nil {
		return &models.ExecuteResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Get result
	result, err := j.GetSubmissionResult(token)
	if err != nil {
		return &models.ExecuteResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	// Parse response
	response := &models.ExecuteResponse{
		Success: true,
		Status:  result.Status.Description,
	}

	if result.Stdout != nil {
		response.Output = *result.Stdout
	}

	if result.Stderr != nil && *result.Stderr != "" {
		response.Error = *result.Stderr
	}

	if result.CompileOutput != nil && *result.CompileOutput != "" {
		response.Error = *result.CompileOutput
	}

	if result.Message != nil && *result.Message != "" {
		if response.Error == "" {
			response.Error = *result.Message
		}
	}

	if result.Time != nil {
		var execTime float64
		fmt.Sscanf(*result.Time, "%f", &execTime)
		response.ExecutionTime = execTime * 1000 // Convert to ms
	}

	if result.Memory != nil {
		response.MemoryKB = *result.Memory
	}

	// If status is not Accepted, mark as unsuccessful for errors
	if result.Status.ID != 3 && response.Error == "" {
		response.Error = result.Status.Description
	}

	return response, nil
}
