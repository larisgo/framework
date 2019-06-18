package Foundation

import (
	"github.com/larisgo/framework/Contracts/Container"
)

type Application interface {
	Container.Container

	/**
	 * Get the version number of the application.
	 *
	 * @return string
	 */
	Version() string

	/**
	 * Get the base path of the Laravel installation.
	 *
	 * @return string
	 */
	BasePath(...string) string

	/**
	 * Get or check the current application environment.
	 *
	 * @return string
	 */
	// environment()

	/**
	 * Determine if the application is running in the console.
	 *
	 * @return bool
	 */
	// runningInConsole()

	/**
	 * Determine if the application is running unit tests.
	 *
	 * @return bool
	 */
	// runningUnitTests()

	/**
	 * Determine if the application is currently down for maintenance.
	 *
	 * @return bool
	 */
	// isDownForMaintenance()

	/**
	 * Register all of the configured providers.
	 *
	 * @return void
	 */
	// registerConfiguredProviders()

	/**
	 * Register a service provider with the application.
	 *
	 * @param  \Illuminate\Support\ServiceProvider|string  provider
	 * @param  bool   force
	 * @return \Illuminate\Support\ServiceProvider
	 */
	Register(interface{}, ...bool) interface{}

	/**
	 * Register a deferred provider and service.
	 *
	 * @param  string  provider
	 * @param  string|null  service
	 * @return void
	 */
	//  registerDeferredProvider(interface{}, service = null);

	/**
	 * Boot the application's service providers.
	 *
	 * @return void
	 */
	Boot()

	/**
	 * Register a new boot listener.
	 *
	 * @param  callable  callback
	 * @return void
	 */

	Booting(func(interface{}))

	/**
	 * Register a new "booted" listener.
	 *
	 * @param  callable  callback
	 * @return void
	 */
	Booted(func(interface{}))

	RegisterConfiguredProviders()

	/**
	 * Get the path to the cached services.php file.
	 *
	 * @return string
	 */
	// getCachedServicesPath()

	/**
	 * Get the path to the cached packages.php file.
	 *
	 * @return string
	 */
	// getCachedPackagesPath()
}
