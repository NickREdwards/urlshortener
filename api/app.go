package main

import (
	"log"
	"net/http"

	"github.com/NickREdwards/urlshortener/api/dal"
	"github.com/gorilla/mux"
)

// App ...
type App struct {
	router  *mux.Router
	service *Service
}

// Initialise ...
func (app *App) Initialise(urlAdder dal.ShortenedURLAdder, urlGetter dal.ShortenedURLGetter) {
	app.router = mux.NewRouter()
	app.service = &Service{}
	app.service.Initialise(app.router, urlAdder, urlGetter)
}

// Run ...
func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":80", app.router))
}
