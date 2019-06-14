package Providers

import (
	"github.com/larisgo/framework/Contracts/Container"
	"github.com/larisgo/framework/Contracts/Foundation"
	"github.com/larisgo/framework/Routing"
	"github.com/larisgo/framework/Support"
)

type RoutingServiceProvider struct {
	*Support.ServiceProvider
}

func NewRoutingServiceProvider(app Foundation.Application) (this *RoutingServiceProvider) {
	this = &RoutingServiceProvider{ServiceProvider: Support.NewServiceProvider(app)}
	return this
}

func (this *RoutingServiceProvider) Register() {
	this.registerRouter()
}

/**
 * Register the router instance.
 *
 * @return void
 */
func (this *RoutingServiceProvider) registerRouter() {
	this.App.Singleton("router", func(app Container.Container) interface{} {
		return Routing.NewRouter(app)
	})
}
