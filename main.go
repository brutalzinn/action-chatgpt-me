package main

import (
	"action-chatgpt-me/commit"
	"action-chatgpt-me/editor"
	"fmt"
	"os"
)

func main() {
	wakaTimeApiKey := os.Getenv("WAKATIME_API_KEY")
	chatGptKey := os.Getenv("CHAT_GPT_API_KEY")
	fmt.Printf("Now you will see. :) %s %s", wakaTimeApiKey, chatGptKey)
	filePath := "README.md"
	markdownTxt, err := editor.ReadMarkdown(filePath)
	if err != nil {
		fmt.Printf("Something does wrong..")
	}
	newTxt := editor.FindAndReplace(markdownTxt, "hello")
	editor.WriteMarkdown(filePath, newTxt)
	commit.Commit("change readme")
}
