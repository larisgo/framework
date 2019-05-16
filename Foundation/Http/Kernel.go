package Http

import (
	"github.com/larisgo/larisgo/Contracts/Http"
	"github.com/larisgo/larisgo/Foundation"
	"github.com/larisgo/larisgo/Routing"
)

type Kernel struct {
	app *Foundation.Application

	Http.Kernel
}

func NewKernel(app *Foundation.Application, router *Routing.Router) (_Kernel *Kernel) {
	_Kernel = &Kernel{}
	_Kernel.app = app

	return _Kernel
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
