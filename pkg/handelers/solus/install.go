package solus

import (
	"bufio"
	"os"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Install installes a program
func Install(pkg string, flags *types.Flags) error {
	PKG = App.Replace(pkg, "", `^\s+|\s+$`)
	if len(PKG) == 0 {
		gui.FriendlyErr("No package(s) specified to remove")
	}
	pkgInfo, err := GetInfo(pkg, flags)
	if err != nil {
		gui.FriendlyErr(err)
	}
	if pkgInfo.Installed {
		gui.Echo(false, "The following package(s) are already installed and wil be not reinstalled")
		output := App.FindAllMatch(pkg, `((\w|\d|\-|\_|\+)+)`, 1)
		gui.ShowList(output, "dashList")
		os.Exit(0)
	}
	if gui.InstallQuestion(PKG, pkgInfo.InstallSize, true) {
		run.Interactive(App, "eopkg --no-color --yes-all install "+pkg, installOutputHandeler)
	}
	return nil
}

// installOutputHandeler handels the default line output of run.Interactive
func installOutputHandeler(line string, tty *os.File, scanner *bufio.Scanner) string {

	commandOutput = append(commandOutput, line)
	needRootErr(line, nil)

	// output the terminal output for the --debug flagg)
	gui.Echof(true, "cmdOut: %q\n", line)

	// Run functions from last line
	toExecuteAtEndOfNextLineWrapper(line)

	// pre definded regular expresion(s)
	DownloadingOrInstalling := `(?i)((downloading)|(installing))\s+(\d+).+(\d+)`
	DownloadingOrInstallLineMatch := `(((\w|-)+)-((\d|\.|\-|\w|\_)+)(\s\[((\w|\d)+)\])?\w*)$`

	if App.NormalMatch(`(?i)Warning: The following package(\(s\)|s)? are already installed and are not going to be installed again`, line) {
		// Check for packages that are already installed error
		gui.Echo(false, "The following package(s) are already installed and wil be not reinstalled")
		nextExecFuncMatchRegx = ``
		toExecuteAtEndOfNextLine = func(line string, extraData types.Flags) {
			output := App.FindAllMatch(line, `((\w|\d|\-|\_|\+)+)`, 1)
			gui.ShowList(output, "dashList")
		}
	} else if App.NormalMatch(DownloadingOrInstalling, line) {
		// Show Download/Install progress
		lastExecLineData = types.Flags{
			"index": App.FindMatch(line, DownloadingOrInstalling, 4),
			"of":    App.FindMatch(line, DownloadingOrInstalling, 5),
			"type":  App.FindMatch(line, DownloadingOrInstalling, 1),
			"regx":  DownloadingOrInstallLineMatch,
		}

		nextExecFuncMatchRegx = DownloadingOrInstallLineMatch

		toExecuteAtEndOfNextLine = func(line string, extraData types.Flags) {
			pkg := App.FindMatch(line, extraData["regx"].(string), 2)
			gui.ProgressIn(extraData["index"].(string), extraData["of"].(string), extraData["type"].(string), pkg)
		}
	} else if App.NormalMatch(`(Installed\s((\w|\s|\-|\_)+)\s*)$`, line) {
		gui.Installed(PKG)
	}

	return ""
}
