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
	fmt.Println(RightPad("Name", " ", leng), "-", packageInf.Name)
	fmt.Println(RightPad("Installed", " ", leng), "-", packageInf.Installed)
	fmt.Println(RightPad("Installled Version", " ", leng), "-", packageInf.InstallledVersion)
	fmt.Println(RightPad("Installled Release", " ", leng), "-", packageInf.InstallledRelease)
	fmt.Println(RightPad("Version", " ", leng), "-", packageInf.Version)
	fmt.Println(RightPad("Release", " ", leng), "-", packageInf.Release)
	fmt.Println(RightPad("Dependencies", " ", leng), "-", packageInf.Dependencies)
	fmt.Println(RightPad("Description", " ", leng), "-", packageInf.Description)
	fmt.Println(RightPad("Licenses", " ", leng), "-", packageInf.Licenses)
	fmt.Println(RightPad("Component", " ", leng), "-", packageInf.Component)
	fmt.Println(RightPad("Architecture", " ", leng), "-", packageInf.Architecture)
	fmt.Println(RightPad("Size", " ", leng), "-", packageInf.InstallSize)
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
