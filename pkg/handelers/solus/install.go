package solus

import (
	"fmt"

	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Install installes a program
func Install(pkg string, flags *types.Flags) error {
	err := run.Interactive(App, "eopkg --no-color install "+pkg, installOutputHandeler)
	fmt.Println(err)
	return nil
}

var installCommandOutput = []string{}

// updateOutputHandeler handels the default line output of run.Interactive
func installOutputHandeler(line string) string {
	installCommandOutput = append(installCommandOutput, line)
	needRootErr(line, nil)
	fmt.Println("\ncmdOut:", line)
	return ""
}
