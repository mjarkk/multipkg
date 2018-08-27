package solus

import (
	"github.com/mjarkk/multipkg/pkg/gui"
)

// needRoot checks if a input string matches a need to root regex
func needRoot(checkstring string) bool {
	return App.NormalMatch("You have to be root for this operation", checkstring)
}

var needRootMsg = "You need root permissions execute this command"

func needRootErr(input string, err error) {
	if needRoot(input) {
		gui.FriendlyErr(needRootMsg)
	} else if err != nil {
		if needRoot(err.Error()) {
			gui.FriendlyErr(needRootMsg)
		} else {
			gui.FriendlyErr("error:", err.Error())
		}
	}
}
