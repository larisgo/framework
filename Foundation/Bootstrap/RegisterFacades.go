package Bootstrap

import (
	"github.com/larisgo/framework/Contracts/Foundation"
	"github.com/larisgo/framework/Support/Facades"
)

type RegisterFacades struct {
}

func (this *RegisterFacades) Bootstrap(app Foundation.Application) {
	Facades.ClearResolvedInstances()
	Facades.SetFacadeApplication(app)
}
