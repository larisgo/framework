package Container

import (
	"fmt"
	ContainerInterface "github.com/larisgo/framework/Contracts/Container"
	"github.com/larisgo/framework/Errors"
	"reflect"
)

type closureT func(ContainerInterface.Container, interface{})
type extenderT func(interface{}, ContainerInterface.Container) interface{}

type BindingsT struct {
	Concrete func(ContainerInterface.Container) interface{}
	Shared   bool
}

type Container struct {
	// instance *Container

	/**
	 * An array of the types that have been resolved.
	 *
	 * @var array
	 */
	resolved map[string]bool

	/**
	 * The current globally available container (if any).
	 */
	bindings map[string]*BindingsT

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
	abstractAliases map[string]map[string]string

	/**
	 * The extension closures for services.
	 *
	 * @var array
	 */
	extenders map[string][]extenderT

	/**
	 * All of the registered tags.
	 *
	 * @var array
	 */
	tags map[string]interface{}

	/**
	 * All of the registered rebound callbacks.
	 *
	 * @var array
	 */
	reboundCallbacks map[string][]closureT

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

	this.resolved = map[string]bool{}
	this.bindings = map[string]*BindingsT{}
	this.methodBindings = map[string]interface{}{}
	this.instances = map[string]interface{}{}
	this.aliases = map[string]string{}
	this.abstractAliases = map[string]map[string]string{}
	this.extenders = map[string][]extenderT{}
	this.tags = map[string]interface{}{}
	this.reboundCallbacks = map[string][]closureT{}
	this.globalResolvingCallbacks = map[string]interface{}{}
	this.globalAfterResolvingCallbacks = map[string]interface{}{}
	this.resolvingCallbacks = map[string]interface{}{}
	this.afterResolvingCallbacks = map[string]interface{}{}

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
	return ok
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
	v, ok := this.aliases[abstract]
	if !ok {
		return abstract
	}
	if v == abstract {
		panic(Errors.NewLogicException(fmt.Sprintf(`[%s] is aliased to itself.`, abstract)))
	}

	return this.GetAlias(v)
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
 * @param  \closure|string|nil  concrete
 * @param  bool  shared
 * @return void
 */
func (this *Container) Bind(abstract string, concrete interface{}, shared ...bool) {
	shared = append(shared, true)
	this.dropStaleInstances(abstract)

	// If the factory is not a Closure, it means it is just a class name which is
	// bound into this container to the abstract type and we will just wrap it
	// up inside its own Closure to give us more convenience when extending.
	var Concrete (func(ContainerInterface.Container) interface{})
	if _, ok := concrete.(func(ContainerInterface.Container) interface{}); !ok {
		Concrete = this.getClosure(abstract, concrete)
	} else {
		Concrete = concrete.(func(ContainerInterface.Container) interface{})
	}

	this.bindings[abstract] = &BindingsT{
		Concrete: Concrete,
		Shared:   shared[0],
	}
	// If the abstract type was already resolved in this container we'll fire the
	// rebound listener so that any objects which have already gotten resolved
	// can have their copy of the object updated via the listener callbacks.
	if this.Resolved(abstract) {
		this.rebound(abstract)
	}
}

/**
 * Get the Closure to be used when building a type.
 *
 * @param  string  abstract
 * @param  string  concrete
 * @return \Closure
 */
func (this *Container) getClosure(abstract string, concrete interface{}) func(ContainerInterface.Container) interface{} {
	return func(container ContainerInterface.Container) interface{} {
		return container.Build(concrete, abstract)
	}
}

/**
 * Fire the "rebound" callbacks for the given abstract type.
 *
 * @param  string  abstract
 * @return void
 */
func (this *Container) rebound(abstract string) {
	instance := this.Make(abstract)

	for _, callback := range this.getReboundCallbacks(abstract) {
		callback(this, instance)
	}
}

/**
 * Get the rebound callbacks for a given type.
 *
 * @param  string  abstract
 * @return array
 */
func (this *Container) getReboundCallbacks(abstract string) []closureT {
	if v, ok := this.reboundCallbacks[abstract]; ok {
		return v
	}

	return []closureT{}
}

/**
 * Resolve the given type from the container.
 *
 * @param  string  abstract
 * @param  slice  parameters
 * @return mixed
 */
func (this *Container) Make(abstract string) interface{} {
	return this.Resolve(abstract)
}

/**
 * Get the concrete type for a given abstract.
 *
 * @param  string  abstract
 * @return mixed   concrete
 */
func (this *Container) getConcrete(abstract string) interface{} {

	// If we don't have a registered resolver or concrete for the type, we'll just
	// assume each type is a concrete name and will attempt to resolve it as is
	// since the container should be able to resolve concretes automatically.
	if bindingsAbstract, ok := this.bindings[abstract]; ok {
		return bindingsAbstract.Concrete
	}

	return nil
}

/**
 * Get the extender callbacks for a given type.
 *
 * @param  string  abstract
 * @return array
 */
func (this *Container) getExtenders(abstract string) []extenderT {
	abstract = this.GetAlias(abstract)

	if extenderAbstract, ok := this.extenders[abstract]; ok {
		return extenderAbstract
	}

	return []extenderT{}
}

/**
 * "Extend" an abstract type in the container.
 *
 * @param  string    abstract
 * @param  \closure  closure
 * @return void
 *
 * @throws \InvalidArgumentException
 */
func (this *Container) Extend(abstract string, closure extenderT) {
	abstract = this.GetAlias(abstract)

	if instancesAbstract, ok := this.instances[abstract]; ok {
		this.instances[abstract] = closure(instancesAbstract, this)

		this.rebound(abstract)
	} else {
		this.extenders[abstract] = append(this.extenders[abstract], closure)

		if this.Resolved(abstract) {
			this.rebound(abstract)
		}
	}
}

/**
 * Alias a type to a different name.
 *
 * @param  string  abstract
 * @param  string  alias
 * @return void
 */
func (this *Container) Alias(abstract string, alias string) {
	this.aliases[alias] = abstract

	this.abstractAliases[abstract][alias] = alias
}

/**
 * Remove an alias from the contextual binding alias cache.
 *
 * @param  string  searched
 * @return void
 */
func (this *Container) removeAbstractAlias(searched string) {
	if _, ok := this.aliases[searched]; !ok {
		return
	}
	for abstract, aliases := range this.abstractAliases {
		for index, alias := range aliases {
			if alias == searched {
				delete(this.abstractAliases[abstract], index)
			}
		}
	}
}

/**
 * Register an existing instance as shared in the container.
 *
 * @param  string  abstract
 * @param  mixed   instance
 * @return mixed
 */
func (this *Container) Instance(abstract string, instance interface{}) interface{} {
	this.removeAbstractAlias(abstract)

	isBound := this.Bound(abstract)

	delete(this.aliases, abstract)

	// We'll check to determine if this type has been bound before, and if it has
	// we will fire the rebound callbacks registered with the container and it
	// can be updated with consuming classes that have gotten resolved here.
	this.instances[abstract] = instance

	if isBound {
		this.rebound(abstract)
	}

	return instance
}

/**
 * Instantiate a concrete instance of the given type.
 *
 * @param  interface{}  concrete
 * @return mixed
 */
func (this *Container) Build(concrete interface{}, abstract string) interface{} {
	if concrete == nil {
		panic(Errors.NewBindingResolutionException(fmt.Sprintf(`Target [%s] is not instantiable.`, abstract)))
	}

	if Closure, ok := concrete.(func(ContainerInterface.Container) interface{}); ok {
		return Closure(this)
	}

	Type := reflect.TypeOf(concrete)
	if T := Type.Kind(); T != reflect.Ptr {
		panic(Errors.NewBindingResolutionException(fmt.Sprintf(`Target [%s] is not pointer.`, abstract)))
	} else {
		if T := Type.Elem().Kind(); T != reflect.Struct {
			panic(Errors.NewBindingResolutionException(fmt.Sprintf(`Target [%s] is not struct.`, abstract)))
		}
		Type = Type.Elem()
	}
	Concrete := reflect.ValueOf(concrete).Elem()
	for i := 0; i < Type.NumField(); i++ { // 遍历字段
		fieldType := Type.Field(i)
		inject := fieldType.Tag.Get("inject") // 获取tag
		if inject == "" {
			continue
		}
		instance, ok := this.instances[inject]
		if !ok {
			if instance = this.resolveClass(inject); instance == nil {
				panic(Errors.NewBindingResolutionException(fmt.Sprintf(`"Unresolvable dependency resolving [%s] in struct %s `, inject, abstract)))
			}
		}
		Concrete.Field(i).Set(reflect.ValueOf(instance))
	}
	return concrete
}

/**
 * Resolve a class based dependency from the container.
 *
 * @param  string parameter
 * @return mixed
 *
 */
func (this *Container) resolveClass(parameter string) interface{} {
	return this.Make(parameter)
}

/**
 * Resolve the given type from the container.
 *
 * @param  string  abstract
 * @return mixed
 */
func (this *Container) Resolve(abstract string) interface{} {
	abstract = this.GetAlias(abstract)

	// If an instance of the type is currently being managed as a singleton we'll
	// just return an existing instance instead of instantiating new instances
	// so the developer can keep using the same objects instance every time.
	if instance, ok := this.instances[abstract]; ok {
		return instance
	}

	object := this.getConcrete(abstract)

	if object == nil {
		return nil
	}

	// We're ready to instantiate an instance of the concrete type registered for
	// the binding. This will instantiate the types, as well as resolve any of
	// its "nested" dependencies recursively until all have gotten resolved.
	object = this.Build(object, abstract)

	// If we defined any extenders for this type, we'll need to spin through them
	// and apply them to the object being built. This allows for the extension
	// of services, such as changing configuration or decorating the object.
	for _, extender := range this.getExtenders(abstract) {
		object = extender(object, this)
	}
	// If the requested type is registered as a singleton we'll want to cache off
	// the instances in "memory" so we can return it later without creating an
	// entirely new instance of an object on each subsequent request for it.
	if this.IsShared(abstract) {
		this.instances[abstract] = object
	}

	this.resolved[abstract] = true

	return object
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

/**
 * Remove a resolved instance from the instance cache.
 *
 * @param  string  abstract
 * @return void
 */
func (this *Container) ForgetInstance(abstract string) {
	delete(this.instances, abstract)
}

/**
 * Clear all of the instances from the container.
 *
 * @return void
 */
func (this *Container) ForgetInstances() {
	this.instances = map[string]interface{}{}
}
