package app

import (
	"errors"
	"io/ioutil"
)

// this package detects the current OS
func detectOs() (string, error) {
	defaultErr := "can't detect OS.."
	out, err := ioutil.ReadFile("/etc/lsb-release")
	if err != nil {
		return "", errors.New(defaultErr)
	}
	match := FindMatch(string(out), "(DISTRIB_ID=)([a-zA-Z]+)", 2)
	if len(match) == 0 {
		return "", errors.New(defaultErr)
	}
	return match, nil
}
