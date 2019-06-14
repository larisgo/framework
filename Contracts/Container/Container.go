package Container

type Container interface {
	Singleton(string, interface{})

	Bind(string, interface{}, ...bool)

	Make(string) interface{}

	Alias(string, string)

	Build(interface{}, string) interface{}
}
