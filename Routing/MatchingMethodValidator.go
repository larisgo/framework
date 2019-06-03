package Routing

import (
	"github.com/larisgo/framework/Http"
)

type MethodValidator struct {
}

func NewMethodValidator() ValidatorInterface {
	return MethodValidator{}
}

func (this MethodValidator) matches(route *Route, request *Http.Request) bool {
	_, ok := route.Methods()[request.GetMethod()]
	return ok
}
