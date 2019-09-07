package main

func main() {
	var app App
	var db PostgresDB
	app.Initialise(&db, &db)
	app.Run()
}
