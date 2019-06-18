package Facades

import (
	"github.com/larisgo/framework/Contracts/Foundation"
)

var App Foundation.Application
var resolvedInstance map[string]interface{}

type Facade struct {
	facadeaccessor string
}

func NewFacade(facadeaccessor string) *Facade {
	return &Facade{facadeaccessor: facadeaccessor}
}

/**
 * Resolve the facade root instance from the container.
 *
 * @param  string|object  name
 * @return mixed
 */
func (this *Facade) resolveFacadeInstance(name string) interface{} {
	if v, ok := resolvedInstance[name]; ok {
		return v
	}
	resolvedInstance[name] = App.Make(this.facadeaccessor)
	return resolvedInstance[name]
}

/**
 * Clear a resolved facade instance.
 *
 * @param  string  name
 * @return void
 */
func ClearResolvedInstance(name string) {
	delete(resolvedInstance, name)
}

/**
 * Clear all of the resolved instances.
 *
 * @return void
 */
func ClearResolvedInstances() {
	resolvedInstance = map[string]interface{}{}
}

/**
 * Get the application instance behind the facade.
 *
 * @return \Illuminate\Contracts\Foundation\Application
 */
func GetFacadeApplication() Foundation.Application {
	return App
}

/**
 * Set the application instance.
 *
 * @param  \Illuminate\Contracts\Foundation\Application  app
 * @return void
 */
func SetFacadeApplication(app Foundation.Application) {
	App = app
}

func (this *Facade) Get() interface{} {
	return this.resolveFacadeInstance(this.facadeaccessor)
}
