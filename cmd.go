package util

import (
	"os/exec"
)

func RunCmd(cmd string, shell bool) ([]byte, error) {
	if shell {
		return exec.Command("bash", "-c", cmd).Output()
	}
	return exec.Command(cmd).Output()
}
