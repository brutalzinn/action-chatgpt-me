package commit

import "os/exec"

func Add() {
	cmd := exec.Command("git add .")
	cmd.Run()
}
func Commit(message string) {
	cmd := exec.Command("git commit -m \" commit test message \" ")
	cmd.Run()
}
func Push() {
	cmd := exec.Command("git push")
	cmd.Run()
}
