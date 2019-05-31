package Container

type Container struct {
	/**
	 * The current globally available container (if any).
	 */
	instance interface{}

	/**
	 * The container's bindings.
	 *
	 * @var slice
	 */
	bindings map[string]interface{}
}

func NewContainer() (this *Container) {
	this = &Container{}
	return this
}

/**
 * Register a shared binding in the container.
 *
 * @param  string  abstract
 * @param  \Closure|string|nil  concrete
 * @return void
 */
func (this *Container) Singleton(abstract string, concrete ...interface{}) {
	concrete = append(concrete, nil)
	this.Bind(abstract, concrete[0], true)
}

/**
 * Register a binding with the container.
 *
 * @param  string  abstract
 * @param  \Closure|string|nil  concrete
 * @param  bool  shared
 * @return void
 */
func (this *Container) Bind(abstract string, concrete interface{}, shared bool) {
	this.bindings[abstract] = []interface{}{
		concrete,
		shared,
	}
}

/**
 * Resolve the given type from the container.
 *
 * @param  string  abstract
 * @param  slice  parameters
 * @return mixed
 */
func (this *Container) Make(abstract interface{}, parameters ...interface{}) interface{} {
	// fmt.Println(this)
	return this.Resolve(abstract, parameters)
}

/**
 * Resolve the given type from the container.
 *
 * @param  string  abstract
 * @param  array  parameters
 * @return mixed
 */
func (this *Container) Resolve(abstract interface{}, parameters ...interface{}) interface{} {
	return abstract
}
