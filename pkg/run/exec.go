package run

import (
	"bufio"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

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

func write(term *os.File, toType string) {
	term.Write([]byte(toType + "\n"))
}

// Interactive runs a command with a function bind input and output
func Interactive(App *types.App, command string, OutputHandler func(line string, tty *os.File, scanner *bufio.Scanner) string) (err error) {

	cmd := exec.Command("bash")
	cmd.Env = procs.Env(map[string]string{"LANG": "en_US.utf8"}, true)

	tty, err := pty.Start(cmd)
	pty.Setsize(tty, &pty.Winsize{
		Rows: 200,
		Cols: 200,
	})

	if err != nil {
		return err
	}

	defer func() { _ = tty.Close() }()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, tty); err != nil {
				log.Printf("error resizing pty: %s", err)
			}
		}
	}()
	ch <- syscall.SIGWINCH

	go func() {
		scanner := bufio.NewScanner(tty)
		for scanner.Scan() {
			toTypeNext := OutputHandler(Cleanup(App, scanner.Text()), tty, scanner)
			if len(toTypeNext) > 0 {
				go write(tty, toTypeNext)
			}
		}
	}()

	go func() {
		write(tty, command+" && exit")
	}()

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}
