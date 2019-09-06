package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// App ...
type App struct {
	router  *mux.Router
	service *Service
}

// Initialise ...
func (app *App) Initialise() {
	app.router = mux.NewRouter()

	db := PostgresDB{}
	app.service = &Service{}

	var adder ShortenedURLAdder = &db
	var getter ShortenedURLGetter = &db

	app.service.Initialise(app.router, adder, getter)
}

// Run ...
func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":80", app.router))
}
