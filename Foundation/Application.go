package Foundation

import (
	"path"
	"path/filepath"
	"strings"
)

const VERSION = "1.0.0"

type App struct {

	/**
	 * the larisgo framework version.
	 *
	 * @var string
	 */
	version string

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
}

/**
 * Get the version number of the application.
 *
 * @return string
 */
func (app *App) Version() string {
	return app.version
}

/**
 * Set the base path for the application.
 *
 * @param  string  basePath
 * @return $this
 */
func (app *App) SetBasePath(basePath string) *App {
	app.basePath = strings.TrimRight(basePath, `\/`)
	return app
}

/**
 * Get the path to the application "app" directory.
 *
 * @param  string  path
 * @return string
 */
func (app *App) Path(_path string) string {
	var appPath string
	if app.appPath == "" {
		appPath = filepath.Clean(path.Join(app.basePath, "app"))
	} else {
		appPath = app.appPath
	}
	if _path != "" {
		return filepath.Clean(path.Join(appPath, _path))
	}
	return appPath
}

func Application(basePath string) (app *App) {

	app = &App{}

	app.version = VERSION

	if basePath != "" {
		app.SetBasePath(basePath)
	}

	return app
}
