package gui

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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
func matchTrue(toTest string) bool {
	return App.Match("yes|Yes|YES|Y|y", toTest)
}

// matchFalse matches a false input like no and n
func matchFalse(toTest string) bool {
	return App.Match("no|No|NO|N|n", toTest)
}

// needToMatchRegex makes sure the output of ask matches a specific regex
func needToMatchRegex(question string, regex string) (string, error) {
	AskTimes := 0
	for AskTimes < 4 {
		AskTimes = AskTimes + 1
		answer, err := Ask(question)
		if err != nil {
			return "", err
		}
		answer = App.Replace(answer, "", `^(\s|\t|\n)+|(\s|\t|\n)+$`)
		if App.NormalMatch(regex, answer) {
			return answer, nil
		}
		Echo(false, "Answer issn't valid")
	}
	return "", errors.New("Wrong answer to much times")
}

// YesNo asks a yes or no question and returns true or false
// this also adds a [y/n] automaticly to the question
func YesNo(question string, defaultAnsw bool) bool {
	yesNoText := " [N/y]"
	if defaultAnsw {
		yesNoText = " [n/Y]"
	}
	out, err := needToMatchRegex(question+yesNoText, `(?i)^(no|n|yes|y|)$`)
	if err == nil && matchTrue(out) {
		return true
	}
	if len(out) == 0 {
		return defaultAnsw
	}
	return false
}

// InstallQuestion Asks the user to if they are oke with installing a package
func InstallQuestion(Name string, Size string, ifAutoYes bool) bool {
	if autoYes {
		return ifAutoYes
	}
	return YesNo("Do you want to continue installing: "+Name+", size: "+Size, ifAutoYes)
}

// UpdateQuestion asks if the package manager can update ... packages
func UpdateQuestion(Size string, ifAutoYes bool) bool {
	if autoYes {
		return ifAutoYes
	}
	return YesNo("Do you want to update these pacakges size: "+Size, ifAutoYes)
}

// RemoveQuestion asks if it's oke to remove a pacakge
func RemoveQuestion(Name string, ifAutoYes bool) bool {
	if autoYes {
		return ifAutoYes
	}
	return YesNo("Do you want to remove: "+Name, ifAutoYes)
}
