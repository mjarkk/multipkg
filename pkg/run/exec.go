package run

import (
	"errors"
	"os/exec"
	"strings"
)

// Run a simple version of exec.Command + cmd.CombinedOutput
func Run(command string) (output string, err error) {
	commands := strings.Split(command, " ")
	if len(commands) < 2 {
		return "", errors.New("command to short")
	}
	cmd := exec.Command(commands[0], commands[1:]...)
	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
