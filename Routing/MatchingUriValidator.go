package Routing

import (
	"github.com/larisgo/framework/Http"
	"regexp"
)

type UriValidator struct {
}

func NewUriValidator() ValidatorInterface {
	return UriValidator{}
}

func (this UriValidator) matches(route *Route, request *Http.Request) bool {
	path := request.Path()
	if path == "/" {
		path = "/"
	} else {
		path = "/" + path
	}

	return regexp.MustCompile(route.GetCompiled().GetRegex()).MatchString(path)
}
