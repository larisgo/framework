package Container

import (
	"fmt"
	"github.com/larisgo/framework/Errors"
)

type Bindings struct {
	Concrete interface{}
	Shared   bool
}

type Container struct {
	instance *Container

	/**
	 * An array of the types that have been resolved.
	 *
	 * @var array
	 */
	resolved map[string]interface{}

	/**
	 * The current globally available container (if any).
	 */
	bindings map[string]*Bindings

	/**
	 * The container's method bindings.
	 *
	 * @var array
	 */
	methodBindings map[string]interface{}

	/**
	 * The container's shared instances.
	 *
	 * @var array
	 */
	instances map[string]interface{}

	/**
	 * The registered type aliases.
	 *
	 * @var array
	 */
	aliases map[string]string

	/**
	 * The registered aliases keyed by the abstract name.
	 *
	 * @var array
	 */
	abstractAliases map[string]interface{}

	/**
	 * The extension closures for services.
	 *
	 * @var array
	 */
	extenders map[string]interface{}

	/**
	 * All of the registered tags.
	 *
	 * @var array
	 */
	tags map[string]interface{}

	/**
	 * The stack of concretions currently being built.
	 *
	 * @var array
	 */
	buildStack map[string]interface{}

	/**
	 * The parameter override stack.
	 *
	 * @var array
	 */
	with map[string]interface{}

	/**
	 * The contextual binding map.
	 *
	 * @var array
	 */
	Contextual map[string]interface{}

	/**
	 * All of the registered rebound callbacks.
	 *
	 * @var array
	 */
	reboundCallbacks map[string]interface{}

	/**
	 * All of the global resolving callbacks.
	 *
	 * @var array
	 */
	globalResolvingCallbacks map[string]interface{}

	/**
	 * All of the global after resolving callbacks.
	 *
	 * @var array
	 */
	globalAfterResolvingCallbacks map[string]interface{}

	/**
	 * All of the resolving callbacks by class type.
	 *
	 * @var array
	 */
	resolvingCallbacks map[string]interface{}

	/**
	 * All of the after resolving callbacks by class type.
	 *
	 * @var array
	 */
	afterResolvingCallbacks map[string]interface{}
}

func NewContainer() (this *Container) {
	this = &Container{}
	return this
}

/**
 * Determine if the given abstract type has been bound.
 *
 * @param  string  abstract
 * @return bool
 */
func (this *Container) Bound(abstract string) bool {
	_, bindingsExists := this.bindings[abstract]
	_, instancesExists := this.instances[abstract]
	return bindingsExists || instancesExists || this.IsAlias(abstract)
}

/**
 *  {@inheritdoc}
 */
func (this *Container) Has(id string) bool {
	return this.Bound(id)
}

/**
 * Determine if the given abstract type has been resolved.
 *
 * @param  string  abstract
 * @return bool
 */
func (this *Container) Resolved(abstract string) bool {
	if this.IsAlias(abstract) {
		abstract = this.GetAlias(abstract)
	}
	_, resolvedExists := this.resolved[abstract]
	_, instancesExists := this.instances[abstract]

	return resolvedExists || instancesExists
}

/**
 * Determine if a given type is shared.
 *
 * @param  string  abstract
 * @return bool
 */
func (this *Container) IsShared(abstract string) bool {
	_, instancesExists := this.instances[abstract]
	bindingsAbstract, bindingsExists := this.bindings[abstract]

	return instancesExists || (bindingsExists && bindingsAbstract.Shared)
}

/**
 * Determine if a given string is an alias.
 *
 * @param  string  name
 * @return bool
 */
func (this *Container) IsAlias(name string) bool {
	_, ok := this.aliases[name]
	return pk
}

/**
 * Get the alias for an abstract if available.
 *
 * @param  string  abstract
 * @return string
 *
 * @throws \LogicException
 */
func (this *Container) GetAlias(abstract string) string {
	if _, ok := this.aliases[abstract]; !ok {
		return abstract
	}

	if v, ok := this.aliases[abstract]; v == abstract {
		painc(Errors.NewLogicException(fmt.Sprintf(`[%s] is aliased to itself.`, abstract)))
	}

	return this.GetAlias(this.aliases[abstract])
}

/**
 * Register a shared binding in the container.
 *
 * @param  string  abstract
 * @param  interface{}  concrete
 * @return void
 */
func (this *Container) Singleton(abstract string, concrete interface{}) {
	this.Bind(abstract, concrete, true)
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
	this.dropStaleInstances(abstract)
	this.bindings[abstract] = &Bindings{
		Concrete: concrete,
		Shared:   shared,
	}
	// If the abstract type was already resolved in this container we'll fire the
	// rebound listener so that any objects which have already gotten resolved
	// can have their copy of the object updated via the listener callbacks.
	if this.Resolved(abstract) {
		this.rebound(abstract)
	}
}

/**
 * Fire the "rebound" callbacks for the given abstract type.
 *
 * @param  string  abstract
 * @return void
 */
func (this *Container) rebound(abstract string, concrete interface{}, shared bool) {
	instance := this.Make(abstract)

	// foreach (this.getReboundCallbacks(abstract) as callback) {
	//     call_user_func(callback, this, instance);
	// }
}

/**
 * Resolve the given type from the container.
 *
 * @param  string  abstract
 * @param  slice  parameters
 * @return mixed
 */
func (this *Container) Make(abstract string, parameters ...interface{}) interface{} {
	return this.Resolve(abstract, parameters...)
}

/**
 * Resolve the given type from the container.
 *
 * @param  string  abstract
 * @param  array  parameters
 * @return mixed
 */
func (this *Container) Resolve(abstract string, parameters ...interface{}) interface{} {
	abstract = this.GetAlias(abstract)

	return abstract
}

/**
 * Drop all of the stale instances and aliases.
 *
 * @param  string  abstract
 * @return void
 */
func (this *Container) dropStaleInstances(abstract string) {
	delete(this.instances, abstract)
	delete(this.aliases, abstract)
}
