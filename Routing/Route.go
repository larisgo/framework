package Routing

type Route struct {
}

func NewRoute(methods []string, uri string, action Action) *Route {
	return &Route{}
}
