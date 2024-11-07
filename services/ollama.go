package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetStudentSummaryFromOllama generates a student summary using Ollama API
func GetStudentSummaryFromOllama(prompt string) (string, error) {
	ollamaAPI := "http://localhost:11434/api/generate" // Ollama API endpoint

	requestBody := map[string]interface{}{
		"model":  "llama3.2",
		"prompt": prompt,
		"stream": false,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Create the HTTP request
	client := &http.Client{}
	req, err := http.NewRequest("POST", ollamaAPI, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Send the request to Ollama API
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request to Ollama failed: %v", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code from Ollama: %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal Ollama response: %v", err)
	}

	// Extract the summary (assuming it's in the 'text' field)
	summary, ok := response["response"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format from Ollama: %v", response)
	}

	// Return the summary
	return summary, nil
}
