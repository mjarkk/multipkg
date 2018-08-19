package types

type Handeler struct {
	Install func()
	Update  func()
	Remove  func()
}
