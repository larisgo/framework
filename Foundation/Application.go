package Foundation

import (
	"fmt"
	"github.com/larisgo/framework/Container"
	RepositoryContract "github.com/larisgo/framework/Contracts/Config"
	FoundationContract "github.com/larisgo/framework/Contracts/Foundation"
	"github.com/larisgo/framework/Contracts/Service"
	"github.com/larisgo/framework/Errors"
	"github.com/larisgo/framework/Providers"
	"path"
	"path/filepath"
	"reflect"
	"strings"
)

const VERSION = "1.0.0"

type Application struct {
	*Container.Container

	/**
	 * the larisgo framework version.
	 *
	 * @var string
	 */
	version string

	/**
	 * The array of booting callbacks.
	 *
	 * @var callable[]
	 */
	bootingCallbacks []func(interface{})

	/**
	 * The array of booted callbacks.
	 *
	 * @var callable[]
	 */
	bootedCallbacks []func(interface{})

	/**
	 * The array of terminating callbacks.
	 *
	 * @var callable[]
	 */
	terminatingCallbacks []interface{}

	/**
	 * All of the registered service providers.
	 *
	 * @var ServiceProvider[]
	 */
	serviceProviders []interface{}

	/**
	 * The names of the loaded service providers.
	 *
	 * @var array
	 */
	loadedProviders map[string]bool

	/**
	 * The deferred services and their providers.
	 *
	 * @var array
	 */
	deferredServices []interface{}

	/**
	 * The base path for the larisgo installation.
	 *
	 * @var string
	 */
	basePath string

	/**
	 * The custom application path defined by the developer.
	 *
	 * @var string
	 */
	appPath string

	/**
	 * Indicates if the application has "booted".
	 *
	 * @var bool
	 */
	booted bool `default:false`

	/**
	 * Indicates if the application has been bootstrapped before.
	 *
	 * @var bool
	 */
	hasBeenBootstrapped bool `default:false`
}

func NewApplication(basePath string) (this *Application) {

	this = &Application{Container: Container.NewContainer()}

	this.version = VERSION
	this.bootingCallbacks = []func(interface{}){}
	this.bootedCallbacks = []func(interface{}){}
	this.terminatingCallbacks = []interface{}{}
	this.serviceProviders = []interface{}{}
	this.loadedProviders = map[string]bool{}
	this.deferredServices = []interface{}{}

	if basePath != "" {
		this.SetBasePath(basePath)
	}
	this.registerBaseBindings()
	this.registerBaseServiceProviders()

	return this
}

/**
 * Get the version number of the application.
 *
 * @return string
 */
func (this *Application) Version() string {
	return this.version
}

/**
 * Register the basic bindingsT into the container.
 *
 * @return void
 */
func (this *Application) registerBaseBindings() {
	// this.SetInstance(this)

	this.Instance("app", this)

	this.Instance("container", this)

	// this.instance(PackageManifest::class, new PackageManifest(
	// new Filesystem, this.basePath(), this.getCachedPackagesPath()
	// ));
}

/**
 * Register all of the base service providers.
 *
 * @return void
 */
func (this *Application) registerBaseServiceProviders() {
	this.Register(Providers.NewRoutingServiceProvider(this))
}

/**
 * Run the given array of bootstrap classes.
 *
 * @param  string[]  bootstrappers
 * @return void
 */
func (this *Application) BootstrapWith(bootstrappers []FoundationContract.BootstrapT) {
	this.hasBeenBootstrapped = true

	for _, bootstrapper := range bootstrappers {
		// this['events'].dispatch('bootstrapping: '.bootstrapper, [this]);
		bootstrapper.Bootstrap(this)
		// this['events'].dispatch('bootstrapped: '.bootstrapper, [this]);
	}
}

/**
 * Get the registered service provider instance if it exists.
 *
 * @param  ServiceProvider|string  provider
 * @return ServiceProvider|null
 */
func (this *Application) GetProvider(provider interface{}) interface{} {
	if provider := this.GetProviders(provider); len(provider) > 0 {
		return provider[0]
	}
	return nil
}

/**
 * Get the registered service provider instances if any exist.
 *
 * @param  ServiceProvider|string  provider
 * @return array
 */
func (this *Application) GetProviders(provider interface{}) []interface{} {
	_tmp := []interface{}{}
	for _, v := range this.serviceProviders {
		if reflect.TypeOf(v).String() == reflect.TypeOf(provider).String() {
			_tmp = append(_tmp, v)
		}
	}
	return _tmp
}

/**
 * Mark the given provider as registered.
 *
 * @param  ServiceProvider  provider
 * @return void
 */
func (this *Application) markAsRegistered(provider interface{}) {
	this.serviceProviders = append(this.serviceProviders, provider)

	this.loadedProviders[reflect.TypeOf(provider).String()] = true
}

/**
 * Boot the given service provider.
 *
 * @param  ServiceProvider  provider
 * @return mixed
 */
func (this *Application) bootProvider(provider interface{}) {
	if p, ok := provider.(Service.BootT); ok {
		p.Boot()
	}
}

/**
 * Register a new boot listener.
 *
 * @param  callable  callback
 * @return void
 */
func (this *Application) Booting(callback func(interface{})) {
	this.bootingCallbacks = append(this.bootingCallbacks, callback)
}

/**
 * Register a new "booted" listener.
 *
 * @param  callable  callback
 * @return void
 */
func (this *Application) Booted(callback func(interface{})) {
	this.bootedCallbacks = append(this.bootedCallbacks, callback)

	if this.IsBooted() {
		this.fireAppCallbacks([]func(interface{}){callback})
	}
}

/**
 * Determine if the application has booted.
 *
 * @return bool
 */
