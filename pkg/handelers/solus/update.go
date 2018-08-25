package solus

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Update a program
func Update(pkg string, flags *types.Flags) error {
	err := run.Interactive(App, "eopkg --no-color upgrade", updateOutputHandeler)
	fmt.Println(err)
	return nil
}

var updateCommandOutput = []string{}

// updateOutputHandeler handels the default line output of run.Interactive
func updateOutputHandeler(line string, tty *os.File, scanner *bufio.Scanner) string {
	updateCommandOutput = append(updateCommandOutput, line)
	needRootErr(line, nil)
	gui.Echo(true, "cmdOut:", line)
	return ""
}
