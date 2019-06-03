package Routing

import (
	"github.com/larisgo/framework/Http"
)

type SchemeValidator struct {
}

func NewSchemeValidator() ValidatorInterface {
	return SchemeValidator{}
}

func (this SchemeValidator) matches(route *Route, request *Http.Request) bool {
	if route.HttpOnly() {
		return !request.Secure()
	} else if route.Secure() {
		return request.Secure()
	}
	return true
}
