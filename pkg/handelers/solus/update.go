package solus

import (
	"fmt"

	"github.com/mjarkk/multipkg/pkg/run"
	"github.com/mjarkk/multipkg/pkg/types"
)

// Update a program
func Update(pkg string, flags *types.Flags) error {
	err := run.Interactive("eopkg --no-color upgrade", updateOutputHandeler, updateErrHandeler)
	fmt.Println(err)
	return nil
}

func updateOutputHandeler(line string) string {
	fmt.Println("out:", line)
	return line
}

func updateErrHandeler(line string) string {
	fmt.Println("err:", line)
	return line
}
