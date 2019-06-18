package Support

import (
	"github.com/larisgo/framework/Contracts/Foundation"
)

type ServiceProvider struct {
	App   Foundation.Application `inject:"app"`
	Defer bool                   `default:false`
}

func NewServiceProvider(app Foundation.Application) *ServiceProvider {
	return &ServiceProvider{
		App: app,
	}
}

/**
 * Determine if the provider is deferred.
 *
 * @return bool
 */
func (this *ServiceProvider) IsDeferred() bool {
	return this.Defer
}
