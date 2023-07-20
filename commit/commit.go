package commit

import "os/exec"

func Commit(message string) {
	exec.Command("git add .")
	exec.Command("git commit -m '"+ message+"'")
	exec.Command("git push")
}


