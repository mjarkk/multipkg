package types

type Flags map[string]interface{}

type HandelerFunc func(string, *Flags) error

type Handeler struct {
	Install   HandelerFunc
	Reinstall HandelerFunc
	Remove    HandelerFunc
	Update    HandelerFunc
	Search    HandelerFunc
	Info      HandelerFunc
}

type Obj = map[string]string
