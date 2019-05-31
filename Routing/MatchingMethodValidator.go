package Routing

type MethodValidator struct {
}

func NewMethodValidator() ValidatorInterface {
	return MethodValidator{}
}

func (this MethodValidator) matches(route *Route, method string) bool {
	_, ok := route.Methods()[method]
	return ok
}
