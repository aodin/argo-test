package server

import (
	"log"
	"net/http"
	"path"

	sql "github.com/aodin/aspect"
	"github.com/aodin/volta/config"
	"github.com/aodin/volta/templates"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	config    config.Config
	db        sql.Connection
	router    *httprouter.Router
	templates *templates.Templates
}

func (app *App) Execute(w http.ResponseWriter, name string, attrs ...templates.Attrs) {
	app.templates.Execute(w, name, attrs...)
}

func (app *App) ListenAndServe() error {
	log.Printf("server: serving on address %s\n", app.config.Address())
	return http.ListenAndServe(app.config.Address(), app.router)
}

func New(c config.Config, db sql.Connection) *App {
	// Parse templates with the static URL as a local template variable
	locals := templates.Attrs{"StaticURL": c.StaticURL}

	app := &App{
		config:    c,
		db:        db,
		templates: templates.New(c.TemplateDir, locals),
		router:    httprouter.New(),
	}

	app.router.ServeFiles(
		path.Join(c.StaticURL, "*filepath"),
		http.Dir(c.StaticDir),
	)

	app.router.GET("/", app.Index)
	app.router.GET("/hello/:name", app.Hello)

	return app
}
