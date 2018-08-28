package arch

import (
	"fmt"

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

// Reinstall handels the reinstall command
func Reinstall(pkg string, flags *types.Flags) error {
	fmt.Println("reinstall is not working yet...")
	return nil
}

// Remove handels the remove command
func Remove(pkg string, flags *types.Flags) error {
	fmt.Println("remove is not working yet...")
	return nil
}

// Search handels the search command
func Search(pkg string, flags *types.Flags) error {
	fmt.Println("search is not working yet...")
	return nil
}

// Info handels the info command
func Info(pkg string, flags *types.Flags) error {
	fmt.Println("info is not working yet...")
	return nil
}
