package openapi

import (
	"ai-artist/chatbot/setting"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Chatbot(userPrompt string) (string, int, int, int) {
	// .env load
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// .env -> API Key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is not set in the .env file")
	}

	// OpenAI API endpoint
	apiURL := "https://api.openai.com/v1/chat/completions"

	// Request
	requestBody := RequestBody{
		Model: setting.Setting.Model,
		Messages: []Message{
			{
				Role:    "user",
				Content: userPrompt,
			},
		},
		MaxTokens:   setting.Setting.MaxToken,
		Temperature: 0.7,
	}

	// JSON encoding
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error encoding JSON: %v", err)
	}

	// Create HTTP POST request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	// Header set
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client
	client := &http.Client{}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to OpenAI: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// check HTTP status code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Non-OK HTTP status: %s\nResponse body: %s", resp.Status, string(body))
	}

	// JSON parse
	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		log.Fatalf("Error parsing JSON response: %v", err)
	}

	// response logging
	if len(responseBody.Choices) > 0 {
		fmt.Println("Assistant:", responseBody.Choices[0].Message.Content)
	} else {
		fmt.Println("No response from assistant.")
	}

	// token usage logging
	log.Printf("Prompt Tokens: %d, Completion Tokens: %d, Total Tokens: %d\n",
		responseBody.Usage.PromptTokens,
		responseBody.Usage.CompletionTokens,
		responseBody.Usage.TotalTokens,
	)

	return responseBody.Choices[0].Message.Content, responseBody.Usage.PromptTokens, responseBody.Usage.CompletionTokens, responseBody.Usage.TotalTokens
}
