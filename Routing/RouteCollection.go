package Routing

type RouteCollection struct {
	routes     map[string]map[string]*Route
	allRoutes  map[string]*Route
	nameList   map[string]*Route
	actionList map[string]*Route
}

func NewRouteCollection() (this *RouteCollection) {
	this = &RouteCollection{}
	this.routes = make(map[string]map[string]*Route)
	this.allRoutes = make(map[string]*Route)
	this.nameList = make(map[string]*Route)
	this.actionList = make(map[string]*Route)
	return this
}

/**
 * Add a Route instance to the collection.
 *
 * @param  Routing\Route  route
 * @return Routing\Route
 */
func (this *RouteCollection) Add(route *Route) *Route {
	this.addToCollections(route)

	this.addLookups(route)

	return route
}

/**
 * Add the given route to the arrays of routes.
 *
 * @param  Routing\Route  route
 * @return void
 */
func (this *RouteCollection) addToCollections(route *Route) {
	domainAndUri := route.GetDomain() + route.Uri()
	var method string
	for method, _ = range route.Methods() {
		if _, ok := this.routes[method]; ok {
			this.routes[method][domainAndUri] = route
		} else {
			this.routes[method] = map[string]*Route{
				domainAndUri: route,
			}
		}
	}
	this.allRoutes[method+domainAndUri] = route
}

/**
 * Add the route to any look-up tables if necessary.
 *
 * @param  Routing\Route  route
 * @return void
 */
func (this *RouteCollection) addLookups(route *Route) {
	// If the route has a name, we will add it to the name look-up table so that we
	// will quickly be able to find any route associate with a name and not have
	// to iterate through every route every time we need to perform a look-up.
	if name := route.GetName(); name != "" {
		this.nameList[name] = route
	}

	// When the route is routing to a controller we will also store the action that
	// is used by the route. This will let us reverse route to controllers while
	// processing a request and easily generate URLs to the given controllers.
	// action := route.Action

	// if action.Controller != "" {
	// 	this.addToActionList(action, route)
	// }
}

/**
 * Add a route to the controller action dictionary.
 *
 * @param  array  action
 * @param  Route  route
 * @return void
 */
func (this *RouteCollection) addToActionList(action Action, route *Route) {
	// this.actionList[action.Controller] = route
}

/**
 * Refresh the name look-up table.
 *
 * This is done in case any names are fluently defined or if routes are overwritten.
 *
 * @return void
 */
func (this *RouteCollection) RefreshNameLookups() {
	this.nameList = make(map[string]*Route, 0)
	for _, route := range this.allRoutes {
		if name := route.GetName(); name != "" {
			this.nameList[name] = route
		}
	}
}

/**
 * Find the first route matching a given request.
 *
 * @param  Http\Request  $request
 * @return Routing\Route
 *
 * @throws \Symfony\Component\HttpKernel\Exception\NotFoundHttpException
 */
func (this *RouteCollection) Match() {
	// routes := this.Get("GET")

	// // First, we will see if we can find a matching route for this current request
	// // method. If we can, great, we can just return it so that it can be called
	// // by the consumer. Otherwise we will check for routes with another verb.
	// route := this.matchAgainstRoutes(routes, "1", true)

	// if route != nil {
	// 	// return route.Bind("request")
	// }

	// If no route was found we will now check if a matching route is specified by
	// another HTTP verb. If it is we will need to throw a MethodNotAllowed and
	// inform the user agent of which HTTP verb it should use for this route.
	// others := this.checkForAlternateVerbs("request")

	// if (count($others) > 0) {
	//     return this.getRouteForMethods($request, $others);
	// }

	// throw new NotFoundHttpException;
}

/**
 * Determine if a route in the array matches the request.
 *
 * @param  array  $routes
 * @param  \Illuminate\Http\Request  $request
 * @param  bool  $includingMethod
 * @return \Illuminate\Routing\Route|null
 */
func (this *RouteCollection) matchAgainstRoutes(routes map[string]*Route, request string, includingMethod bool) *Route {
	_fallbacks := []*Route{}
	_routes := []*Route{}
	for _, route := range routes {
		if route.IsFallback {
			_fallbacks = append(_fallbacks, route)
		} else {
			_routes = append(_routes, route)
		}
	}
	_routes = append(_routes, _fallbacks...)
	for _, route := range _routes {
		if route.Matches(request, includingMethod) {
			return route
		}
	}
	return nil
}

/**
 * Determine if any routes match on another HTTP verb.
 *
 * @param  \Illuminate\Http\Request  $request
 * @return array
 */
/**
  protected function checkForAlternateVerbs($request)
  {
      $methods = array_diff(Router::$verbs, [$request->getMethod()]);

      // Here we will spin through all verbs except for the current request verb and
      // check to see if any routes respond to them. If they do, we will return a
      // proper error response with the correct headers on the response string.
      $others = [];

      foreach ($methods as $method) {
          if (! is_null($this->matchAgainstRoutes($this->get($method), $request, false))) {
              $others[] = $method;
          }
      }

      return $others;
  }
*/

/**
 * Get a route (if necessary) that responds when other available methods are present.
 *
 * @param  \Illuminate\Http\Request  $request
 * @param  array  $methods
 * @return \Illuminate\Routing\Route
 *
 * @throws \Symfony\Component\HttpKernel\Exception\MethodNotAllowedHttpException
 */
/**
  protected function getRouteForMethods($request, array $methods)
  {
      if ($request->method() === 'OPTIONS') {
          return (new Route('OPTIONS', $request->path(), function () use ($methods) {
              return new Response('', 200, ['Allow' => implode(',', $methods)]);
          }))->bind($request);
      }

      $this->methodNotAllowed($methods);
  }
*/

/**
 * Get routes from the collection by method.
 *
 * @param  string  method
 * @return map[string]*Route
 */
func (this *RouteCollection) Get(method ...string) map[string]*Route {
	if len(method) > 0 {
		if routes, ok := this.routes[method[0]]; ok {
			return routes
		} else {
			return nil
		}
	}
	return this.GetRoutes()
}

/**
 * Determine if the route collection contains a given named route.
 *
 * @param  string  name
 * @return bool
 */
func (this *RouteCollection) HasNamedRoute(name string) bool {
	_, ok := this.nameList[name]
	return ok
}

/**
 * Get a route instance by its name.
 *
 * @param  string  name
 * @return Routing\Route|nil
 */
func (this *RouteCollection) GetByName(name string) *Route {
	if route, ok := this.nameList[name]; ok {
		return route
	}
	return nil
}

/**
 * Get all of the routes in the collection.
 *
 * @return map[string]*Route
 */
func (this *RouteCollection) GetRoutes() map[string]*Route {
	return this.allRoutes
}

/**
 * Get all of the routes keyed by their HTTP verb / method.
 *
 * @return map[string]*Route
 */
func (this *RouteCollection) GetRoutesByMethod() map[string]map[string]*Route {
	return this.routes
}

/**
 * Get all of the routes keyed by their name.
 *
 * @return map[string]*Route
 */
func (this *RouteCollection) GetRoutesByName() map[string]*Route {
	return this.nameList
}

/**
 * Count the number of items in the collection.
 *
 * @return int
 */
func (this *RouteCollection) Count() int {
	return len(this.GetRoutes())
}
