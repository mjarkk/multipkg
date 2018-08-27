package solus

import (
	"bufio"
	"errors"
	"os"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Update a program
func Update(pkg string, flags *types.Flags) error {
	command := "eopkg --no-color --dry-run upgrade"
	if !App.NormalMatch(`^(\s*)$`, pkg) {
		command = command + " " + pkg
	}
	out, err := run.Run(command)
	needRootErr(out, err)

	testRegx := `(?i)(The following packages will be upgraded:)\s*\n(((\w|\-|\_|\.|\d|\+)*(\s|\n|\t))*)(Total size).+:\s*((\d|\.)+\s*(\w+)?)`
	if !App.NormalMatch(testRegx, out) {
		gui.FriendlyErr(errors.New("There are no pacakges to update"))
	}

	toUpdate := App.FindMatch(out, testRegx, 2)
	toUpdate = App.Replace(toUpdate, "", `^(\s|\t|\n)+|(\s|\t|\n)+$`)
	toUpdateArr := App.FindAllMatch(toUpdate, `((\w|\d|\-|\_|\+|\.)+)`, 1)
	gui.Echo(false, "The following package(s) will be upgraded:")
	gui.ShowList(toUpdateArr, "normal")

	if gui.UpdateQuestion(App.FindMatch(out, testRegx, 7), true) {
		run.Interactive(App, "eopkg --no-color --yes-all upgrade", updateOutputHandeler)
	}

	return nil
}

var updateCommandOutput = []string{}

// updateOutputHandeler handels the default line output of run.Interactive
func updateOutputHandeler(line string, tty *os.File, scanner *bufio.Scanner) string {
	updateCommandOutput = append(updateCommandOutput, line)
	needRootErr(line, nil)
	gui.Echo(true, "cmdOut:", line)

	// pre definded regular expresion(s)
	DownloadingOrInstalling := `(?i)((downloading)|(installing))\s+(\d+).+(\d+)`
	DownloadingOrInstallLineMatch := `(((\w|-)+)-((\d|\.|\-|\w|\_)+)(\s\[((\w|\d)+)\])?\w*)$`

	if App.NormalMatch(DownloadingOrInstalling, line) {
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
	}

	return ""
}
