package arch

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Update handels the update command
func Update(pkg string, flags *types.Flags) error {

	// TODO: remove later
	fmt.Println("update is not working yet...")

	run.Interactive(App, "pacman -Syuu", func(line string, tty *os.File, scanner *bufio.Scanner) string {
		return installUpdateProcess(line)
	})

	return nil
}

// Install handels the install command
func Install(pkg string, flags *types.Flags) error {

	// TODO: remove later
	fmt.Println("install is not working yet...")

	PKG = App.Replace(pkg, "", `^\s+|\s+$`)
	gui.NoPkgsInstall(PKG)

	run.Interactive(App, "pacman -Syuu "+PKG, func(line string, tty *os.File, scanner *bufio.Scanner) string {
		return installUpdateProcess(line)
	})

	return nil
}

func installUpdateProcess(line string) string {

	commandOutput = append(commandOutput, line)

	gui.Echo(true, "cmdOut:", line)

	return ""
}
