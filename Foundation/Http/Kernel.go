package Http

import (
	"fmt"
	// "github.com/larisgo/framework/Errors"
	FoundationContract "github.com/larisgo/framework/Contracts/Foundation"
	"github.com/larisgo/framework/Foundation"
	"github.com/larisgo/framework/Foundation/Bootstrap"
	"github.com/larisgo/framework/Http"
	"github.com/larisgo/framework/Routing"
	"net/http"
	"runtime"
)

type Kernel struct {
	App           *Foundation.Application `inject:"app"`
	Router        *Routing.Router         `inject:"router"`
	bootstrappers []FoundationContract.BootstrapT
}

func NewKernel() (this *Kernel) {
	this = &Kernel{}

	this.bootstrappers = []FoundationContract.BootstrapT{
		// &Bootstrap.LoadEnvironmentVariables{},
		&Bootstrap.LoadConfiguration{},
		// &Bootstrap.HandleExceptions{},
		&Bootstrap.RegisterFacades{},
		&Bootstrap.RegisterProviders{},
		&Bootstrap.BootProviders{},
	}

	return this
}

func (this *Kernel) Bootstrap() {
	if !this.App.HasBeenBootstrapped() {
		this.App.BootstrapWith(this.bootstrappers)
	}
}

func (this *Kernel) Handle() {
	this.Bootstrap()

	fmt.Println(`Larisgo development server started: <http://127.0.0.1:8000>`)
	panic(http.ListenAndServe("127.0.0.1:8000", this))
	// http.ListenAndServeTLS(addr, certFile, keyFile, this)
}

func (this *Kernel) SendRequestThroughRouter(request *Http.Request) {
	this.Router.Dispatch(request).Send()
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
// 		// this.App.instance('request', request);
// 		return this.Router.Dispatch(request)
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
			// fmt.Printf("%+v\n%s\n", err, string(buf[:n]))
			fmt.Fprintf(response, "%+v\n%s\n", err, string(buf[:n]))
			// }
		}
	}()
	this.SendRequestThroughRouter(Http.NewRequest(this.App, response, request))
}

/**
 * Call the terminate method on any terminable middleware.
 *
 * @param  Http\Request  request
 * @param  Http\Response  response
 * @return void
 */
func (this *Kernel) Terminate() {
	this.App.Terminate()
}

func (this *Kernel) GetApplication() *Foundation.Application {
	return this.App
}
