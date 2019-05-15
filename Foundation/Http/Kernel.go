package Http

import (
	"github.com/larisgo/larisgo/Contracts/Http"
	"github.com/larisgo/larisgo/Foundation"
)

type Kernel struct {
	app *Foundation.Application

	Http.Kernel
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

func NewKernel(app *Foundation.Application) (kernel *Kernel) {
	kernel = &Kernel{}
	kernel.app = app

	return kernel
}
