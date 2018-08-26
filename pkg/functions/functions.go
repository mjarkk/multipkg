package funs

import "github.com/mjarkk/multipkg/pkg/types"

// InArr checks if an item is in a array
func InArr(testArr []string, toFind string) bool {
	for _, item := range testArr {
		if item == toFind {
			return true
		}
	}
	return false
}

// InArrRegx checks if an item is in a array using a regular expression
func InArrRegx(app *types.App, testArr []string, regx string) bool {
	for _, item := range testArr {
		if app.NormalMatch(regx, item) {
			return true
		}
	}
	return false
}
