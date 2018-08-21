package solus

import (
	"github.com/mjarkk/multipkg/pkg/gui"
	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Search for a program
func Search(pkg string, flags *types.Flags) error {
	out, err := run.Run("eopkg search --no-color --name" + pkg)
	needRootErr(out, err)
	regx := `((\w|-)+)(\t|\s)+-\s{0,1}((\(|\)|\w|\.|\s|,|\w-\w)+)(\n((\s|\t)$|)|$)`
	titles := app.FindAllMatch(out, regx, 1)
	Descriptions := app.FindAllMatch(out, regx, 4)
	returnVal := []types.PkgSearchOut{}
	for i, title := range titles {
		returnVal = append(returnVal, types.PkgSearchOut{
			Name:        title,
			Description: app.CleanupCli(Descriptions[i]),
		})
	}
	gui.PrintPkgSearch(&types.PkgSearchList{
		List: returnVal,
	})
	return nil
}
