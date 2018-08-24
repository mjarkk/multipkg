package gui

import (
	"fmt"
	"os"
	"strings"

	"github.com/mjarkk/multipkg/pkg/functions"
)

// debugMode tells the program to print more info if needed
var debugMode = false

// Echo prints something to the screen
func Echo(debug bool, inputs ...interface{}) {
	if !debug || debugMode {
		fmt.Println(inputs...)
	}
}

// Setup the gui package
func Setup() {
	debugMode = funs.InArr(os.Args[1:], "--debug")
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
