package main

import (
	"fmt"
	"os"
)

func main() {
	wakaTimeApiKey := os.Getenv("WAKATIME_API_KEY")
	chatGptKey := os.Getenv("CHAT_GPT_API_KEY")
	fmt.Printf("Now you will see. :) %s %s", wakaTimeApiKey, chatGptKey)
}
