package solus

import (
	"fmt"
	"strings"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// app variable contains functions from other packages
var app *types.App

// Setup will be called to setup this pacakge
func Setup(App *types.App) *types.Handeler {
	app = App
	return &types.Handeler{
		Install:   Install,
		Reinstall: Reinstall,
		Remove:    Remove,
		Update:    Update,
		Search:    Search,
		Info:      Info,
	}
}

// Install installes a program
func Install(pkg string, flags *types.Flags) error {
	fmt.Println("called: Install")
	return nil
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

// Update a program
func Update(pkg string, flags *types.Flags) error {
	fmt.Println("called: Update")
	return nil
}

// Search for a program
func Search(pkg string, flags *types.Flags) error {
	out, err := run.Run("eopkg search --no-color --name" + pkg)
	needRootErr(out, err)
	regx := `((\w|-)+)(\t|\s)+-\s{0,1}((\(|\)|\w|\.|\s|,|\w-\w)+)(\n((\s|\t)$|)|$)`
	titles := app.FindAllMatch(out, regx, 1)
	Descriptions := app.FindAllMatch(out, regx, 4)
	returnVal := []types.PkgSearchOut{}
	for i, title := range titles {
		returnVal = append(returnVal, types.PkgSearchOut{
			Name:        title,
			Description: app.CleanupCli(Descriptions[i]),
		})
	}
	gui.PrintPkgSearch(&types.PkgSearchList{
		List: returnVal,
	})
	return nil
}

// Info about a program
func Info(pkg string, flags *types.Flags) error {
	out, err := run.Run("eopkg info " + pkg)
	needRootErr(out, err)
	returnVal := &types.PkgInfo{
		Name:              app.CleanFindMatch(out, `(Name(\s|\t)+:\s{0,})((\w|\s|\d|\.)+)`, 3),
		Installed:         app.NormalMatch("Installed package:", out),
		InstallledVersion: app.CleanFindMatch(out, `Installed package:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,})((\w|\s|\d|\.)+)`, 5),
		InstallledRelease: app.CleanFindMatch(out, `Installed package:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,}).+(release:\s{0,})((\s|\d|\.)+)`, 6),
		Version:           app.CleanFindMatch(out, `Package found in.{0,}:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,})((\w|\s|\d|\.)+)`, 5),
		Release:           app.CleanFindMatch(out, `Package found in.{0,}:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,}).+(release:\s{0,})((\s|\d|\.)+)`, 6),
		Dependencies:      strings.Split(app.CleanFindMatch(out, `(Dependencies(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-)+)\n`, 4), " "),
		Description:       app.CleanFindMatch(out, `(Description(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+)(\n\w+)(\s|\t)+:`, 4),
		Licenses:          strings.Split(app.CleanFindMatch(out, `(Licenses(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+)(\n\w+)(\s|\t)+:`, 4), " "),
		Component:         app.CleanFindMatch(out, `(Component(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+)(\n\w+)(\s|\t)+:`, 4),
	}
	gui.PrintPkgInfo(returnVal)
	return nil
}

func needRoot(checkstring string) bool {
	return app.NormalMatch("You have to be root for this operation", checkstring)
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
