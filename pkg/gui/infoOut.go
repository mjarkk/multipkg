package gui

import (
	"fmt"

	"github.com/mjarkk/multipkg/pkg/types"
)

// This file prints the output of every info command

// PrintPkgInfo prints out a programs info info
func PrintPkgInfo(packageInf *types.PkgInfo) {
	fmt.Println("Name:", packageInf.Name)
	fmt.Println("Installed:", packageInf.Installed)
	fmt.Println("Installled Version:", packageInf.InstallledVersion)
	fmt.Println("Installled Release:", packageInf.InstallledRelease)
	fmt.Println("Version:", packageInf.Version)
	fmt.Println("Release:", packageInf.Release)
	fmt.Println("Dependencies:", packageInf.Dependencies)
	fmt.Println("Description:", packageInf.Description)
	fmt.Println("Licenses:", packageInf.Licenses)
	fmt.Println("Component:", packageInf.Component)

}
