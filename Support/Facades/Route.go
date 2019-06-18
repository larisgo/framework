package Facades

import (
	"github.com/larisgo/framework/Routing"
)

var Route func() *Routing.Router = func() *Routing.Router {
	return NewFacade("router").Get().(*Routing.Router)
}
