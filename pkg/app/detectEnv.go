package app

import (
	"errors"

	"github.com/mjarkk/multipkg/pkg/run"
)

// this package detects the current OS

// detectOs detects the current os from a command
func detectOs() (string, error) {
	_, err := run.Run("pacman --version")
	if err == nil {
		return "Arch", nil
	}
	_, err = run.Run("eopkg --version")
	if err == nil {
		return "Solus", nil
	}
	return "", errors.New("Can't detect OS")
}
