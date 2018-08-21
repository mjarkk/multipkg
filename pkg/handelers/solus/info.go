package solus

import (
	"strings"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

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
