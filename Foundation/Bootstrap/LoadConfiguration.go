package Bootstrap

import (
	AppConfig "App/Config"
	"github.com/larisgo/framework/Config"
	RepositoryContract "github.com/larisgo/framework/Contracts/Config"
	"github.com/larisgo/framework/Contracts/Foundation"
	"github.com/larisgo/framework/Errors"
)

type LoadConfiguration struct {
}

/**
 * Bootstrap the given application.
 *
 * @param  Foundation\Application  app
 * @return void
 */
func (this *LoadConfiguration) Bootstrap(app Foundation.Application) {

	// Next we will spin through all of the configuration files in the configuration
	// directory and load each one into the repository. This will make all of the
	// options available to the developer for use in various parts of this app.
	config := Config.NewRepository()
	app.Instance("config", config)

	this.loadConfigurationFiles(app, config)

	// Finally, we will set the application's environment based on the configuration
	// values that were loaded. We will pass a callback which will be used to get
	// the environment in a web context where an "--env" switch is not present.
	// app.DetectEnvironment(func() string {
	// 	if v, ok := config.Get("app.env", "production").(string); ok {
	// 		return v
	// 	}
	// 	return "production"
	// })
	// date_default_timezone_set(config.get('app.timezone', 'UTC'));

	// mb_internal_encoding('UTF-8');
}

/**
 * Load the configuration items from all of the files.
 *
 * @param  Foundation\Application  app
 * @param  Config\Repository  repository
 * @return void
 *
 * @throws \Exception
 */
func (this *LoadConfiguration) loadConfigurationFiles(app Foundation.Application, repository RepositoryContract.Repository) {

	if _, ok := AppConfig.Map["app"]; !ok {
		panic(Errors.NewException(`Unable to load the "app" configuration file.`))
	}

	for key, value := range AppConfig.Map {
		repository.Set(key, value)
	}
}
