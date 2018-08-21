package run

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/ionrock/procs"
)

// Run a simple version of exec.Command + cmd.CombinedOutput
func Run(command string) (output string, err error) {
	commands := strings.Split(command, " ")
	if len(commands) < 2 {
		return "", errors.New("command to short")
	}

	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Env = procs.Env(map[string]string{"LANG": "en_US.utf8"}, true)

	bytes, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// Interactive runs a command with a function bind input and output
func Interactive(command string, OutputHandler func(line string) string, ErrHandler func(line string) string) (err error) {
	p := procs.NewProcess(command)
	p.Env = map[string]string{"LANG": "en_US.utf8"}
	p.OutputHandler = OutputHandler
	p.ErrHandler = ErrHandler
	return p.Run()
}
