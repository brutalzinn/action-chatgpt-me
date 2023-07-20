package commit

import (
	"fmt"
	"log"
	"os/exec"
)

func Add() {
	_, err := exec.Command("git", "add", ".").Output()
	if err != nil {
		log.Fatal(err)
	}
}
func Commit(message string) {
	commitMessage := fmt.Sprintf("-m '%s'", message)
	_, err := exec.Command("git", "commit", commitMessage).Output()
	if err != nil {
		log.Fatal(err)
	}
}
func Push() {
	_, err := exec.Command("git", "push").Output()
	if err != nil {
		log.Fatal(err)
	}
}
