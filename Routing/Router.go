package Routing

import (
	"fmt"
	"github.com/larisgo/framework/Contracts/Container"
	"github.com/larisgo/framework/Http"
	"strings"
)

type Action func(*Http.Request) *Http.Response

/**
 * All of the verbs supported by the router.
 *
 * @var map[string]bool
 */
var Verbs map[string]bool = map[string]bool{"GET": true, "HEAD": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true, "OPTIONS": true}

type Router struct {
	routes  *RouteCollection
	current *Route

	container Container.Container

	middleware       map[string]interface{}
	middlewareGroups map[string]interface{}

	/**
	 * The priority-sorted list of middleware.
	 *
	 * Forces the listed middleware to always be in the given order.
	 *
	 * @var array
	 */
	MiddlewarePriority map[string]interface{}
	/**
	 * The registered route value binders.
	 *
	 * @var array
	 */
	binders map[string]interface{}
	/**
	 * The globally available parameter patterns.
	 *
	 * @var array
	 */
	patterns map[string]string

	groupStack []map[string]string
}

func NewRouter(container Container.Container) (this *Router) {
	this = &Router{}
	this.container = container
	this.routes = NewRouteCollection()
	this.patterns = map[string]string{}
	this.groupStack = []map[string]string{}
	return this
}

func (this *Router) Get(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(map[string]bool{"GET": true, "HEAD": true}, uri, Action(action))
}

func (this *Router) Post(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(map[string]bool{"POST": true}, uri, Action(action))
}

func (this *Router) Put(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(map[string]bool{"PUT": true}, uri, Action(action))
}

func (this *Router) Patch(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(map[string]bool{"PATCH": true}, uri, Action(action))
}

func (this *Router) Delete(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(map[string]bool{"DELETE": true}, uri, Action(action))
}

func (this *Router) Options(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(map[string]bool{"OPTIONS": true}, uri, Action(action))
}

func (this *Router) Head(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(map[string]bool{"HEAD": true}, uri, Action(action))
}

func (this *Router) Any(uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.AddRoute(Verbs, uri, Action(action))
}

/**
 * Register a new Fallback route with the router.
 *
 * @param  \Closure|array|string|callable|null  action
 * @return \Illuminate\Routing\Route
 */
func (this *Router) Fallback(action func(*Http.Request) *Http.Response, methods ...string) *Route {
	placeholder := "fallbackPlaceholder"
	if len(methods) == 0 {
		methods = []string{"GET"}
	}
	return this.Match(methods, fmt.Sprintf("{{%s}}", placeholder), action).Where(placeholder, ".*").Fallback()
}

func (this *Router) Match(methods []string, uri string, action func(*Http.Request) *Http.Response) *Route {
	var _methods map[string]bool
	for _, v := range methods {
		_methods[strings.ToUpper(v)] = true
	}
	return this.AddRoute(_methods, uri, action)
}

func (this *Router) AddRoute(methods map[string]bool, uri string, action func(*Http.Request) *Http.Response) *Route {
	return this.routes.Add(this.createRoute(methods, uri, Action(action)))
}

func (this *Router) createRoute(methods map[string]bool, uri string, action Action) *Route {
	return NewRoute(methods, this.prefix(uri), action).SetRouter(this)
}

/**
 * Create a route group with shared attributes.
 *
 * @param  array  attributes
 * @param  \Closure|string  routes
 * @return void
 */
func (this *Router) Group(attributes map[string]string, routes func(*Router)) {
	this.updateGroupStack(attributes)

	// Once we have updated the group stack, we'll load the provided routes and
	// merge in the group's attributes when the routes are created. After we
	// have created the routes, we will pop the attributes off the stack.
	this.loadRoutes(routes)

	this.groupStack = this.groupStack[:len(this.groupStack)-1]
}

/**
 * Update the group stack with the given attributes.
 *
 * @param  map[string]string  attributes
 * @return void
 */
func (this *Router) updateGroupStack(attributes map[string]string) {
	if len(this.groupStack) > 0 {
		attributes = this.MergeWithLastGroup(attributes)
	}
	this.groupStack = append(this.groupStack, attributes)
}

/**
 * Merge the given array with the last group stack.
 *
 * @param  array  new
 * @return array
 */
