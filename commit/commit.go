package commit

import (
	"fmt"
	"os/exec"
)

type GitConfig struct {
	Username string
	Email    string
}

func SetConfig(gitCfg GitConfig) error {

	_, err := exec.Command("bash", "-c", "git config --global user.name '"+gitCfg.Username+"'").Output()
	if err != nil {
		return err
	}
	_, err = exec.Command("bash", "-c", "git config --global user.email '"+gitCfg.Email+"'").Output()
	return err
}

func Add() error {
	_, err := exec.Command("bash", "-c", "git add .").Output()
	return err
}
func Commit(message string) error {
	commitMessage := fmt.Sprintf("-m '%s'", message)
	_, err := exec.Command("bash", "-c", "git commit", commitMessage).Output()
	return err
}
func Push() error {
	_, err := exec.Command("bash", "-c", "git push origin main").Output()
	return err
}
