package Routing

import (
	"fmt"
	"github.com/larisgo/framework/Errors"
	"github.com/larisgo/framework/Http"
)

type routeAction struct {
	Domain     string
	Prefix     string
	Uses       Action
	As         string
	Middleware []string
}

/**
 * Parse the route action into a standard array.
 *
 * @param  uri string  action Action
 * @return *routeAction
 *
 */
func (this *routeAction) Parse(uri string, action Action) *routeAction {
	if action == nil {
		this.Uses = Action(func(*Http.Request) *Http.Response {
			panic(Errors.NewLogicException(fmt.Sprintf("Route for [%s] has no action.", uri)))
		})
		return this
	}
	this.Uses = action
	return this
}

func RouteAction() (this *routeAction) {
	this = &routeAction{}
	this.Middleware = []string{}
	return this
}
