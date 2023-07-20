package commit

import "os/exec"

func Commit(message string) {
	exec.Command("git add -A")
	exec.Command("git commit -m \" commit test message \" ")
	exec.Command("git push")
}
