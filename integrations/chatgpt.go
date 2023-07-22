package integrations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	openAIURL = "https://api.openai.com/v1/chat/completions"
)

type ChatGPTRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

// chat gpt generate itsself integation. Soo easy when you can document you own model and share with a gooooood documentation.
func GetChatGPTResponse(prompt string) (string, error) {
	apiKey := os.Getenv("CHAT_GPT_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OpenAI API key not found. Set OPENAI_API_KEY environment variable")
	}
	client := &http.Client{}

	requestData := ChatGPTRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: "You are a head hunter and needs a developer that can solve you problems.",
			},
			{
				Role:    "user",
				Content: prompt,
			},
		},
		MaxTokens: 100,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var chatGPTResponse ChatGPTResponse
	err = json.NewDecoder(resp.Body).Decode(&chatGPTResponse)
	if err != nil {
		return "", err
	}

	if len(chatGPTResponse.Choices) > 0 {
		return chatGPTResponse.Choices[0].Message.Content, nil
	}

	return "", nil
}
