package app

import (
	"fmt"
	"log"

	"github.com/mjarkk/multipkg/pkg/run"
)

// this package detects the current OS
func detectOs() string {
	out, err := run.Run("cat /etc/lsb-release")
	if err != nil {
		log.Fatal("can't detect OS..")
	}
	fmt.Println("detect OS out:", out)
	return "Solus"
}
