package Routing

type RouteCollection struct {
	routes     []*Route
	allRoutes  []*Route
	nameList   []*Route
	actionList []*Route
}

func NewRouteCollection() (_RouteCollection *RouteCollection) {
	_RouteCollection = &RouteCollection{}
	_RouteCollection.routes = make([]*Route, 0)
	_RouteCollection.allRoutes = make([]*Route, 0)
	_RouteCollection.nameList = make([]*Route, 0)
	_RouteCollection.actionList = make([]*Route, 0)
	return _RouteCollection
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
	// $domainAndUri = $route->getDomain().$route->uri();

	// foreach ($route->methods() as $method) {
	//     $this->routes[$method][$domainAndUri] = $route;
	// }

	// $this->allRoutes[$method.$domainAndUri] = $route;
}

/**
 * Add the route to any look-up tables if necessary.
 *
 * @param  Routing\Route  route
 * @return void
 */
func (this *RouteCollection) addLookups(route *Route) {
	// // If the route has a name, we will add it to the name look-up table so that we
	// // will quickly be able to find any route associate with a name and not have
	// // to iterate through every route every time we need to perform a look-up.
	// if ($name = $route->getName()) {
	//     $this->nameList[$name] = $route;
	// }

	// // When the route is routing to a controller we will also store the action that
	// // is used by the route. This will let us reverse route to controllers while
	// // processing a request and easily generate URLs to the given controllers.
	// $action = $route->getAction();

	// if (isset($action['controller'])) {
	//     $this->addToActionList($action, $route);
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
	// $this->actionList[trim($action['controller'], '\\')] = $route;
}
