package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	apiKey := os.Getenv("XAI_API_KEY")
	prompt := "give me a summary of the ffiec aio handbook" // Your prompt here

	// Prepare request body
	reqBody := map[string]interface{}{
		"model": "grok-beta",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}
	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		os.Exit(1)
	}

	// Print the request body for debugging
	fmt.Println("Request Body:", string(jsonBody))

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://api.x.ai/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	// Parse response
	var response struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshaling response:", err)
		os.Exit(1)
	}

	spew.Dump(response)

	// Print the answer
	if len(response.Choices) > 0 {
		fmt.Println(response.Choices[0].Message.Content)
	} else {
		fmt.Println("No response choices found.")
	}
}
