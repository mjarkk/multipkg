package solus

import "github.com/mjarkk/multipkg/pkg/types"

// this file contains all variables

var commandOutput = []string{}

var toExecuteAtEndOfNextLine = func(line string, extraData types.Flags) {}

var lastExecLineData = types.Flags{}
var nextExecFuncMatchRegx = ``

// this function will be executed after every
func toExecuteAtEndOfNextLineWrapper(line string) {
	if App.NormalMatch(nextExecFuncMatchRegx, line) {
		toExecuteAtEndOfNextLine(line, lastExecLineData)
		// reset the toExecuteAtEndOfNextLine to the default value
		toExecuteAtEndOfNextLine = func(line string, extraData types.Flags) {}
		nextExecFuncMatchRegx = ``
	}
}

// PKG is the package(s) that will be handeld by install, remove, update, etc...
var PKG string
