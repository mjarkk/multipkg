package gui

import (
	"fmt"
	"os"
	"strings"

	"github.com/mjarkk/multipkg/pkg/functions"
	"github.com/mjarkk/multipkg/pkg/types"
)

// debugMode tells the program to print more info if needed
var debugMode = false
var autoYes = false

// App variable contains functions from other packages
var App *types.App

// Echo prints something to the screen
func Echo(debug bool, inputs ...interface{}) {
	if !debug || debugMode {
		fmt.Println(inputs...)
	}
}

// Echof ueses Printf instaid of Println
func Echof(debug bool, input string, arguments ...interface{}) {
	if !debug || debugMode {
		fmt.Printf(input, arguments...)
	}
}

// Setup the gui package
func Setup(app *types.App) {
	App = app
	debugMode = funs.InArr(os.Args[1:], "--debug")
	autoYes = funs.InArr(os.Args[1:], "--yes") || funs.InArrRegx(app, os.Args[1:], `^(\s*\-(\d|\w)*y(\d|\w)*\s*)$`)
}

// ShowList shows a list of inputed strings
// lsitType = "dashList" || "normal"
func ShowList(input []string, listType string) {
	if listType == "dashList" {
		for _, item := range input {
			fmt.Println(" -", item)
		}
	} else {
		// normal
		fmt.Println(strings.Join(input, ", "))
	}
}
