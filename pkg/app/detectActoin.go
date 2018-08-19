package app

import (
	"fmt"
	"os"
	"regexp"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/types"
)

// DetectRunAction from command line aruments what to do next and do that
func DetectRunAction(handeler *types.Handeler) {
	cliActions := actionToRun()
	fmt.Println(cliActions)
}

func match(regx string, arg string) bool {
	matched, err := regexp.MatchString("^"+regx+"$", arg)
	if err != nil {
		return false
	}
	return matched
}

func matchFullFlag(full string, letter string, arg string) bool {
	return match("--"+full+"|-\\w{0,}"+letter+"\\w{0,}", arg)
}

func matchFlag(full string, arg string) bool {
	return match("--"+full, arg)
}

func actionToRun() types.Obj {
	out := make(types.Obj)
	out["command"] = ""
	out["force"] = ""
	out["commandArg"] = ""

	for _, arg := range os.Args[1:] {
		runCommand := out["command"] == ""

		if match("install|in|i", arg) && runCommand {
			out["command"] = "install"
		} else if match("reinstall|rein|ri", arg) && runCommand {
			out["command"] = "reinstall"
		} else if match("remove|re|r", arg) && runCommand {
			out["command"] = "remove"
		} else if match("update|up|u", arg) && runCommand {
			out["command"] = "remove"
		} else if match("[^-].+", arg) && !runCommand {
			out["commandArg"] = out["commandArg"] + " " + arg
		} else if matchFullFlag("force", "f", arg) {
			out["force"] = "true"
		} else if runCommand {
			gui.FriendlyErr("multipkg: command " + arg + " not found")
		}

	}

	if out["command"] == "" {
		gui.FriendlyErr("No command given \nUse: multipkg --help for more info")
	}

	return out
}
