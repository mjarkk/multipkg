package app

import (
	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/handelers"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Run will setup the app
func Run() {
	app := &types.App{
		NormalMatch:    NormalMatch,
		Match:          Match,
		FindMatch:      FindMatch,
		Replace:        Replace,
		CleanupCli:     CleanupCli,
		CleanFindMatch: CleanFindMatch,
		FindAllMatch:   FindAllMatch,
	}
	gui.Setup(app)
	oss, err := detectOs()
	gui.CheckErr(err)
	handeler, err := handelers.Setup(app, oss)
	gui.CheckErr(err)
	DetectRunAction(handeler)
}
