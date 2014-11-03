package server

import (
	"net/http"

	"github.com/aodin/volta/templates"
	"github.com/julienschmidt/httprouter"
)

func (app *App) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	app.Execute(w, "index")
}

func (app *App) Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	app.Execute(w, "hello", templates.Attrs{"Name": name})
}
