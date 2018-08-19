package app

import (
	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/handelers"
)

// Run will setup the app
func Run() {
	oss, err := detectOs()
	gui.CheckErr(err)
	handeler, err := handelers.Setup(oss)
	gui.CheckErr(err)
	DetectRunAction(handeler)
}
