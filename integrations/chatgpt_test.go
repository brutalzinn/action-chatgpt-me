package integrations

import (
	"os"
	"testing"
)

func TestIntegrationGetChatGPTResponse(t *testing.T) {
	// Get the OpenAI API key from an environment variable
	apiKey := os.Getenv("CHAT_GPT_API_KEY")
	if apiKey == "" {
		t.Fatal("OpenAI API key not found. Set OPENAI_API_KEY environment variable to run the test.")
	}
	prompt := "Say hello to me."
	response, err := GetChatGPTResponse(prompt)
	if err != nil {
		t.Errorf("Error getting GPT-3.5 response: %v", err)
	}
	if len(response) == 0 {
		t.Errorf("Empty response received from GPT-3.5 API.")
	}
}
