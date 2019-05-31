package Routing

import (
	"regexp"
)

type UriValidator struct {
}

func NewUriValidator() ValidatorInterface {
	return UriValidator{}
}

func (this UriValidator) matches(route *Route, patch string) bool {
	if patch == "/" {
		patch = "/"
	} else {
		patch = "/" + patch
	}

	return regexp.MustCompile(route.GetCompiled().GetRegex()).MatchString(patch)
}
