package handelers

import (
	"errors"

	"github.com/mjarkk/multipkg/pkg/handelers/arch"
	"github.com/mjarkk/multipkg/pkg/handelers/solus"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Setup is started as first function of this program
func Setup(App *types.App, osName string) (*types.Handeler, error) {
	switch osName {
	case "Solus":
		return solus.Setup(App), nil
	case "Arch":
		return arch.Setup(App), nil
	}
	return &types.Handeler{}, errors.New("No handeler found for this OS")
}
