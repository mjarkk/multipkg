package solus

import (
	"fmt"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/types"
)

// App variable contains functions from other packages
var App *types.App

// Setup will be called to setup this pacakge
func Setup(app *types.App) *types.Handeler {
	App = app
	return &types.Handeler{
		Install:   Install,
		Reinstall: Reinstall,
		Remove:    Remove,
		Update:    Update,
		Search:    Search,
		Info:      Info,
	}
}

// Reinstall reinstalles a program
func Reinstall(pkg string, flags *types.Flags) error {
	fmt.Println("called: Reinstall")
	return nil
}

// Remove a program
func Remove(pkg string, flags *types.Flags) error {
	fmt.Println("called: Remove")
	return nil
}

func needRoot(checkstring string) bool {
	return App.NormalMatch("You have to be root for this operation", checkstring)
}

func needRootErr(input string, err error) {
	needRootErr := "You need root permissions execute this command"
	if err != nil {
		if needRoot(err.Error()) {
			gui.FriendlyErr(needRootErr)
		} else {
			gui.FriendlyErr("error:", err.Error())
		}
	}
	if needRoot(input) {
		gui.FriendlyErr(needRootErr)
	}
}
