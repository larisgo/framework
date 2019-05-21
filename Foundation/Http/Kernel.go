package Http

import (
	"github.com/larisgo/larisgo/Contracts/Http"
	"github.com/larisgo/larisgo/Foundation"
	"github.com/larisgo/larisgo/Routing"
)

type Kernel struct {
	app    *Foundation.Application
	router *Routing.Router
	Http.Kernel
}

func NewKernel(app *Foundation.Application, router *Routing.Router) (this *Kernel) {
	this = &Kernel{}
	this.app = app
	this.router = router

	return this
}

func (this *Kernel) Handle() {

}

func (this *Kernel) SendRequestThroughRouter() *Routing.Router {
	return this.router
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
