package Http

import (
	"github.com/larisgo/larisgo/Foundation/Application"
)

type Kernel struct {
	app *Application
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
func Kernel() *Kernel {
	return &Kernel{}
}
