package main

import "github.com/NickREdwards/urlshortener/api/dal"

func main() {
	var app App
	var db dal.PostgresDB
	app.Initialise(&db, &db)
	app.Run()
}
