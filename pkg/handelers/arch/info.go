package arch

import (
	"errors"
	"strings"

	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Info handels the info command
func Info(pkg string, flags *types.Flags) error {
	pkgInf, err := getInfo(pkg)
	if err != nil {
		gui.FriendlyErr(err)
	}

	gui.PrintPkgInfo(pkgInf)

	return nil
}

func getInfo(pkg string) (*types.PkgInfo, error) {
	PKG = App.Replace(pkg, "", `^\s+|\s+$`)

	outLocal, err := run.Run("pacman -Qi " + PKG)
	gui.Echof(true, "cmdOut: %q, error: %v\n", outLocal, err)

	needRootErr(outLocal, nil)

	Installed := false
	InstallledVersion := ""
	if !App.NormalMatch(`error\: package .+ was not found`, outLocal) {
		Installed = true
		InstallledVersion = App.CleanFindMatch(outLocal, `(?i)Version\s+:\s*(.+)`, 1)
	}

	outInter, err := run.Run("pacman -Si " + PKG)
	gui.Echof(true, "cmdOut: %q, error: %v\n", outInter, err)
	if App.NormalMatch(`error: package .+ was not found`, outInter) {
		return &types.PkgInfo{}, errors.New("Package not found")
	}

	return &types.PkgInfo{
		Name:              App.CleanFindMatch(outInter, `(?i)Name\s+:\s*(.+)`, 1),
		Installed:         Installed,
		InstallledVersion: InstallledVersion,
		Version:           App.CleanFindMatch(outInter, `(?i)Version\s+:\s*(.+)`, 1),
		Dependencies:      strings.Split(App.CleanFindMatch(outInter, `(?i)Depends On\s+:\s+((\w|\s|\d|\.|-)+)\n`, 1), " "),
		Description:       App.CleanFindMatch(outInter, `(?i)Description\s+:\s*(.+)`, 1),
		Licenses:          strings.Split(App.CleanFindMatch(outInter, `(?i)Licenses\s+:\s+((\w|\s|\d|\.|-)+)\n`, 1), " "),
		Architecture:      App.CleanFindMatch(outInter, `(?i)Architecture\s+:\s*(.+)`, 1),
		InstallSize:       App.CleanFindMatch(outInter, `(?i)Installed Size\s+:\s*(.+)`, 1),
		Component:         "", // issn't provided by pacman -Si <package name>
		Release:           "", // issn't provided by pacman -Si <package name>
		InstallledRelease: "", // issn't provided by pacman -Si <package name>
	}, nil
}
