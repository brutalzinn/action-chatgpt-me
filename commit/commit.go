package commit

import (
	"log"
	"os/exec"
)

func Add() {
	_, err := exec.Command("bash", "-c", "git add .").Output()
	if err != nil {
		log.Fatal(err)
	}
}
func Commit(message string) {
	_, err := exec.Command("bash", "-c", "git commit -m 'commit readme.md change'").Output()
	if err != nil {
		log.Fatal(err)
	}
}
func Push() {
	_, err := exec.Command("bash", "-c", "git push").Output()
	if err != nil {
		log.Fatal(err)
	}
}
