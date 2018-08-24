package solus

import (
	"os"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Install installes a program
func Install(pkg string, flags *types.Flags) error {
	PKG = App.Replace(pkg, "", `^\s+|\s+$`)
	run.Interactive(App, "eopkg --no-color install "+pkg, installOutputHandeler)
	gui.Echo(false, "exiting..")
	return nil
}

// updateOutputHandeler handels the default line output of run.Interactive
func installOutputHandeler(line string, tty *os.File) string {

	commandOutput = append(commandOutput, line)
	needRootErr(line, nil)

	// output the terminal output for the --debug flagg)
	gui.Echo(true, "cmdOut:", line)

	// Check if the app need install confirmation
	if App.NormalMatch(`(?i)total size of package.{0,}:\s{0,2}`, line) {
		InstallSize := App.FindMatch(line, `:\s{0,2}([0-9]{0,3}(,[0-9]{0,3})?(\.[0-9]{0,4})?.+)`, 1)
		if gui.InstallQuestion(App, PKG, InstallSize) {
			return "yes"
		}
		return "no"
	}

	// Run functions from last line
	toExecuteAtEndOfNextLineWrapper(line)

	// Check for packages that are already installed error
	if App.NormalMatch(`(?i)Warning: The following package(\(s\)|s)? are already installed and are not going to be installed again`, line) {
		gui.Echo(false, "The following package(s) are already installed and wil be not reinstalled")
		toExecuteAtEndOfNextLine = func(line string) {
			output := App.FindAllMatch(line, `((\w|\d|-)+)`, 1)
			gui.ShowList(output, "dashList")
		}
	}

	return ""
}
