package Routing

import (
	"github.com/larisgo/framework/Http"
)

type ValidatorInterface interface {
	matches(*Route, *Http.Request) bool
}
