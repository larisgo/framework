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

func (this *Kernel) SendRequestThroughRouter(request *fasthttp.RequestCtx) {
	this.router.Dispatch(Http.Capture(request)).Send()
	// this.dispatchToRouter()()
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
