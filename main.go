package main

import (
	"action-chatgpt-me/editor"
	"action-chatgpt-me/integrations"
	"fmt"
	"strings"
)

func main() {
	filePath := "README.md"
	originalMarkdown, err := editor.ReadMarkdown(filePath)
	if err != nil {
		fmt.Printf("Something does wrong..")
	}
	fmt.Printf("step find and place.")
	newMarkdown := GetMarkdownText()
	newTxt := editor.FindAndReplace(originalMarkdown, newMarkdown)
	editor.WriteMarkdown(filePath, newTxt)
	fmt.Printf("ok.")
}

func GetChatgptText() string {
	userStats, err := integrations.GetUserStats()
	var text strings.Builder
	if err != nil {
		text.WriteString("oh lord..")
		return ""
	}
	text.WriteString("my most used operating systems of the week:")
	for _, item := range userStats.Data.OperationalSystem[:3] {
		text.WriteString(item.Name + ",")
	}
	text.WriteString("My coding activities and time used of this week:")
	for _, item := range userStats.Data.Category[:3] {
		text.WriteString(item.Name + " " + item.Time + ",")
	}
	text.WriteString("languages used this week:")
	for _, item := range userStats.Data.Language {
		text.WriteString(item.Name + ",")
	}
	textToChatgpt := text.String()
	chatgptResponse, err := integrations.GetChatGPTResponse(textToChatgpt)
	if err != nil {
		text.WriteString("oh lord..")
		return ""
	}
	return chatgptResponse
}

func GetMarkdownText() string {
	chatgptText := GetChatgptText()
	var text strings.Builder
	text.WriteString(chatgptText)
	text.WriteString("did you like this readme? Take a look at <a href='https://github.com/brutalzinn/action-chatgpt-me'>Click here</a>")
	result := text.String()
	return result
}