func (this *Router) MergeWithLastGroup(_new map[string]string) map[string]string {
	return NewRouteGroup().Merge(_new, this.groupStack[len(this.groupStack)-1])
}

/**
 * Load the provided routes.
 *
 * @param  \Closure|string  routes
 * @return void
 */
func (this *Router) loadRoutes(routes func(*Router)) {
	routes(this)
}

/**
 * Get the prefix from the last group on the stack.
 *
 * @return string
 */
func (this *Router) GetLastGroupPrefix() string {
	if len(this.groupStack) > 0 {
		last := this.groupStack[len(this.groupStack)-1]
		if v, ok := last["prefix"]; ok {
			return v
		}
	}
	return ""
}

/**
 * Prefix the given URI with the last prefix.
 *
 * @param  string  uri
 * @return string
 */
func (this *Router) prefix(uri string) string {
	if prefix := strings.Trim(strings.Trim(this.GetLastGroupPrefix(), "/")+"/"+strings.Trim(uri, "/"), "/"); prefix != "" {
		return prefix
	}
	return "/"
}

/**
 * Return the response returned by the given route.
 *
 * @param  string  $name
 * @return mixed
 */
func (this *Router) RespondWithRoute(request *Http.Request, name string) interface{} {
	route := this.routes.GetByName(name).Bind(request)

	return this.runRoute(request, route)
}

func (this *Router) GetRoutes() *RouteCollection {
	return this.routes
}

/**
 * Dispatch the request to the application.
 *
 * @param  \Illuminate\Http\Request  request
 * @return \Illuminate\Http\Response|\Illuminate\Http\JsonResponse
 */
func (this *Router) Dispatch(request *Http.Request) *Http.Response {

	return this.DispatchToRoute(request)
}

/**
 * Dispatch the request to a route and return the response.
 *
 * @param  \Illuminate\Http\Request  request
 * @return mixed
 */
func (this *Router) DispatchToRoute(request *Http.Request) *Http.Response {
	return this.runRoute(request, this.findRoute(request))
}

/**
 * Find the route matching a given request.
 *
 * @param  \Illuminate\Http\Request  request
 * @return \Illuminate\Routing\Route
 */
func (this *Router) findRoute(request *Http.Request) *Route {
	route := this.routes.Match(request)
	this.current = route
	// this.container.instance(Route::class, route);

	return route
}

/**
 * Return the response for the given route.
 *
 * @param  \Illuminate\Http\Request  $request
 * @param  \Illuminate\Routing\Route  $route
 * @return mixed
 */
func (this *Router) runRoute(request *Http.Request, route *Route) *Http.Response {
	// request.SetRouteResolver(func() interface{} {
	// 	return route
	// })
	// $this->events->dispatch(new Events\RouteMatched($route, $request));

	return this.PrepareResponse(request, this.runRouteWithinStack(route, request))
}

/**
 * Run the given route within a Stack "onion" instance.
 *
 * @param  \Illuminate\Routing\Route  $route
 * @param  \Illuminate\Http\Request  $request
 * @return mixed
 */
func (this *Router) runRouteWithinStack(route *Route, request *Http.Request) *Http.Response {
	// $shouldSkipMiddleware = $this->container->bound('middleware.disable') &&
	//                         $this->container->make('middleware.disable') === true;

	// $middleware = $shouldSkipMiddleware ? [] : $this->gatherRouteMiddleware($route);

	return this.PrepareResponse(request, route.Run(request))
}

/**
 * Create a response instance from the given value.
 *
 * @param  \Symfony\Component\HttpFoundation\Request  $request
 * @param  mixed  $response
 * @return \Illuminate\Http\Response|\Illuminate\Http\JsonResponse
 */
func (this *Router) PrepareResponse(request *Http.Request, response *Http.Response) *Http.Response {
	return this.ToResponse(request, response)
}

/**
 * Static version of prepareResponse.
 *
 * @param  \Symfony\Component\HttpFoundation\Request  $request
 * @param  mixed  $response
 * @return \Illuminate\Http\Response|\Illuminate\Http\JsonResponse
 */
func (this *Router) ToResponse(request *Http.Request, response *Http.Response) *Http.Response {
	return response.Prepare(request)
}
