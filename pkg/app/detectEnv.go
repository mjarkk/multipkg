package app

import (
	"errors"
	"io/ioutil"
	"regexp"
)

// this package detects the current OS
func detectOs() (string, error) {
	defaultErr := "can't detect OS.."
	out, err := ioutil.ReadFile("/etc/lsb-release")
	if err != nil {
		return "", errors.New(defaultErr)
	}
	re := regexp.MustCompile("(DISTRIB_ID=)([a-zA-Z]+)")
	match := re.FindStringSubmatch(string(out))[2]
	if len(match) == 0 {
		return "", errors.New(defaultErr)
	}
	return match, nil
}
