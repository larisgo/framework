package Support

import (
	"github.com/larisgo/framework/Contracts/Foundation"
)

type ServiceProvider struct {
	App Foundation.Application
}

func NewServiceProvider(app Foundation.Application) *ServiceProvider {
	return &ServiceProvider{
		App: app,
	}
}
