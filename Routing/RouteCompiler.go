package Routing

import (
	"fmt"
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
func (this *RouteCompiler) Compile() {
	optionals := this.getOptionalParameters()
	uri := regexp.MustCompile(`\{(\w+?)\?\}`).ReplaceAllString(this.route.Uri(), `{$1}`)
	fmt.Println(uri)

	// return (
	//     new SymfonyRoute($uri, $optionals, $this->route->wheres, ['utf8' => true], $this->route->getDomain() ?: '')
	// )->compile();
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
