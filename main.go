package main

import (
	"action-chatgpt-me/commit"
	"action-chatgpt-me/editor"
	"fmt"
)

func main() {
	filePath := "README.md"
	gitConfig := commit.GitConfig{
		Username: "chatgptme",
		Email:    "chatgptme@noreply.com",
	}
	commit.SetConfig(gitConfig)
	markdownTxt, err := editor.ReadMarkdown(filePath)
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
