package gui

import (
	"fmt"

	"github.com/mjarkk/multipkg/pkg/types"
)

// This file prints the output of every info command

// PrintPkgInfo prints out a programs info info
func PrintPkgInfo(packageInf *types.PkgInfo) {
	fmt.Println(packageInf)
}
