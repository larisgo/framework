package Routing

import (
	"net/http"
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

type Router struct {
	routes                *RouteCollection
	RedirectTrailingSlash bool
	RedirectFixedPath     bool
	NotFound              http.Handler
	MethodNotAllowed      http.Handler
	PanicHandler          func(http.ResponseWriter, *http.Request, interface{})
}

func NewRouter() (this *Router) {
	this = &Router{
		RedirectTrailingSlash: true,
		RedirectFixedPath:     true,
	}
	this.routes = NewRouteCollection()
	return this
}

var _ http.Handler = NewRouter()

func (this *Router) Get(uri string, action func()) *Route {
	return this.AddRoute([]string{"GET", "HEAD"}, uri, Action(action))
}

func (this *Router) Post(uri string, action func()) *Route {
	return this.AddRoute([]string{"POST"}, uri, Action(action))
}

func (this *Router) Put(uri string, action func()) *Route {
	return this.AddRoute([]string{"PUT"}, uri, Action(action))
}

func (this *Router) Patch(uri string, action func()) *Route {
	return this.AddRoute([]string{"PATCH"}, uri, Action(action))
}

func (this *Router) Delete(uri string, action func()) *Route {
	return this.AddRoute([]string{"DELETE"}, uri, Action(action))
}

func (this *Router) Options(uri string, action func()) *Route {
	return this.AddRoute([]string{"OPTIONS"}, uri, Action(action))
}

func (this *Router) Head(uri string, action func()) *Route {
	return this.AddRoute([]string{"HEAD"}, uri, Action(action))
}

func (this *Router) Any(uri string, action func()) *Route {
	return this.AddRoute([]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, uri, Action(action))
}

func (this *Router) Match(methods []string, uri string, action func()) *Route {
	return this.AddRoute(methods, uri, action)
}

func (this *Router) AddRoute(methods []string, uri string, action func()) *Route {
	return this.routes.Add(this.createRoute(methods, uri, Action(action)))
}

func (this *Router) createRoute(methods []string, uri string, action Action) *Route {
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

// HandlerFunc is an adapter which allows the usage of an http.HandlerFunc as a
// request handle.
// func (this *Router) HandlerFunc(method, uri string, handler http.HandlerFunc) {
// 	this.AddRoute(method, uri, handler)
// }

func (this *Router) ServeFiles(uri string, root http.FileSystem) {
	// if len(uri) < 10 || uri[len(uri)-10:] != "/*fileuri" {
	// 	panic("uri must end with /*fileuri in uri '" + uri + "'")
	// }

	// fileServer := http.FileServer(root)

	// this.GET(uri, func(w http.ResponseWriter, req *http.Request, ps Params) {
	// 	req.URL.Path = ps.ByName("fileuri")
	// 	fileServer.ServeHTTP(w, req)
	// })
}

func (this *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// if this.PanicHandler != nil {
	// 	// defer this.recv(w, req)
	// }

	// uri := req.URL.Path

	// if root := this.trees[req.Method]; root != nil {
	// 	if handle, ps, tsr := root.getValue(uri); handle != nil {
	// 		handle(w, req, ps)
	// 		return
	// 	} else if req.Method != "CONNECT" && uri != "/" {
	// 		code := 301 // Permanent redirect, request with GET method
	// 		if req.Method != "GET" {
	// 			// Temporary redirect, request with same method
	// 			// As of Go 1.3, Go does not support status code 308.
	// 			code = 307
	// 		}

	// 		if tsr && this.RedirectTrailingSlash {
	// 			if len(uri) > 1 && uri[len(uri)-1] == '/' {
	// 				req.URL.Path = uri[:len(uri)-1]
	// 			} else {
	// 				req.URL.Path = uri + "/"
	// 			}
	// 			http.Redirect(w, req, req.URL.String(), code)
	// 			return
	// 		}

	// 		// Try to fix the request uri
	// 		if this.RedirectFixedPath {
	// 			fixedPath, found := root.findCaseInsensitivePath(
	// 				CleanPath(uri),
	// 				this.RedirectTrailingSlash,
	// 			)
	// 			if found {
	// 				req.URL.Path = string(fixedPath)
	// 				http.Redirect(w, req, req.URL.String(), code)
	// 				return
	// 			}
	// 		}
	// 	}
	// }

	// if req.Method == "OPTIONS" {
	// 	// // Handle OPTIONS requests
	// 	// if allow := this.allowed(uri, req.Method); len(allow) > 0 {
	// 	// 	w.Header().Set("Allow", allow)
	// 	// 	return
	// 	// }
	// }

	// // Handle 404
	// if this.NotFound != nil {
	// 	this.NotFound.ServeHTTP(w, req)
	// } else {
	http.NotFound(w, req)
	// }
}
