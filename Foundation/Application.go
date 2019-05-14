package Foundation

import (
	"github.com/larisgo/larisgo/Container"
	"path"
	"path/filepath"
	"strings"
)

const VERSION = "1.0.0"

type App struct {
	*Container.Container

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
func (this *App) Version() string {
	return this.version
}

/**
 * Set the base path for the application.
 *
 * @param  string  basePath
 * @return $this
 */
func (this *App) SetBasePath(basePath string) *App {
	this.basePath = strings.TrimRight(basePath, `\/`)
	return this
}

/**
 * Get the path to the application "app" directory.
 *
 * @param  string  path
 * @return string
 */
func (this *App) Path(_path ...string) string {
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
func (this *App) UseAppPath(_path ...string) *App {
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
func (this *App) BasePath(_path ...string) string {
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
func (this *App) BootstrapPath(_path ...string) string {
	_path = append(_path, "")
	return filepath.Clean(path.Join(this.basePath, "bootstrap", _path[0]))
}

/**
 * Get the path to the application configuration files.
 *
 * @param  string  path Optionally, a path to append to the config path
 * @return string
 */
func (this *App) ConfigPath(_path ...string) string {
	_path = append(_path, "")
	return filepath.Clean(path.Join(this.basePath, "config", _path[0]))
}

/**
 * Get the path to the database directory.
 *
 * @param  string  path Optionally, a path to append to the database path
 * @return string
 */
func (this *App) DatabasePath(_path ...string) string {
	_path = append(_path, "")
	return filepath.Clean(path.Join(this.basePath, "database", _path[0]))
}

func Application(basePath string) (app *App) {

	app = &App{}

	app.version = VERSION

	if basePath != "" {
		app.SetBasePath(basePath)
	}

	return app
}
