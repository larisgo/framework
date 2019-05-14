package Container

type Container struct {
	/**
	 * The current globally available container (if any).
	 */
	instance interface{}
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
}
