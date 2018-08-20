package app

import (
	"os"
	"regexp"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/types"
)

// DetectRunAction from command line aruments what to do next and do that
func DetectRunAction(handeler *types.Handeler) {
	cliActions := actionToRun()
	firstArg := cliActions["commandArg"]
	flags := &types.Flags{
		"force": len(cliActions["force"]) > 0,
	}
	switch cliActions["command"] {
	case "Install":
		handeler.Install(firstArg, flags)
	case "Reinstall":
		handeler.Reinstall(firstArg, flags)
	case "Remove":
		handeler.Remove(firstArg, flags)
	case "Update":
		handeler.Update(firstArg, flags)
	case "Search":
		handeler.Search(firstArg, flags)
	case "Info":
		handeler.Info(firstArg, flags)
	}
}

func match(regx string, arg string) bool {
	matched, err := regexp.MatchString("^("+regx+")$", arg)
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
			out["command"] = "Install"
		} else if match("reinstall|rein|ri", arg) && runCommand {
			out["command"] = "Reinstall"
		} else if match("remove|re|r", arg) && runCommand {
			out["command"] = "Remove"
		} else if match("update|up|u", arg) && runCommand {
			out["command"] = "Update"
		} else if match("search|se|s", arg) && runCommand {
			out["command"] = "Search"
		} else if match("info|inf", arg) && runCommand {
			out["command"] = "Info"
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
