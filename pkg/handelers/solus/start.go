package solus

import (
	"fmt"

	"github.com/mjarkk/multipkg/pkg/types"
)

func Setup() *types.Handeler {
	return &types.Handeler{
		Install:   Install,
		Reinstall: Reinstall,
		Remove:    Remove,
		Update:    Update,
		Search:    Search,
		Info:      Info,
	}
}

// Install installes a program
func Install(pkg string, flags *types.Flags) error {
	fmt.Println("called: Install")
	return nil
}

// Reinstall reinstalles a program
func Reinstall(pkg string, flags *types.Flags) error {
	fmt.Println("called: Reinstall")
	return nil
}

// Remove a program
func Remove(pkg string, flags *types.Flags) error {
	fmt.Println("called: Remove")
	return nil
}

// Update a program
func Update(pkg string, flags *types.Flags) error {
	fmt.Println("called: Update")
	return nil
}

// Search for a program
func Search(pkg string, flags *types.Flags) error {
	fmt.Println("called: Search")
	return nil
}

// Info about a program
func Info(pkg string, flags *types.Flags) error {
	fmt.Println("called: Info")
	return nil
}
