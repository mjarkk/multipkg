package types

// Flags type for input flags
type Flags map[string]interface{}

// HandelerFunc is every main function in handelers/**/*.go
type HandelerFunc func(string, *Flags) error

// Handeler is what a handeler file must return when ran setup
type Handeler struct {
	Install   HandelerFunc
	Reinstall HandelerFunc
	Remove    HandelerFunc
	Update    HandelerFunc
	Search    HandelerFunc
	Info      HandelerFunc
}

// Obj is just a wrapper around map[string]string
type Obj = map[string]string

// PkgInfo is what needs to printed out in gui/infoOut.go
type PkgInfo struct {
	Name              string   // git
	Installed         bool     // true
	InstallledVersion string   // 2.18.0
	InstallledRelease string   // 70
	Version           string   // 2.18.0
	Release           string   // 70
	Dependencies      []string // [glibc zlib expat curl openssl perl-authen-sasl perl-error ...]
	Description       string   // Git is a fast, scalable, distributed revision control system with an...
	Licenses          []string // GPL-2.0
	Component         string   // programming.tools
}

// PkgSearchOut is every array item in PkgSearchList
type PkgSearchOut struct {
	Name        string // git
	Description string // Git is a fast, scalable, distributed revision control system with an...
}

// PkgSearchList is what will be send to the print output off package search output
type PkgSearchList struct {
	List []PkgSearchOut
}

// App can be used in other packages to execute functions from package app
type App struct {
	NormalMatch    func(regx string, arg string) bool
	Match          func(regx string, arg string) bool
	FindMatch      func(input string, regx string, selector int) string
	Replace        func(toReplace string, Replaceval string, regx string) string
	CleanupCli     func(input string) string
	CleanFindMatch func(input string, regx string, selector int) string
	FindAllMatch   func(input string, regx string, selector int) []string
}
