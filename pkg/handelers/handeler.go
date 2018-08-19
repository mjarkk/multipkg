package handelers

import (
	"errors"

	"github.com/mjarkk/multipkg/pkg/handelers/solus"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Setup sets up this package with the right functions
func Setup(osName string) (*types.Handeler, error) {
	if osName == "Solus" {
		return solus.Setup(), nil
	}
	return &types.Handeler{}, errors.New("No handeler found for this OS")
}
