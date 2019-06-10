package Http

import (
	"fmt"
	// "github.com/larisgo/framework/Errors"
	"github.com/larisgo/framework/Foundation"
	"github.com/larisgo/framework/Http"
	"github.com/larisgo/framework/Routing"
	"net/http"
	"runtime"
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
	panic(http.ListenAndServe("127.0.0.1:8000", this))
	// http.ListenAndServeTLS(addr, certFile, keyFile, this)
}

func (this *Kernel) SendRequestThroughRouter(request *Http.Request) {
	this.router.Dispatch(request).Send()
	// request.Error(fasthttp.StatusMessage(fasthttp.StatusNotFound),
	// fasthttp.StatusNotFound)
}

/**
 * Get the route dispatcher callback.
 *
 * @return \Closure
 */
// func (this *Kernel) dispatchToRouter() {
// 	return func(request *Http.Request) {
// 		// $this->app->instance('request', $request);
// 		return this.router.Dispatch(request)
// 	}
// }

/**
 * Get the route dispatcher callback.
 *
 * @return \Closure
 */
// func (this *Kernel) dispatchToRouter() *Routing.Router {
// }

func (this *Kernel) ServeFile(uri string, root string) {
	// if len(uri) < 10 || uri[len(uri)-10:] != "/*fileuri" {
	// 	panic("uri must end with /*fileuri in uri '" + uri + "'")
	// }

	// fileServer := http.FileServer(root)

	// this.GET(uri, func(w http.ResponseWriter, req *http.Request, ps Params) {
	// 	req.URL.Path = ps.ByName("fileuri")
	// 	fileServer.ServeHTTP(w, req)
	// })
}

func (this *Kernel) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			// if e, ok := err.(Errors.Exception); ok {
			// 	fmt.Printf("%+v\n%s\n", e.GetMessage(), string(buf[:n]))
			// 	fmt.Fprintf(response, "%+v\n%s\n", e.GetMessage(), string(buf[:n]))
			// } else {
			fmt.Printf("%+v\n%s\n", err, string(buf[:n]))
			fmt.Fprintf(response, "%+v\n%s\n", err, string(buf[:n]))
			// }
		}
	}()
	this.SendRequestThroughRouter(Http.NewRequest(this.app, response, request))
}

/**
 * Call the terminate method on any terminable middleware.
 *
 * @param  Http\Request  request
 * @param  Http\Response  response
 * @return void
 */
func (this *Kernel) Terminate(request *Http.Request, response *Http.Response) {
	this.app.Terminate()
}

func (this *Kernel) GetApplication() *Foundation.Application {
	return this.app
}
