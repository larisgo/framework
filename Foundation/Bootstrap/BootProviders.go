package Bootstrap

import (
	"github.com/larisgo/framework/Contracts/Foundation"
)

type BootProviders struct {
}

func (this *BootProviders) Bootstrap(app Foundation.Application) {
	app.Boot()
}
