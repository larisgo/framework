package Http

import (
	"github.com/larisgo/framework/Foundation"
	"github.com/larisgo/framework/Http"
	"github.com/larisgo/framework/Routing"
	"github.com/valyala/fasthttp"
)

type Kernel struct {
	app    *Foundation.Application
	router *Routing.Router
}

func NewKernel(app *Foundation.Application, router *Routing.Router) (this *Kernel) {
	this = &Kernel{}
	this.app = app
	this.router = router

	return this
}

func (this *Kernel) Bootstrap() {
}

func (this *Kernel) Handle() {
}

func (this *Kernel) SendRequestThroughRouter() *Routing.Router {
	return this.router
}

/**
 * Get the route dispatcher callback.
 *
 * @return \Closure
 */
// func (this *Kernel) dispatchToRouter() *Routing.Router {
// }

func (this *Kernel) ServeFiles(uri string, root string) {
	// if len(uri) < 10 || uri[len(uri)-10:] != "/*fileuri" {
	// 	panic("uri must end with /*fileuri in uri '" + uri + "'")
	// }

	// fileServer := http.FileServer(root)

	// this.GET(uri, func(w http.ResponseWriter, req *http.Request, ps Params) {
	// 	req.URL.Path = ps.ByName("fileuri")
	// 	fileServer.ServeHTTP(w, req)
	// })
}

// Handler makes the router implement the fasthttp.ListenAndServe interface.
func (this *Kernel) Handler(request *fasthttp.RequestCtx) {
	this.router.Dispatch(Http.Capture(request))
	request.Error(fasthttp.StatusMessage(fasthttp.StatusNotFound),
		fasthttp.StatusNotFound)
}

// func (this *Kernel) ServeHTTP(response http.ResponseWriter, request *http.Request) {
// 	this.router.Dispatch(Http.Capture(request))
// 	// if this.PanicHandler != nil {
// 	// 	// defer this.recv(w, req)
// 	// }

// 	// uri := req.URL.Path

// 	// if root := this.trees[req.Method]; root != nil {
// 	// 	if handle, ps, tsr := root.getValue(uri); handle != nil {
// 	// 		handle(w, req, ps)
// 	// 		return
// 	// 	} else if req.Method != "CONNECT" && uri != "/" {
// 	// 		code := 301 // Permanent redirect, request with GET method
// 	// 		if req.Method != "GET" {
// 	// 			// Temporary redirect, request with same method
// 	// 			// As of Go 1.3, Go does not support status code 308.
// 	// 			code = 307
// 	// 		}

// 	// 		if tsr && this.RedirectTrailingSlash {
// 	// 			if len(uri) > 1 && uri[len(uri)-1] == '/' {
// 	// 				req.URL.Path = uri[:len(uri)-1]
// 	// 			} else {
// 	// 				req.URL.Path = uri + "/"
// 	// 			}
// 	// 			http.Redirect(w, req, req.URL.String(), code)
// 	// 			return
// 	// 		}

// 	// 		// Try to fix the request uri
// 	// 		if this.RedirectFixedPath {
// 	// 			fixedPath, found := root.findCaseInsensitivePath(
// 	// 				CleanPath(uri),
// 	// 				this.RedirectTrailingSlash,
// 	// 			)
// 	// 			if found {
// 	// 				req.URL.Path = string(fixedPath)
// 	// 				http.Redirect(w, req, req.URL.String(), code)
// 	// 				return
// 	// 			}
// 	// 		}
// 	// 	}
// 	// }

// 	// if req.Method == "OPTIONS" {
// 	// 	// // Handle OPTIONS requests
// 	// 	// if allow := this.allowed(uri, req.Method); len(allow) > 0 {
// 	// 	// 	w.Header().Set("Allow", allow)
// 	// 	// 	return
// 	// 	// }
// 	// }

// 	// // Handle 404
// 	// if this.NotFound != nil {
// 	// 	this.NotFound.ServeHTTP(w, req)
// 	// } else {
// 	http.NotFound(response, request)
// 	// }
// }

/**
 * Call the terminate method on any terminable middleware.
 *
 * @param  Http\Request  request
 * @param  Http\Response  response
 * @return void
 */
func (this *Kernel) Terminate(request interface{}, response interface{}) {
	this.app.Terminate()
}

func (this *Kernel) GetApplication() *Foundation.Application {
	return this.app
}
