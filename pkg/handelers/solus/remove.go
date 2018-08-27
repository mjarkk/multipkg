package solus

import (
	"bufio"
	"os"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Remove a program
func Remove(pkg string, flags *types.Flags) error {
	PKG = App.Replace(pkg, "", `^\s+|\s+$`)
	if len(PKG) == 0 {
		gui.FriendlyErr("No package(s) specified to remove")
	}
	pkgInfo, err := GetInfo(pkg, flags)
	if err != nil {
		gui.FriendlyErr(err)
	}
	if !pkgInfo.Installed {
		gui.Echo(false, "The following package(s) can't be removed because they are not intalled")
		output := App.FindAllMatch(pkg, `((\w|\d|\-|\_|\+)+)`, 1)
		gui.ShowList(output, "dashList")
		os.Exit(0)
	}
	if gui.RemoveQuestion(PKG, true) {
		run.Interactive(App, "eopkg --no-color --yes-all remove "+pkg, removeOutputHandeler)
	}
	return nil
}

// removeOutputHandeler handels the default line output of run.Interactive
func removeOutputHandeler(line string, tty *os.File, scanner *bufio.Scanner) string {
	commandOutput = append(commandOutput, line)
	needRootErr(line, nil)

	// output the terminal output for the --debug flagg)
	gui.Echof(true, "cmdOut: %q\n", line)

	// Run functions from last line
	toExecuteAtEndOfNextLineWrapper(line)

	// pre defined regexes
	removePackageDune := `(?i)removing\s*package\s*((\w|\d|\-|\_|\+)+).+Removed`
	removePackageNotDune := `((?i)removing\s*package\s*((\w|\d|\-|\_|\+)+).{0,5})$`

	// Removing package
	if App.NormalMatch(removePackageNotDune, line) {
		gui.Echo(false, "Removing", App.FindMatch(line, removePackageNotDune, 2), "...")
	} else if App.NormalMatch(removePackageDune, line) {
		gui.Echo(false, "Removed", App.FindMatch(line, removePackageDune, 1))
	}

	return ""
}
