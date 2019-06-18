package Providers

import (
	"github.com/larisgo/framework/Contracts/Foundation"
	"github.com/larisgo/framework/Routing"
	"github.com/larisgo/framework/Support"
	// "github.com/larisgo/framework/Support/Facades"
)

type RouteServiceProvider struct {
	*Support.ServiceProvider
}

func NewRouteServiceProvider(app Foundation.Application) (this *RouteServiceProvider) {
	this = &RouteServiceProvider{ServiceProvider: Support.NewServiceProvider(app)}
	return this
}

func (this *RouteServiceProvider) Register() {
	//
}

/**
 * Bootstrap any application services.
 *
 * @return void
 */
func (this *RouteServiceProvider) Boot() {
	this.App.Booted(func(app interface{}) {
		this.App.Make("router").(*Routing.Router).GetRoutes().RefreshNameLookups()
	})
}
