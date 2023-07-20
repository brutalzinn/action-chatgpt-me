package main

import (
	"action-chatgpt-me/commit"
	"action-chatgpt-me/editor"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := strings.Join(os.Args[1:], " ")
	chatGptKey := args[0]
	wakaTimeApiKey := args[1]
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
