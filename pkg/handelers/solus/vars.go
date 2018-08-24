package solus

// this file contains all variables

var commandOutput = []string{}
var toExecuteAtEndOfNextLine = func(line string) {}

// this function will be executed after every
func toExecuteAtEndOfNextLineWrapper(line string) {
	toExecuteAtEndOfNextLine(line)
	// reset the toExecuteAtEndOfNextLine to the default value
	toExecuteAtEndOfNextLine = func(line string) {}
}

// PKG is the package(s) that will be insatlled
var PKG string
