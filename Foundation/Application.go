package Foundation

type App struct {
	VERSION string
}

func (app *App) Version() string {
	return app.VERSION
}

func Application() (app *App) {
	app = &App{}
	app.VERSION = "1.0.0"
	return app
}
