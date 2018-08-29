package gui

import (
	"fmt"
	"strings"

	"github.com/mjarkk/multipkg/pkg/types"
)

// This file prints the output of every info command

// ReturnLongest returns the longes string in a []string
func ReturnLongest(inputs []string) int {
	longest := 1
	for _, item := range inputs {
		if len(item) > longest {
			longest = len(item)
		}
	}
	return longest
}

// PrintPkgInfo prints out a programs info info
func PrintPkgInfo(packageInf *types.PkgInfo) {
	leng := ReturnLongest([]string{"Installled Version", "Installled Release"})
	emptryVars := []string{}
	printLn := func(key string, value interface{}) {
		if value == "" {
			emptryVars = append(emptryVars, key)
		} else {
			fmt.Println(RightPad(key, " ", leng), "-", value)
		}
	}
	printLn("Name", packageInf.Name)

	printLn("Installed", packageInf.Installed)
	if packageInf.Installed {
		printLn("Installled Version", packageInf.InstallledVersion)
		printLn("Installled Release", packageInf.InstallledRelease)
	}
	printLn("Version", packageInf.Version)
	printLn("Release", packageInf.Release)
	printLn("Dependencies", packageInf.Dependencies)
	printLn("Description", packageInf.Description)
	printLn("Licenses", packageInf.Licenses)
	printLn("Component", packageInf.Component)
	printLn("Architecture", packageInf.Architecture)
	printLn("Size", packageInf.InstallSize)
	if len(emptryVars) > 0 {
		Echo(false, "\nthe folowing item where not found:")
		ShowList(emptryVars, "normal")
	}
}

// PrintPkgSearch prints out a list of found packages from a repo
func PrintPkgSearch(SearchOut *types.PkgSearchList) {
	testArr := []string{}
	for _, item := range SearchOut.List {
		testArr = append(testArr, item.Name)
	}
	leng := ReturnLongest(testArr)
	for _, searchRes := range SearchOut.List {
		fmt.Println(RightPad(searchRes.Name, " ", leng), " - ", searchRes.Description)
	}
}

// RightPad view: https://github.com/git-time-metric/gtm/blob/master/util/string.go#L53-L88
func RightPad(s string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// ProgressIn shows the progress in the downloads, installes and so on
func ProgressIn(item string, from string, action string, pkg string) {
	progress := "[" + RightPad(item, " ", len(from)) + "/" + from + "]"
	fmt.Println(progress, action+":", pkg)
}

// Installed shows that a package is installed
func Installed(pkg string) {
	Echo(false, "Installed:", pkg)
}
