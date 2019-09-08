package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// App ...
type App struct {
	router  *mux.Router
	service *Service
}

// Initialise ...
func (app *App) Initialise(serviceParams serviceParams) {
	rand.Seed(time.Now().UTC().UnixNano())

	app.router = mux.NewRouter()
	app.service = &Service{}
	app.service.Initialise(app.router, serviceParams)
}

// Run ...
func (app *App) Run() {
	log.Fatal(http.ListenAndServe(":80", app.router))
}
