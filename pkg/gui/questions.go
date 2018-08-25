package gui

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/mjarkk/multipkg/pkg/types"
)

// Ask Asks a question to the user and returns the output
func Ask(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + "? ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return text, nil
}

// matchTrue matches a true input like yes and y
func matchTrue(app *types.App, toTest string) bool {
	return app.Match("yes|Yes|YES|Y|y", toTest)
}

// matchFalse matches a false input like no and n
func matchFalse(app *types.App, toTest string) bool {
	return app.Match("no|No|NO|N|n", toTest)
}

// needToMatchRegex makes sure the output of ask matches a specific regex
func needToMatchRegex(app *types.App, question string, regex string) (string, error) {
	AskTimes := 0
	for AskTimes < 4 {
		AskTimes = AskTimes + 1
		answer, err := Ask(question)
		if err != nil {
			return "", err
		}
		answer = app.Replace(answer, "", `^(\s|\t|\n)+|(\s|\t|\n)+$`)
		if app.NormalMatch(regex, answer) {
			return answer, nil
		}
		Echo(false, "Answer issn't valid")
	}
	return "", errors.New("Wrong answer to much times")
}

// YesNo asks a yes or no question and returns true or false
// this also adds a [y/n] automaticly to the question
func YesNo(app *types.App, question string) bool {
	out, err := needToMatchRegex(app, question+" [y/n]", `(?i)^(no|n|yes|y)$`)
	if err == nil && matchTrue(app, out) {
		return true
	}
	return false
}

// InstallQuestion Asks the user to if they are oke with installing a package
func InstallQuestion(app *types.App, Name string, Size string) bool {
	return YesNo(app, "Do you want to continue installing: "+Name+", size: "+Size)
}
