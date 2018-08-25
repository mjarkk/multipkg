package solus

import "github.com/mjarkk/multipkg/pkg/types"

// this file contains all variables

var commandOutput = []string{}
var toExecuteAtEndOfNextLine = func(line string, extraData types.Flags) {}

var lastExecLineData = types.Flags{}

// this function will be executed after every
func toExecuteAtEndOfNextLineWrapper(line string) {
	toExecuteAtEndOfNextLine(line, lastExecLineData)
	// reset the toExecuteAtEndOfNextLine to the default value
	toExecuteAtEndOfNextLine = func(line string, extraData types.Flags) {}
}

// PKG is the package(s) that will be insatlled
var PKG string