func (this *Application) IsBooted() bool {
	return this.booted
}

/**
 * Boot the application's service providers.
 *
 * @return void
 */
func (this *Application) Boot() {
	if this.booted {
		return
	}

	// Once the application has booted we will also fire some "booted" callbacks
	// for any listeners that need to do work after this initial booting gets
	// finished. This is useful when ordering the boot-up processes we run.
	this.fireAppCallbacks(this.bootingCallbacks)

	for _, p := range this.serviceProviders {
		this.bootProvider(p)
	}

	this.booted = true

	this.fireAppCallbacks(this.bootedCallbacks)
}

/**
 * Call the booting callbacks for the application.
 *
 * @param  callable[]  callbacks
 * @return void
 */
func (this *Application) fireAppCallbacks(callbacks []func(interface{})) {
	for _, callback := range callbacks {
		callback(this)
	}
}

/**
 * Register all of the configured providers.
 *
 * @return void
 */
func (this *Application) RegisterConfiguredProviders() {
	for _, instance := range this.Make("config").(RepositoryContract.Repository).Get("app.providers").([]Service.Provider) {
		this.Register(this.Build(instance, reflect.TypeOf(instance).String()))
	}
}

/**
 * Register a service provider with the application.
 *
 * @param  ServiceProvider|string  provider
 * @param  bool   force
 * @return ServiceProvider
 */
func (this *Application) Register(provider interface{}, force ...bool) interface{} {
	force = append(force, false)
	_provider := reflect.TypeOf(provider)
	if T := _provider.Kind(); T != reflect.Ptr {
		if T != reflect.Struct {
			panic(Errors.NewRuntimeException(fmt.Sprintf(`Provider is a [%s] is not a struct.`, _provider.String())))
		}
	} else {
		if T := _provider.Elem().Kind(); T != reflect.Struct {
			panic(Errors.NewRuntimeException(fmt.Sprintf(`Provider is a [%s] is not a struct.`, _provider.String())))
		}
	}

	if registered := this.GetProvider(provider); registered != nil && !force[0] {
		return registered
	}

	// If the given "provider" is a string, we will resolve it, passing in the
	// application instance automatically for the developer. This is simply
	// a more convenient way of specifying your service provider classes.
	// if is_string(provider) {
	// 	provider = this.resolveProvider(provider)
	// }

	if p, ok := provider.(Service.RegisterT); ok {
		p.Register()
	}

	// If there are bindingsT / singletonsT set as properties on the provider we
	// will spin through them and register them with the application, which
	// serves as a convenience layer while registering a lot of bindingsT.
	if p, ok := provider.(Service.BindingsT); ok {
		for key, value := range p.Bindings() {
			this.Bind(key, value)
		}
	}

	if p, ok := provider.(Service.SingletonsT); ok {
		for key, value := range p.Singletons() {
			this.Singleton(key, value)
		}
	}

	this.markAsRegistered(provider)

	// If the application has already booted, we will call this boot method on
	// the provider class so it has an opportunity to do its boot logic and
	// will be ready for any usage by this developer's application logic.
	if this.booted {
		this.bootProvider(provider)
	}

	return provider
}

/**
 * Determine if the application has been bootstrapped before.
 *
 * @return bool
 */
func (this *Application) HasBeenBootstrapped() bool {
	return this.hasBeenBootstrapped
}

/**
 * Set the base path for the application.
 *
 * @param  string  basePath
 * @return this
 */
func (this *Application) SetBasePath(basePath string) *Application {
	this.basePath = strings.TrimRight(basePath, `\/`)
	return this
}

/**
 * Get the path to the application "app" directory.
 *
 * @param  string  path
 * @return string
 */
func (this *Application) Path(_path ...string) string {
	_path = append(_path, "")

	var appPath string

	if this.appPath == "" {
		appPath = filepath.Clean(path.Join(this.basePath, "app"))
	} else {
		appPath = this.appPath
	}

	if _path[0] != "" {
		return filepath.Clean(path.Join(appPath, _path[0]))
	}

	return appPath
}

/**
 * Set the application directory.
 *
 * @param  string  path
 * @return this
 */
func (this *Application) UseAppPath(_path ...string) *Application {
	_path = append(_path, "")
	this.appPath = _path[0]
	return this
}

/**
 * Get the base path of the Laravel installation.
 *
 * @param  string  path Optionally, a path to append to the base path
 * @return string
 */
func (this *Application) BasePath(_path ...string) string {
	_path = append(_path, "")
	if this.appPath == "" {
		return filepath.Clean(path.Join(this.basePath, _path[0]))
	}
	return this.appPath
}

/**
 * Get the path to the bootstrap directory.
 *
 * @param  string  path Optionally, a path to append to the bootstrap path
 * @return string
 */
func (this *Application) BootstrapPath(_path ...string) string {
	_path = append(_path, "")
	return filepath.Clean(path.Join(this.basePath, "bootstrap", _path[0]))
}

/**
 * Get the path to the application configuration files.
 *
 * @param  string  path Optionally, a path to append to the config path
 * @return string
 */
func (this *Application) ConfigPath(_path ...string) string {
	_path = append(_path, "")
	return filepath.Clean(path.Join(this.basePath, "config", _path[0]))
}

/**
 * Get the path to the database directory.
 *
 * @param  string  path Optionally, a path to append to the database path
 * @return string
 */
func (this *Application) DatabasePath(_path ...string) string {
	_path = append(_path, "")
	return filepath.Clean(path.Join(this.basePath, "database", _path[0]))
}

/**
 * Terminate the application.
 *
 * @return void
 */
func (this *Application) Terminate() {
	// foreach (this.terminatingCallbacks as terminating) {
	//     this.call(terminating);
	// }
}
