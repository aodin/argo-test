package server

import (
	"log"
	"net/http"
	"path"

	"github.com/aodin/argo"
	sql "github.com/aodin/aspect"
	"github.com/aodin/volta/config"
	"github.com/aodin/volta/templates"
	"github.com/julienschmidt/httprouter"

	db "github.com/aodin/argo-test/db"
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

func New(c config.Config, conn sql.Connection) *App {
	// Parse templates with the static URL as a local template variable
	locals := templates.Attrs{"StaticURL": c.StaticURL}

	app := &App{
		config:    c,
		db:        conn,
		templates: templates.NewWithDelims(c.TemplateDir, `<%`, `%>`, locals),
		router:    httprouter.New(),
	}

	// Just an API for now
	api := argo.New().SetPrefix("/api/")
	api.Add(
		"companies",
		argo.NewJSONResource(conn, db.Companies),
		db.Companies.PrimaryKey()...,
	)

	api.Add(
		"industries",
		argo.NewJSONResource(conn, db.Industries),
		db.Industries.PrimaryKey()...,
	)

	api.Add(
		"company-industries",
		argo.NewJSONResource(conn, db.CompanyIndustries),
		db.CompanyIndustries.PrimaryKey()...,
	)

	// Attach the api
	app.router.Handler("GET", "/api/*api", api)
	app.router.Handler("POST", "/api/*api", api)
	app.router.Handler("PATCH", "/api/*api", api)
	app.router.Handler("PUT", "/api/*api", api)
	app.router.Handler("DELETE", "/api/*api", api)

	// Static files
	app.router.ServeFiles(
		path.Join(c.StaticURL, "*filepath"),
		http.Dir(c.StaticDir),
	)

	// Default routes
	app.router.GET("/", app.Index)
	app.router.GET("/hello/:name", app.Hello)

	return app
}
