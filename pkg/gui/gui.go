package gui

import "fmt"

// Echo prints something to the screen
func Echo(debug bool, inputs ...interface{}) {
	if !debug {
		fmt.Println(inputs...)
	}
}
