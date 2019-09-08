package main

const (
	postgresHost     = "db"
	postgresPort     = 5432
	postgresUser     = "postgres"
	postgresPassword = "docker"
	postgresDBName   = "urlshortener"

	shortCodeLength = 6
)

func main() {
	var app App
	db := PostgresDB{config: PostgresDBConfig{host: postgresHost, port: postgresPort, user: postgresUser, password: postgresPassword, dbname: postgresDBName}}
	serviceParams := serviceParams{&db, &db, &db, &db}

	app.Initialise(serviceParams)
	app.Run()
}
