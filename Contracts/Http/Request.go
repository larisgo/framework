package Http

type Request interface {
	GetPathInfo() string
	Method() string
	Get(string, ...string) string
	Gets(string, ...[]string) []string
	GetMethod() string
	GetHost() string
	Secure() bool
	Path() string
	UserAgent() string
	IsJson() bool
}
