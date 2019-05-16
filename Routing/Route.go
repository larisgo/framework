package Routing

type Route struct {
	Uri     string
	Methods []string
	Action  map[string]interface{}
}

func NewRoute(methods []string, uri string, action Action) (_Route *Route) {
	_Route = &Route{}
	_Route.Uri = uri
	_Route.Methods = methods
	_Route.Action = _Route.parseAction(action)
	return _Route
}

/**
 * Parse the route action into a standard array.
 *
 * @param  callable|array|null  $action
 * @return map[string]interface{}
 *
 * @throws \UnexpectedValueException
 */
func (this *Route) parseAction(action Action) map[string]interface{} {
	return map[string]interface{}{
		"a": nil,
	}
	// return RouteAction::parse($this->uri, $action);
}
