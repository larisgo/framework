package Routing

type SchemeValidator struct {
}

func NewSchemeValidator() ValidatorInterface {
	return SchemeValidator{}
}

func (this SchemeValidator) matches(route *Route, method string) bool {
	if route.HttpOnly() {
		// return ! $request->secure();
	} else if route.Secure() {
		// return $request->secure();
	}
	return true
}
