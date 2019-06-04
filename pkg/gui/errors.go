package gui

import (
	"fmt"
	"os"
	"strings"

	"github.com/mjarkk/multipkg/pkg/types"
)

// CheckErr checks if there is an error if so it logs it error
func CheckErr(err error) {
	if err != nil {
		fmt.Println("CRITICAL ERROR:", err)
		os.Exit(1)
	}
}

// FriendlyErr returns a error message without making a user super scared
func FriendlyErr(errors ...interface{}) {
	fmt.Println(errors...)
	os.Exit(0)
}

// NoPkgsInstall give an error message that there are no packages to insatll
func NoPkgsInstall(pkg string) {
	if len(pkg) == 0 {
		FriendlyErr("No package(s) specified to install")
	}
}

// CheckBeForeInstall checks if the input pkg was right
func CheckBeForeInstall(pkg string, getInfo func(string) (*types.PkgInfo, error)) string {
	pkg = App.Replace(pkg, "", `^\s+|\s+$`)
	NoPkgsInstall(pkg)

	PKGs := strings.Split(pkg, " ")
	doNotExsist := []string{}
	for _, item := range PKGs {
		info, _ := getInfo(item)
		if !info.Installed {
			doNotExsist = append(doNotExsist, item)
		}
	}

	if len(doNotExsist) > 0 {
		Echo(false, "The following package(s) are already installed and wil be not installed")
		ShowList(doNotExsist, "dashList")
		os.Exit(0)
	}
	return pkg
}
