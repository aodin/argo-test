package server

import (
	"net/http"

	"github.com/aodin/volta/templates"
	"github.com/julienschmidt/httprouter"
)

func (app *App) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	app.Execute(w, "index")
}

func (app *App) Companies(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	app.Execute(w, "companies")
}

func (app *App) Industries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	app.Execute(w, "industries")
}

func (app *App) Detail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	app.Execute(w, "detail", templates.Attrs{"ID": ps.ByName("id")})
}

func (app *App) Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	app.Execute(w, "hello", templates.Attrs{"Name": name})
}
