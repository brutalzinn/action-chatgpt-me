package commit

import (
	"log"
	"os/exec"
)

func Fetch() {
	_, err := exec.Command("git", "fetch").Output()
	if err != nil {
		log.Fatal(err)
	}
}
func Add() {
	_, err := exec.Command("git", "add", ".").Output()
	if err != nil {
		log.Fatal(err)
	}
}
func Commit(message string) {
	_, err := exec.Command("git", "commit", "-m", message).Output()
	if err != nil {
		log.Fatal(err)
	}
}
func Push() {
	_, err := exec.Command("git", "push", "origin", "main").Output()
	if err != nil {
		log.Fatal(err)
	}
}
