package commit

import "os/exec"

func Add() {
	exec.Command("git add -A")
}
func Commit(message string) {
	exec.Command("git commit -m \" commit test message \" ")
}
func Push() {
	exec.Command("git push")
}
