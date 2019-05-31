package Routing

import (
	"regexp"
)

type RouteCompiler struct {
	route *Route
}

/**
 * Create a new Route compiler instance.
 *
 * @param  Routing\Route  route
 * @return void
 */
func NewRouteCompiler(route *Route) *RouteCompiler {
	return &RouteCompiler{route: route}
}

/**
 * Compile the route.
 *
 * @return Routing\CompiledRoute
 */
func (this *RouteCompiler) Compile() *CompiledRoute {
	optionals := this.getOptionalParameters()
	uri := regexp.MustCompile(`\{(\w+?)\?\}`).ReplaceAllString(this.route.Uri(), `{$1}`)
	return NewSymfonyRoute(uri, optionals, map[string]string{}, map[string]interface{}{"utf8": true}, this.route.GetDomain(), map[string]bool{}, map[string]bool{}, "").Compile()
}

/**
 * Get the optional parameters for the route.
 *
 * @return array
 */
func (this *RouteCompiler) getOptionalParameters() (v map[string]string) {

	v = make(map[string]string)
	for _, match := range regexp.MustCompile(`\{(\w+?)\?\}`).FindAllStringSubmatch(this.route.Uri(), -1) {
		if len(match) == 2 {
			v[match[1]] = ""
		}
	}
	return v
}
