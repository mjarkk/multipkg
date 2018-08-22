package run

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/ionrock/procs"
	"github.com/kr/pty"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Cleanup cleans the output of the live output
func Cleanup(App *types.App, input string) string {
	return App.Replace(input, "[m", `((\\|\[)([0-9a-z]|;){1,}?[a-z])`)
}

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
func Interactive(App *types.App, command string, OutputHandler func(line string) string) (err error) {
	commands := strings.Split(command, " ")
	if len(commands) < 2 {
		return errors.New("command to short")
	}

	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Env = procs.Env(map[string]string{"LANG": "en_US.utf8"}, true)

	tty, err := pty.Start(cmd)
	pty.Setsize(tty, &pty.Winsize{
		Rows: 200,
		Cols: 200,
	})

	if err != nil {
		return err
	}

	defer tty.Close()
	go func() {
		scanner := bufio.NewScanner(tty)
		for scanner.Scan() {
			OutputHandler(Cleanup(App, scanner.Text()))
		}
	}()
	go func() {
		io.Copy(tty, os.Stdin)
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
