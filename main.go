package main

import (
	"action-chatgpt-me/commit"
	"action-chatgpt-me/editor"
	"fmt"
)

func main() {
	filePath := "README.md"
	markdownTxt, err := editor.ReadMarkdown(filePath)
	gitConfig := commit.GitConfig{
		Username: "chatgptme",
		Email:    "chatgptme@noreply.com",
	}
	commit.SetConfig(gitConfig)
	if err != nil {
		fmt.Printf("Something does wrong..")
	}
	newTxt := editor.FindAndReplace(markdownTxt, "done.")
	editor.WriteMarkdown(filePath, newTxt)
	commit.Add()
	commit.Commit("some update message here")
	commit.Push()
	fmt.Printf("ok.")
}
