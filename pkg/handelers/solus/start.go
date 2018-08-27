package solus

import (
	"github.com/mjarkk/multipkg/pkg/types"
)

// App variable contains functions from other packages
var App *types.App

// Setup will be called to setup this pacakge
func Setup(app *types.App) *types.Handeler {
	App = app
	return &types.Handeler{
		Install:   Install,
		Reinstall: Reinstall,
		Remove:    Remove,
		Update:    Update,
		Search:    Search,
		Info:      Info,
	}
}
