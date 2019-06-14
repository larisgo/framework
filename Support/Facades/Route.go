package Facades

var Route func() interface{} = func() interface{} {
	return NewFacade("router").GetRaw()
}
