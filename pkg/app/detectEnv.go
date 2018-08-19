package app

import (
	"errors"
	"regexp"

	"github.com/mjarkk/multipkg/pkg/run"
)

// this package detects the current OS
func detectOs() (string, error) {
	defaultErr := "can't detect OS.."
	out, err := run.Run("cat /etc/lsb-release")
	if err != nil {
		return "", errors.New(defaultErr)
	}
	re := regexp.MustCompile("(DISTRIB_ID=)([a-zA-Z]+)")
	match := re.FindStringSubmatch(out)[2]
	if len(match) == 0 {
		return "", errors.New(defaultErr)
	}
	return match, nil
}
