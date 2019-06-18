package Service

type RegisterT interface {
	Register()
}

type BootT interface {
	Boot()
}

type BindingsT interface {
	Bindings() map[string]interface{}
}

type SingletonsT interface {
	Singletons() map[string]interface{}
}

type IsDeferredT interface {
	IsDeferred() bool
}

type Provider interface {
	RegisterT
	BootT
	BindingsT
	SingletonsT
	IsDeferredT
}
