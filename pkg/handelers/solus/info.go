package solus

import (
	"errors"
	"strings"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Info about a program
func Info(pkg string, flags *types.Flags) error {
	data, err := GetInfo(pkg)
	if err != nil {
		gui.FriendlyErr()
	}
	gui.PrintPkgInfo(data)
	return nil
}

// GetInfo gets info about the package and wraps it inside a object
func GetInfo(pkg string) (*types.PkgInfo, error) {
	out, err := run.Run("eopkg info --no-color " + pkg)
	needRootErr(out, err)
	returnVal := &types.PkgInfo{
		Name:              App.CleanFindMatch(out, `(Name(\s|\t)+:\s{0,})((\w|\s|\d|\.)+)`, 3),
		Installed:         App.NormalMatch("Installed package:", out),
		InstallledVersion: App.CleanFindMatch(out, `Installed package:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,})((\w|\s|\d|\.)+)`, 5),
		InstallledRelease: App.CleanFindMatch(out, `Installed package:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,}).+(release:\s{0,})((\s|\d|\.)+)`, 6),
		Version:           App.CleanFindMatch(out, `Package found in.{0,}:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,})((\w|\s|\d|\.)+)`, 5),
		Release:           App.CleanFindMatch(out, `Package found in.{0,}:(\n|\s){0,20}(Name(\s|\t){0,}:\s{0,}).+(version:\s{0,}).+(release:\s{0,})((\s|\d|\.)+)`, 6),
		Dependencies:      strings.Split(App.CleanFindMatch(out, `(Dependencies(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-)+)\n`, 4), " "),
		Description:       App.CleanFindMatch(out, `(Description(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+)(\n\w+)(\s|\t)+:`, 4),
		Licenses:          strings.Split(App.CleanFindMatch(out, `(Licenses(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+)(\n\w+)(\s|\t)+:`, 4), " "),
		Component:         App.CleanFindMatch(out, `(Component(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+)(\n\w+)(\s|\t)+:`, 4),
		Architecture:      App.CleanFindMatch(out, `Package found in(.|\n)*(Architecture(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+),`, 5),
		InstallSize:       App.CleanFindMatch(out, `Package found in(.|\n)*(Architecture(\s|\t)+:(\s|\t)+)((\w|\s|\d|\.|-|,)+),\s*Installed\s*Size:(\s|\t)*((\d|\.)*(\s\w*)?)`, 8),
	}
	if returnVal.Name == "" {
		return &types.PkgInfo{}, errors.New("package not found")
	}
	return returnVal, nil
}
