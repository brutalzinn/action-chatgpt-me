package commit

import (
	"fmt"
	"os/exec"
)

func SetConfig() error {
	_, err := exec.Command("bash", "-c", "git", "config", "--global user.name 'chatgptme'").Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("bash", "-c", "git", "config", "--global user.email '<>'").Output()
	return err
}

func Add() error {
	_, err := exec.Command("bash", "-c", "git add -A").Output()
	return err
}
func Commit(message string) error {
	commitMessage := fmt.Sprintf("-m '%s'", message)
	_, err := exec.Command("bash", "-c", "git", "commit", commitMessage).Output()
	return err
}
func Push() error {
	_, err := exec.Command("bash", "-c", "git", "push").Output()
	return err
}
