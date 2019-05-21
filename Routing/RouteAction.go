package Routing

import (
	"fmt"
)

type routeAction struct {
	Domain     string
	Prefix     string
	Uses       Action
	As         string
	Controller string
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
		this.Uses = Action(func() {
			panic(fmt.Sprintf("Route for [%s] has no action.", uri))
		})
		return this
	}
	this.Uses = action
	return this
}

func RouteAction() *routeAction {
	return &routeAction{}
}
