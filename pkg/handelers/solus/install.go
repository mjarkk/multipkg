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
	gui.NoPkgsInstall(PKG)

	PKGs := App.FindAllMatch(PKG, `((\w|\d|\-|\_|\+)+)`, 1)
	Installed := ""
	for _, PKGsItem := range PKGs {
		toAdd, err := GetInfo(PKGsItem, flags)
		if err != nil {
			gui.FriendlyErr(err)
		}
		if toAdd.Installed {
			if len(Installed) > 0 {
				Installed = Installed + " " + PKGsItem
			} else {
				Installed = PKGsItem
			}
		}
	}

	if len(Installed) > 0 {
		gui.Echo(false, "The following package(s) are already installed and wil be not installed")
		output := App.FindAllMatch(Installed, `((\w|\d|\-|\_|\+)+)`, 1)
		gui.ShowList(output, "dashList")
		os.Exit(0)
	}

	out, err := run.Run("eopkg --dry-run --no-color install " + PKG)
	needRootErr(out, err)
	installSize := App.FindMatch(out, `Total size of package\(s\):\s*((\d|\w|\.)+(\s(\w|\d)+)?)`, 1)

	if gui.InstallQuestion(PKG, installSize, true) {
		run.Interactive(App, "eopkg --no-color --yes-all install "+PKG, installOutputHandeler)
	}

	return nil
}

// installOutputHandeler handels the default line output of run.Interactive
func installOutputHandeler(line string, tty *os.File, scanner *bufio.Scanner) string {

	commandOutput = append(commandOutput, line)
	needRootErr(line, nil)

	// output the terminal output for the --debug flagg
	gui.Echo(true, "cmdOut:", line)

	// Run functions from last line
	toExecuteAtEndOfNextLineWrapper(line)

	// pre definded regular expresion(s)
	DownloadingOrInstalling := `(?i)((downloading)|(installing))\s+(\d+).+(\d+)`
	DownloadingOrInstallLineMatch := `(((\w|-)+)-((\d|\.|\-|\w|\_)+)((((\s\[((\w|\d)+)\])?\w*)|(\(\d.+))|(\s*\((\w|\d|\.)+(\s*\w*)?\)\s*\d*\%.+)))$`
	Installed := `(Installed\s((\w|\s|\-|\_)+)\s*)$`

	if App.NormalMatch(`Package (\w|\d|\-|\_|\.|\+)+ found in`, line) {
		return ""
	} else if App.NormalMatch(`(?i)Warning: The following package(\(s\)|s)? are already installed and are not going to be installed again`, line) {
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
	} else if App.NormalMatch(Installed, line) {
		gui.Installed(App.FindMatch(line, Installed, 2))
	}

	return ""
}
