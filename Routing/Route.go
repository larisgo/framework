package Routing

import (
	"github.com/larisgo/larisgo/Support"
	"strings"
)

var Validators []ValidatorInterface

type Route struct {
	uri        string
	methods    map[string]bool
	Action     *routeAction
	IsFallback bool
	compiled   *CompiledRoute
	router     *Router
	http       bool
	https      bool
}

func NewRoute(methods map[string]bool, uri string, action Action) (this *Route) {
	this = &Route{}
	this.uri = uri
	this.methods = methods
	this.Action = this.parseAction(action)
	return this
}

/**
 * Parse the route action into a standard array.
 *
 * @param  Action  action
 * @return *routeAction
 */
func (this *Route) parseAction(action Action) *routeAction {
	return RouteAction().Parse(this.uri, action)
}

/**
 * Run the route action and return the response.
 *
 * @return mixed
 */
func (this *Route) Run() {
	// $this->container = $this->container ?: new Container;
	this.runCallable()
	return
}

/**
 * Run the route action and return the response.
 *
 * @return mixed
 */
func (this *Route) runCallable() {
	this.Action.Uses()
	return
}

/**
 * Determine if the route matches given request.
 *
 * @param  Http\Request  $request
 * @param  bool  $includingMethod
 * @return bool
 */
func (this *Route) Matches(request string, includingMethod bool) bool {
	this.compileRoute()

	for _, validator := range this.GetValidators() {
		if _, ok := validator.(MethodValidator); !includingMethod && ok {
			continue
		}
		if !validator.matches(this, request) {
			return false
		}
	}

	return true
}

/**
 * Compile the route into a Symfony CompiledRoute instance.
 *
 * @return \Symfony\Component\Routing\CompiledRoute
 */
func (this *Route) compileRoute() *CompiledRoute {
	if this.compiled == nil {
		this.compiled = NewRouteCompiler(this).Compile()
	}

	return this.compiled
}

/**
 * Bind the route to a given request for execution.
 *
 * @param  \Illuminate\Http\Request  $request
 * @return $this
 */
// func (this *Route) compileRoute(request string) *Route {
// this.compileRoute()

// $this->parameters = (new RouteParameterBinder($this))
// ->parameters($request);
//
// $this->originalParameters = $this->parameters;

// return this
// }

/**
 * Set the router instance on the route.
 *
 * @param  \Illuminate\Routing\Router  $router
 * @return $this
 */
func (this *Route) SetRouter(router *Router) *Route {
	this.router = router
	return this
}

/**
 * Mark this route as a fallback route.
 *
 * @return this
 */
func (this *Route) Fallback() *Route {
	this.IsFallback = true
	return this
}

/**
 * Get the HTTP verbs the route responds to.
 *
 * @return array
 */
func (this *Route) Methods() map[string]bool {
	return this.methods
}

/**
 * Determine if the route only responds to HTTP requests.
 *
 * @return bool
 */
func (this *Route) HttpOnly() bool {
	return this.http
}

/**
 * Determine if the route only responds to HTTPS requests.
 *
 * @return bool
 */
func (this *Route) HttpsOnly() bool {
	return this.Secure()
}

/**
 * Determine if the route only responds to HTTPS requests.
 *
 * @return bool
 */
func (this *Route) Secure() bool {
	return this.https
}

/**
 * Get or set the domain for the route.
 *
 * @param  string  domain
 * @return this
 */
func (this *Route) Domain(_domain string) *Route {
	this.Action.Domain = _domain
	return this
}

/**
 * Get the domain defined for the route.
 *
 * @return string
 */
func (this *Route) GetDomain() string {
	if this.Action.Domain != "" {
		return strings.Replace(strings.Replace(this.Action.Domain, "http://", "", 1), "https://", "", 1)
	}
	return ""
}

/**
 * Get the prefix of the route instance.
 *
 * @return string
 */
func (this *Route) GetPrefix() string {
	return this.Action.Prefix
}

/**
 * Add a prefix to the route URI.
 *
 * @param  string  prefix
 * @return $this
 */
func (this *Route) Prefix(prefix string) *Route {
	this.uri = strings.Trim(strings.TrimRight(prefix, "/")+"/"+strings.TrimLeft(this.uri, "/"), "/")
	return this
}

/**
 * Get the URI associated with the route.
 *
 * @return string
 */
func (this *Route) Uri() string {
	return this.uri
}

/**
 * Set the URI that the route responds to.
 *
 * @param  string  uri
 * @return $this
 */
func (this *Route) SetUri(uri string) *Route {
	this.uri = uri
	return this
}

/**
 * Get the name of the route instance.
 *
 * @return string
 */
func (this *Route) GetName() string {
	return this.Action.As
}

/**
 * Add or change the route name.
 *
 * @param  string  $name
 * @return $this
 */
func (this *Route) Name(name string) *Route {
	this.Action.As = this.Action.As + name
	return this
}

/**
 * Determine whether the route's name matches the given patterns.
 *
 * @param  mixed  ...patterns
 * @return bool
 */
func (this *Route) Named(patterns ...string) bool {
	var routeName string

	if routeName = this.GetName(); routeName == "" {
		return false
	}

	for _, pattern := range patterns {
		if Support.Str(routeName).Is([]string{pattern}) {
			return true
		}
	}

	return false
}

/**
 * Set the handler for the route.
 *
 * @param  \Closure|string  $action
 * @return $this
 */
func (this *Route) Uses(action func()) *Route {
	this.Action.Uses = Action(action)
	return this
}

/**
 * Set the action array for the route.
 *
 * @param  array  $action
 * @return $this
 */
func (this *Route) SetAction(action *routeAction) *Route {
	this.Action = action

	return this
}

/**
 * Get or set the middlewares attached to the route.
 *
 * @param  string middleware
 * @return this
 */
func (this *Route) Middleware(middleware ...string) *Route {
	if len(middleware) > 0 {
		this.Action.Middleware = append(this.Action.Middleware, middleware...)
	}
	return this
}

/**
 * Get the compiled version of the route.
 *
 * @return \Symfony\Component\Routing\CompiledRoute
 */
func (this *Route) GetCompiled() *CompiledRoute {
	return this.compiled
}

/**
 * Get the route validators for the instance.
 *
 * @return array
 */
func (this *Route) GetValidators() []ValidatorInterface {
	if Validators != nil {
		return Validators
	}

	// To match the route, we will use a chain of responsibility pattern with the
	// validator implementations. We will spin through each one making sure it
	// passes and then we will know if the route as a whole matches request.
	Validators = []ValidatorInterface{
		NewUriValidator(), NewMethodValidator(),
		NewSchemeValidator(), NewHostValidator(),
	}
	return Validators
}
