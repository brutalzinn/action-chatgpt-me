package main

import (
	"fmt"
	"os"
)

func main() {
	wakaTimeApiKey := os.Getenv("WAKATIME_API_KEY")
	chatGptKey := os.Getenv("CHAT_GPT_API_KEY")
	fmt.Println("No. you will not cant see any of my keys here. :) %s %s", wakaTimeApiKey, chatGptKey)
}
