package Routing

import (
	"regexp"
)

type HostValidator struct {
}

func NewHostValidator() ValidatorInterface {
	return HostValidator{}
}

func (this HostValidator) matches(route *Route, host string) bool {
	_regex := route.GetCompiled().GetHostRegex()
	if _regex == "" {
		return true
	}

	return regexp.MustCompile(_regex).MatchString(host)
}
