package Routing

import (
	"fmt"
	"github.com/larisgo/framework/Http"
	"strings"
)

type Action func()

type Args map[string]string

func (args Args) Get(key string) string {
	if value, ok := args[key]; ok {
		return value
	}
	return ""
}

var Verbs map[string]bool = map[string]bool{"GET": true, "HEAD": true, "POST": true, "PUT": true, "PATCH": true, "DELETE": true, "OPTIONS": true}

type Router struct {
	routes *RouteCollection
}

func NewRouter() (this *Router) {
	this = &Router{}
	this.routes = NewRouteCollection()
	return this
}

func (this *Router) Get(uri string, action func()) *Route {
	return this.AddRoute(map[string]bool{"GET": true, "HEAD": true}, uri, Action(action))
}

func (this *Router) Post(uri string, action func()) *Route {
	return this.AddRoute(map[string]bool{"POST": true}, uri, Action(action))
}

func (this *Router) Put(uri string, action func()) *Route {
	return this.AddRoute(map[string]bool{"PUT": true}, uri, Action(action))
}

func (this *Router) Patch(uri string, action func()) *Route {
	return this.AddRoute(map[string]bool{"PATCH": true}, uri, Action(action))
}

func (this *Router) Delete(uri string, action func()) *Route {
	return this.AddRoute(map[string]bool{"DELETE": true}, uri, Action(action))
}

func (this *Router) Options(uri string, action func()) *Route {
	return this.AddRoute(map[string]bool{"OPTIONS": true}, uri, Action(action))
}

func (this *Router) Head(uri string, action func()) *Route {
	return this.AddRoute(map[string]bool{"HEAD": true}, uri, Action(action))
}

func (this *Router) Any(uri string, action func()) *Route {
	return this.AddRoute(Verbs, uri, Action(action))
}

func (this *Router) Match(methods []string, uri string, action func()) *Route {
	var _methods map[string]bool
	for _, v := range methods {
		_methods[strings.ToUpper(v)] = true
	}
	return this.AddRoute(_methods, uri, action)
}

func (this *Router) AddRoute(methods map[string]bool, uri string, action func()) *Route {
	return this.routes.Add(this.createRoute(methods, uri, Action(action)))
}

func (this *Router) createRoute(methods map[string]bool, uri string, action Action) *Route {
	return NewRoute(methods, this.prefix(uri), action).SetRouter(this)
}

/**
 * Prefix the given URI with the last prefix.
 *
 * @param  string  uri
 * @return string
 */
func (this *Router) prefix(uri string) string {
	if prefix := strings.Trim(uri, "/"); prefix != "" {
		return prefix
	}
	return "/"
}

func (this *Router) GetRoutes() *RouteCollection {
	return this.routes
}

/**
 * Dispatch the request to the application.
 *
 * @param  \Illuminate\Http\Request  $request
 * @return \Illuminate\Http\Response|\Illuminate\Http\JsonResponse
 */
func (this *Router) Dispatch(request *Http.Request) {
	fmt.Printf("%+v\n", request.Context)
	fmt.Printf("%+v\n", request.GetMethod())
	// $this->currentRequest = $request;

	// return $this->dispatchToRoute($request);
}

/**
 * Dispatch the request to a route and return the response.
 *
 * @param  \Illuminate\Http\Request  $request
 * @return mixed
 */
/**
  public function dispatchToRoute(Request $request)
  {
      return $this->runRoute($request, $this->findRoute($request));
  }
*/

/**
 * Find the route matching a given request.
 *
 * @param  \Illuminate\Http\Request  $request
 * @return \Illuminate\Routing\Route
 */
/**
  protected function findRoute($request)
  {
      $this->current = $route = $this->routes->match($request);

      $this->container->instance(Route::class, $route);

      return $route;
  }
*/
