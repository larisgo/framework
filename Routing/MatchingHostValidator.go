package Routing

import (
	"github.com/larisgo/framework/Http"
	"regexp"
)

type HostValidator struct {
}

func NewHostValidator() ValidatorInterface {
	return HostValidator{}
}

func (this HostValidator) matches(route *Route, request *Http.Request) bool {
	_regex := route.GetCompiled().GetHostRegex()
	if _regex == "" {
		return true
	}

	return regexp.MustCompile(_regex).MatchString(request.GetHost())
}
